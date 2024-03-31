package main

import (
	"log"

	"github.com/Kei-K23/go-otp/cmd/api"
	"github.com/Kei-K23/go-otp/internal/config"
	"github.com/Kei-K23/go-otp/internal/db"
)

func main() {
	sqlDB, err := db.NewDB(config.Env.DB_CONNECTION)

	if err != nil {
		log.Fatal(err)
	}

	db.InitDB(sqlDB)

	server := api.NewAPIServer(api.APIServer{Addr: ":8080", DB: sqlDB})

	server.Serve()
}
