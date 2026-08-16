package repository

import (
	"errors"
	"training-plan-api/helper"
	"training-plan-api/model"

	"gorm.io/gorm"
)

type NotificationRepositoryImpl struct {
	Db *gorm.DB
}

func NewNotificationRepositoryImpl(db *gorm.DB) NotificationRepository {
	return &NotificationRepositoryImpl{Db: db}
}

func (r *NotificationRepositoryImpl) Save(n *model.Notification) error {
	return r.Db.Create(n).Error
}

func (r *NotificationRepositoryImpl) FindByUserID(userID uint, offset, limit int) ([]model.Notification, int64, error) {
	var notifications []model.Notification
	var total int64

	base := r.Db.Model(&model.Notification{}).Where("user_id = ?", userID)

	if err := base.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.Db.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&notifications).Error

	return notifications, total, err
}

func (r *NotificationRepositoryImpl) FindByIDAndUserID(id, userID uint) (*model.Notification, error) {
	var n model.Notification
	err := r.Db.Where("id = ? AND user_id = ?", id, userID).First(&n).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.NotFound("notification not found")
		}
		return nil, err
	}
	return &n, nil
}

func (r *NotificationRepositoryImpl) CountUnreadByUserID(userID uint) (int64, error) {
	var count int64
	err := r.Db.Model(&model.Notification{}).
		Where("user_id = ? AND is_read = false", userID).
		Count(&count).Error
	return count, err
}

func (r *NotificationRepositoryImpl) MarkAsRead(id uint) error {
	return r.Db.Model(&model.Notification{}).Where("id = ?", id).Update("is_read", true).Error
}

func (r *NotificationRepositoryImpl) MarkAllReadByUser(userID uint) error {
	return r.Db.Model(&model.Notification{}).Where("user_id = ? AND is_read = false", userID).Update("is_read", true).Error
}

func (r *NotificationRepositoryImpl) Delete(id uint) error {
	result := r.Db.Delete(&model.Notification{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return helper.NotFound("notification not found")
	}
	return nil
}
