# New APIs — Postman Testing Reference

Endpoints added during the notification + push-notification work.
All endpoints live under the existing backend at `http://localhost:8080`.

---

## Authentication (required for every endpoint below)

All endpoints require a JWT in the `Authorization` header.

1. Login first via one of:
   - `POST /api/v1/auth/manager/login` → manager role
   - `POST /api/v1/auth/staff/login` → staff role
   - `POST /api/v1/auth/login` → generic (any role)
2. Copy the `accessToken` from the response.
3. In Postman: **Authorization → Type: Bearer Token → paste the token**.
   Or set header: `Authorization: Bearer <token>`.

**Important — role prefix in the URL must match your account's role:**
- If you logged in as a manager → use `/manager/...`
- If you logged in as a staff → use `/staff/...`
- Mixing them returns `403 Forbidden`.

---

## 1. Notifications (in-app)

### 1.1 List notifications (paginated)
- **Method:** `GET`
- **URL:** `http://localhost:8080/api/v1/{role}/notifications?page=1&limit=20`
- **Query params:** `page` (default 1), `limit` (default 10, max 100)
- **Response 200:**
```json
{
  "status": "SUCCESS",
  "message": "",
  "data": {
    "items": [
      {
        "id": 12,
        "userId": 5,
        "type": "certificate_approved",
        "title": "Certificate Approved",
        "message": "Your certificate for \"Safety Training\" has been approved.",
        "isRead": false,
        "createdAt": "2026-05-28T09:12:33Z"
      }
    ],
    "meta": { "page": 1, "limit": 20, "totalItems": 1, "totalPages": 1 }
  }
}
```
- **Notes:** Sorted newest first. `type` is one of `training_registered`, `certificate_approved`, `certificate_rejected`.

---

### 1.2 Unread count (used for the sidebar badge)
- **Method:** `GET`
- **URL:** `http://localhost:8080/api/v1/{role}/notifications/unread-count`
- **Response 200:**
```json
{
  "status": "SUCCESS",
  "message": "",
  "data": { "unread": 3 }
}
```

---

### 1.3 Mark one notification as read
- **Method:** `PUT`
- **URL:** `http://localhost:8080/api/v1/{role}/notifications/{id}/read`
- **Path param:** `id` — notification ID (from the list endpoint)
- **Response 200:**
```json
{ "status": "SUCCESS", "message": "Notification marked as read" }
```
- **Errors:** `404` if the notification doesn't belong to the current user.

---

### 1.4 Mark all notifications as read
- **Method:** `PUT`
- **URL:** `http://localhost:8080/api/v1/{role}/notifications/read-all`
- **Response 200:**
```json
{ "status": "SUCCESS", "message": "All notifications marked as read" }
```

---

### 1.5 Delete one notification
- **Method:** `DELETE`
- **URL:** `http://localhost:8080/api/v1/{role}/notifications/{id}`
- **Response 200:**
```json
{ "status": "SUCCESS", "message": "Notification deleted" }
```

---

## 2. Web Push (browser push notifications)

### 2.1 Get the VAPID public key
- **Method:** `GET`
- **URL:** `http://localhost:8080/api/v1/{role}/notifications/vapid-public-key`
- **Response 200:**
```json
{
  "status": "SUCCESS",
  "message": "",
  "data": { "publicKey": "BJArsFb-3m3DLoXXJ0Sho0puNCP_GrvQ06TKEVHTE2u3dJ6kWx7__2MfB3R2O-I8tDHlsf5F_TLBQqvwhVIQJWE" }
}
```
- **Notes:** Frontend uses this to build a `PushSubscription`. It's not secret — safe to expose to clients.

---

### 2.2 Subscribe to push notifications
- **Method:** `POST`
- **URL:** `http://localhost:8080/api/v1/{role}/notifications/push-subscribe`
- **Body (JSON):**
```json
{
  "endpoint": "https://fcm.googleapis.com/fcm/send/abcdef123...",
  "keys": {
    "p256dh": "BNcRdreALRFXTkOOUHK1EtK2wtaz5Ry4YfYCA_0QTpQtUbVlUls0VJXg7A8u-Ts1XbjhazAkj7I99e8QcYP7DkM",
    "auth":   "tBHItJI5svbpez7KI4CCXg"
  }
}
```
- **Response 200:**
```json
{ "status": "SUCCESS", "message": "Push subscription saved" }
```
- **Notes:** Postman is *not* the right tool to acquire a real `PushSubscription` — that comes from `navigator.pushManager.subscribe()` in the browser. Use this endpoint mainly to verify the response shape.

