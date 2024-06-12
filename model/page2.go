package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type CustomModel struct {
	Id       uint `gorm:"primaryKey" json:"id"`
	CreateAt time.Time
	UpdateAt *time.Time // Chú ý UpdateAt là con trỏ *time.Time
	DeleteAt gorm.DeletedAt
}

// Trước khi tạo mới
func (cm *CustomModel) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	cm.CreateAt = now
	// cm.UpdateAt = &now // Gán giá trị thời gian hiện tại cho UpdateAt
	return nil
}

// Trước khi cập nhật
func (cm *CustomModel) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	cm.UpdateAt = &now // Cập nhật giá trị thời gian hiện tại cho UpdateAt
	return nil
}

type Page2 struct {
	CustomModel
	Label    string
	Key      string
	Link     string
	Metadata datatypes.JSON
}
