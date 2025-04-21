package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tranvinh21/fastext-be-go/cmd/services/auth"
	"github.com/tranvinh21/fastext-be-go/cmd/services/user"
	"github.com/tranvinh21/fastext-be-go/config"
	"gorm.io/gorm"
)

type APIServer struct {
	address string
	db      *gorm.DB
}

func NewAPIServer(db *gorm.DB, port string) *APIServer {
	return &APIServer{
		address: port,
		db:      db,
	}
}

func (s *APIServer) Run() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.Envs.CORS.WHITELIST_DOMAINS,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowCredentials: true,
	}))
	app.Get("/check", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	apiRoutes := app.Group("/api")
	userHandler := user.SetupUserRoutes(user.NewUserStore(s.db))
	userHandler.RegisterRoutes(apiRoutes)

	authHandler := auth.SetupAuthRoutes(auth.NewAuthStore(s.db))
	authHandler.RegisterRoutes(apiRoutes)

	err := app.Listen(":" + s.address)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
