package main

import (
	"log"

	"github.com/EmiliodDev/GinTasker/internal/config"
	"github.com/EmiliodDev/GinTasker/internal/models"
	"github.com/EmiliodDev/GinTasker/internal/routes"
	"github.com/EmiliodDev/GinTasker/pkg/database"
)

func main() {
    config.LoadConfig()

    db, err := database.Connect()
    if err != nil {
        log.Fatalf("DB connection failed, %v", err)
    }

    if err := models.Migrate(db); err != nil {
        log.Fatalf("Error running migrations, %v", err)
    }

    r := routes.SetupRouter(db)

    log.Printf("Server running on port %s", config.AppConfig.Port)
    if err := r.Run(":" + config.AppConfig.Port); err != nil {
        log.Fatalf("Failed to start server, %v", err)
    }
}
