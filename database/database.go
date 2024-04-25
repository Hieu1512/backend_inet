package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"main/model" // tên gói chứa file model
)

// DB là biến cơ sở dữ liệu mà các gói khác có thể truy cập
var DB *gorm.DB

// InitDB khởi tạo kết nối cơ sở dữ liệu và tự động migrate các bảng
func InitDB() {
	dsn := "yuno:yuno@tcp(127.0.0.1:3306)/demo1?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	// Tự động migrate các bảng
	err = DB.AutoMigrate(&model.Page2{})
	if err != nil {
		log.Fatalln(err)
	}
}
