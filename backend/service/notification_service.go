package service

import (
	"encoding/json"
	"log"
	"math"
	"net/http"

	"training-plan-api/data/response"
	"training-plan-api/model"
	"training-plan-api/repository"

	webpush "github.com/SherClockHolmes/webpush-go"
)

type NotificationServiceImpl struct {
	repo            repository.NotificationRepository
	pushRepo        repository.PushSubscriptionRepository
	vapidPublicKey  string
	vapidPrivateKey string
	vapidSubject    string
}

func NewNotificationServiceImpl(
	repo repository.NotificationRepository,
	pushRepo repository.PushSubscriptionRepository,
	vapidPublicKey, vapidPrivateKey, vapidSubject string,
) NotificationService {
	return &NotificationServiceImpl{
		repo:            repo,
		pushRepo:        pushRepo,
		vapidPublicKey:  vapidPublicKey,
		vapidPrivateKey: vapidPrivateKey,
		vapidSubject:    vapidSubject,
	}
}

func (s *NotificationServiceImpl) Create(userID uint, notifType model.NotificationType, title, message string) error {
	n := &model.Notification{
		UserID:  userID,
		Type:    notifType,
		Title:   title,
		Message: message,
	}
	if err := s.repo.Save(n); err != nil {
		return err
	}

	// Fire push notifications in the background (best effort).
	go s.sendPushAsync(userID, title, message)

	return nil
}

// sendPushAsync sends a Web Push to all of the user's registered devices.
// Failures are logged but never bubble up — the in-app notification has already
// been saved, so push is just a delivery enhancement.
func (s *NotificationServiceImpl) sendPushAsync(userID uint, title, body string) {
	if s.vapidPublicKey == "" || s.vapidPrivateKey == "" {
		return // VAPID not configured — silently skip
	}

	subs, err := s.pushRepo.FindByUserID(userID)
	if err != nil {
		log.Println("⚠ push: failed to load subscriptions:", err)
		return
	}

	payload, err := json.Marshal(map[string]string{
		"title":   title,
		"message": body,
		"url":     "/notifications",
	})
	if err != nil {
		log.Println("⚠ push: failed to marshal payload:", err)
		return
	}

	for _, sub := range subs {
		webSub := &webpush.Subscription{
			Endpoint: sub.Endpoint,
			Keys: webpush.Keys{
				P256dh: sub.P256dh,
				Auth:   sub.Auth,
			},
		}

		resp, err := webpush.SendNotification(payload, webSub, &webpush.Options{
			Subscriber:      s.vapidSubject,
			VAPIDPublicKey:  s.vapidPublicKey,
			VAPIDPrivateKey: s.vapidPrivateKey,
			TTL:             60, // seconds — drop if undelivered after 1 min
		})
		if err != nil {
			log.Println("⚠ push: send failed:", err)
			continue
		}
		// 404/410 = subscription is gone (user uninstalled the PWA, denied permissions, etc.)
		// Clean it up so we don't keep trying.
		if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusGone {
			_ = s.pushRepo.DeleteByEndpoint(sub.Endpoint)
		}
		resp.Body.Close()
	}
}

func (s *NotificationServiceImpl) FindByUser(userID uint, page, limit int) (response.PaginatedResponse[response.NotificationResponse], error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 10
	}
	offset := (page - 1) * limit

	notifications, total, err := s.repo.FindByUserID(userID, offset, limit)
	if err != nil {
		return response.PaginatedResponse[response.NotificationResponse]{}, err
	}

	items := make([]response.NotificationResponse, 0, len(notifications))
	for _, n := range notifications {
		items = append(items, response.NotificationResponse{
			ID:        n.ID,
			UserID:    n.UserID,
			Type:      string(n.Type),
			Title:     n.Title,
			Message:   n.Message,
			IsRead:    n.IsRead,
			CreatedAt: n.CreatedAt,
		})
	}

	return response.PaginatedResponse[response.NotificationResponse]{
		Items: items,
		Meta: response.PaginationMeta{
			Page:       page,
			Limit:      limit,
			TotalItems: total,
			TotalPages: int(math.Ceil(float64(total) / float64(limit))),
		},
	}, nil
}

func (s *NotificationServiceImpl) CountUnread(userID uint) (int64, error) {
	return s.repo.CountUnreadByUserID(userID)
}

func (s *NotificationServiceImpl) MarkAsRead(id, userID uint) error {
	if _, err := s.repo.FindByIDAndUserID(id, userID); err != nil {
		return err
	}
	return s.repo.MarkAsRead(id)
}

func (s *NotificationServiceImpl) MarkAllRead(userID uint) error {
	return s.repo.MarkAllReadByUser(userID)
}

func (s *NotificationServiceImpl) Delete(id, userID uint) error {
	if _, err := s.repo.FindByIDAndUserID(id, userID); err != nil {
		return err
	}
	return s.repo.Delete(id)
}

// -------- Web Push subscription management --------

func (s *NotificationServiceImpl) GetVAPIDPublicKey() string {
	return s.vapidPublicKey
}

func (s *NotificationServiceImpl) SubscribePush(userID uint, endpoint, p256dh, auth, userAgent string) error {
	sub := &model.PushSubscription{
		UserID:    userID,
		Endpoint:  endpoint,
		P256dh:    p256dh,
		Auth:      auth,
		UserAgent: userAgent,
	}
	return s.pushRepo.Save(sub)
}

func (s *NotificationServiceImpl) UnsubscribePush(userID uint, endpoint string) error {
	return s.pushRepo.DeleteByUserAndEndpoint(userID, endpoint)
}

// SendTestPush is a diagnostic helper. It immediately sends a test push to every
// subscription registered for the user and returns the per-subscription result
// so the UI can show what went wrong.
func (s *NotificationServiceImpl) SendTestPush(userID uint) PushTestResult {
	result := PushTestResult{
		VAPIDConfigured: s.vapidPublicKey != "" && s.vapidPrivateKey != "",
		Results:         []PushSubscriptionResult{},
	}

	subs, err := s.pushRepo.FindByUserID(userID)
	if err != nil {
		log.Println("push-test: failed to load subs:", err)
		return result
	}
	result.SubscriptionsFound = len(subs)

	if !result.VAPIDConfigured || len(subs) == 0 {
		return result
	}

	payload, _ := json.Marshal(map[string]string{
		"title":   "Test Push",
		"message": "If you see this, Web Push is working!",
		"url":     "/notifications",
	})

	for _, sub := range subs {
		prefix := sub.Endpoint
		if len(prefix) > 60 {
			prefix = prefix[:60] + "..."
		}
		r := PushSubscriptionResult{EndpointPrefix: prefix}

		resp, sendErr := webpush.SendNotification(payload, &webpush.Subscription{
			Endpoint: sub.Endpoint,
			Keys:    webpush.Keys{P256dh: sub.P256dh, Auth: sub.Auth},
		}, &webpush.Options{
			Subscriber:      s.vapidSubject,
			VAPIDPublicKey:  s.vapidPublicKey,
			VAPIDPrivateKey: s.vapidPrivateKey,
			TTL:             60,
		})
		if sendErr != nil {
			r.Error = sendErr.Error()
			log.Println("push-test: send error:", sendErr)
		} else {
			r.StatusCode = resp.StatusCode
			log.Printf("push-test: endpoint=%s status=%d", prefix, resp.StatusCode)
			if resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusGone {
				_ = s.pushRepo.DeleteByEndpoint(sub.Endpoint)
				r.Error = "subscription expired (cleaned up)"
			}
			resp.Body.Close()
		}
		result.Results = append(result.Results, r)
	}

	return result
}
