package main

import (
	"log"

	"github.com/EmiliodDev/GinTasker/internal/config"
	"github.com/EmiliodDev/GinTasker/pkg/database"
)

func main() {
    config.LoadConfig()

    db, err := database.Connect()
    if err != nil {
        log.Fatalf("DB connection failed, %v", err)
    }
    defer db.Close()

    r := 

}
