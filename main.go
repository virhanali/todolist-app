package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Routes(app fiber.Router) {
	r := app.Group("/todos")
	r.Post("/create", CreateTodo)
	r.Get("/", GetTodo)
	r.Put("/:id", UpdateTodo)
	r.Delete("/:id", DeleteTodo)
}

func main() {
	db, err = gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Todo{})

	app := fiber.New()
	Routes(app)
	app.Listen(":8080")

}
