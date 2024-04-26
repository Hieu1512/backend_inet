package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type CustomModel struct {
	Id       uint `gorm:"primaryKey" json:"id"`
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt gorm.DeletedAt
}
type Page2 struct {
	CustomModel
	Label    string
	Key      string
	Link     string
	Metadata datatypes.JSON
}
