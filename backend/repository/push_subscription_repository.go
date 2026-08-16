package repository

import (
	"training-plan-api/model"

	"gorm.io/gorm"
)

type PushSubscriptionRepositoryImpl struct {
	Db *gorm.DB
}

func NewPushSubscriptionRepositoryImpl(db *gorm.DB) PushSubscriptionRepository {
	return &PushSubscriptionRepositoryImpl{Db: db}
}

// Save stores a subscription. If a subscription with the same endpoint already
// exists (same browser, different/same user), it is replaced.
func (r *PushSubscriptionRepositoryImpl) Save(sub *model.PushSubscription) error {
	// Replace any existing row for this endpoint (handles re-subscribe / account switch)
	if err := r.Db.Where("endpoint = ?", sub.Endpoint).
		Delete(&model.PushSubscription{}).Error; err != nil {
		return err
	}
	return r.Db.Create(sub).Error
}

func (r *PushSubscriptionRepositoryImpl) FindByUserID(userID uint) ([]model.PushSubscription, error) {
	var subs []model.PushSubscription
	err := r.Db.Where("user_id = ?", userID).Find(&subs).Error
	return subs, err
}

func (r *PushSubscriptionRepositoryImpl) DeleteByEndpoint(endpoint string) error {
	return r.Db.Where("endpoint = ?", endpoint).
		Delete(&model.PushSubscription{}).Error
}

func (r *PushSubscriptionRepositoryImpl) DeleteByUserAndEndpoint(userID uint, endpoint string) error {
	return r.Db.Where("user_id = ? AND endpoint = ?", userID, endpoint).
		Delete(&model.PushSubscription{}).Error
}
