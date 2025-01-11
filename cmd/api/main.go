package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mwombeki6/e_water-backend/config"
	"github.com/mwombeki6/e_water-backend/db"
	"github.com/mwombeki6/e_water-backend/handlers"
	//"github.com/mwombeki6/e_water-backend/middlewares"
	"github.com/mwombeki6/e_water-backend/repositories"
	"github.com/mwombeki6/e_water-backend/services"
)

func main()  {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName: "AxEnergies",
		ServerHeader: "Fiber",
	})

	//Repositories
	authRepository := repositories.NewAuthRepository(db)

	//Service
	authService := services.NewAuthService(authRepository)

	// Routing
	server := app.Group("/api")

	// Handlers
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	//privateRoutes := server.Use(middlewares.AuthProtected(db))

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))

}