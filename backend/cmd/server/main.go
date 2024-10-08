package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/mohithchintu/College_Book/backend/config"
	"github.com/mohithchintu/College_Book/backend/db"
	"github.com/mohithchintu/College_Book/backend/routes"
)

func main() {
	config.LoadConfig()
	db.MongoConnect()
	app := fiber.New()
	routes.UserRoutes(app)
	routes.MaterialRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, MongoDB!")
	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := app.Listen("0.0.0.0:" + config.AppConfig.Port); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	<-quit

	log.Println("Shutting down server...")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}
	db.MongoDisconnect()

	log.Println("Server shutdown completed.")
}
