package main

import (
	"log"

	"main/database"
	"main/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Kết nối cơ sở dữ liệu và tự động migrate các bảng
	database.InitDB()

	// Khởi tạo ứng dụng Fiber
	app := fiber.New()
	app.Use(cors.New())

	// Thiết lập các tuyến đường (routes)
	router.SetupRoutes(app)
	// Chạy ứng dụng trên cổng 3000
	log.Fatal(app.Listen(":3000"))
}