---

### 2.3 Unsubscribe from push
- **Method:** `POST`
- **URL:** `http://localhost:8080/api/v1/{role}/notifications/push-unsubscribe`
- **Body (JSON):**
```json
{ "endpoint": "https://fcm.googleapis.com/fcm/send/abcdef123..." }
```
- **Response 200:**
```json
{ "status": "SUCCESS", "message": "Push subscription removed" }
```

---

### 2.4 Send a test push (diagnostic)
- **Method:** `POST`
- **URL:** `http://localhost:8080/api/v1/{role}/notifications/push-test`
- **Body:** none
- **Response 200:**
```json
{
  "status": "SUCCESS",
  "message": "",
  "data": {
    "vapidConfigured": true,
    "subscriptionsFound": 2,
    "results": [
      { "endpoint": "https://fcm.googleapis.com/fcm/send/abc...", "statusCode": 201 },
      { "endpoint": "https://updates.push.services.mozilla.com/wpush/v2/xyz...", "statusCode": 410, "error": "subscription expired (cleaned up)" }
    ]
  }
}
```
- **Notes:**
  - `vapidConfigured: false` → backend is missing `VAPID_PUBLIC_KEY` / `VAPID_PRIVATE_KEY` in `app.env`.
  - `subscriptionsFound: 0` → this user never enabled push in a browser yet.
  - `statusCode: 201` → push was successfully sent to the push service (FCM/Mozilla/Apple); the OS handles delivery from there.
  - `statusCode: 404` or `410` → subscription is expired; the backend auto-deletes it.

---

## Notification triggers (no new endpoints, but behavior to know)

The backend automatically creates an in-app notification **and** fires a Web Push when:

| Event | Recipient | Notification Type | Title | Message |
|---|---|---|---|---|
| Manager registers staff to a training plan via `POST /api/v1/manager/training-plans/:trainingPlanId/registrations` | Each registered staff | `training_registered` | "Training Registration" | `You have been registered for "<plan name>".` |
| Admin approves a certificate via `PUT /api/v1/admin/certificates/:id/approve` | Certificate owner | `certificate_approved` | "Certificate Approved" | `Your certificate for "<training name>" has been approved.` |
| Admin rejects a certificate via `PUT /api/v1/admin/certificates/:id/reject` | Certificate owner | `certificate_rejected` | "Certificate Rejected" | `Your certificate for "<training name>" has been rejected.` |

So to **test the full chain** in Postman:

1. Login as admin → call `PUT /api/v1/admin/certificates/:id/approve` on a pending cert.
2. Login as the staff who owns that cert.
3. Call `GET /api/v1/staff/notifications` → you should see the new `certificate_approved` row.
4. Call `GET /api/v1/staff/notifications/unread-count` → unread = 1.
5. Call `PUT /api/v1/staff/notifications/{id}/read` → marks it read.
6. Call `GET /api/v1/staff/notifications/unread-count` again → unread = 0.

---

## Postman tips

- **Set the Bearer token once on the collection**, not per request. Open the collection → Authorization → Type: Bearer Token → paste. Every request inherits it.
- **Two environments**: create a `manager-token` and `staff-token` env var so you can swap between role-prefixed URLs and accounts quickly.
- **Rate limit**: backend allows 300 requests per 60 seconds per IP. If you hammer it you'll get a JSON `429`:
  ```json
  { "status": "ERROR", "message": "Too many requests. Please slow down and try again in a moment." }
  ```

---

## Summary table

| Method | Path (use `{role}` = `manager` or `staff`) | Purpose |
|---|---|---|
| GET    | `/api/v1/{role}/notifications` | List notifications (paginated) |
| GET    | `/api/v1/{role}/notifications/unread-count` | Unread count for sidebar badge |
| PUT    | `/api/v1/{role}/notifications/{id}/read` | Mark one as read |
| PUT    | `/api/v1/{role}/notifications/read-all` | Mark all as read |
| DELETE | `/api/v1/{role}/notifications/{id}` | Delete one |
| GET    | `/api/v1/{role}/notifications/vapid-public-key` | Frontend bootstrapping for Web Push |
| POST   | `/api/v1/{role}/notifications/push-subscribe` | Save a browser's push subscription |
| POST   | `/api/v1/{role}/notifications/push-unsubscribe` | Remove a push subscription |
| POST   | `/api/v1/{role}/notifications/push-test` | Send a diagnostic push to the current user |
