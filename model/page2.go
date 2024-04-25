package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Page2 struct {
	gorm.Model
	Label    string
	Key      string
	Link     string
	Metadata datatypes.JSON
}
