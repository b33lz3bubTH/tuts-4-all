package database

import (
	"github.com/tuts-4-all/backend/internal/models"
	"gorm.io/gorm"
)

type NotificationRepository interface {
	Create(notification *models.Notification) error
	GetByID(id uint) (*models.Notification, error)
	List(offset, limit int) ([]models.Notification, int64, error)
	Update(notification *models.Notification) error
	Delete(id uint) error
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db: db}
}

func (r *notificationRepository) Create(notification *models.Notification) error {
	return r.db.Create(notification).Error
}

func (r *notificationRepository) GetByID(id uint) (*models.Notification, error) {
	var n models.Notification
	if err := r.db.First(&n, id).Error; err != nil {
		return nil, err
	}
	return &n, nil
}

func (r *notificationRepository) List(offset, limit int) ([]models.Notification, int64, error) {
	var notifications []models.Notification
	var count int64
	r.db.Model(&models.Notification{}).Count(&count)
	result := r.db.Order("created_at desc").Offset(offset).Limit(limit).Find(&notifications)
	return notifications, count, result.Error
}

func (r *notificationRepository) Update(notification *models.Notification) error {
	return r.db.Save(notification).Error
}

func (r *notificationRepository) Delete(id uint) error {
	return r.db.Delete(&models.Notification{}, id).Error
}
