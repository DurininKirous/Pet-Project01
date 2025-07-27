package main

import (
	"project01/app/internal/db/connection"
	"project01/app/internal/db/config"
	"project01/app/internal/db/users"
	"project01/app/internal/server"
	"context"
	"net/http"
	"log"
)

func main () {
	cfg := config.Load()

	ctx := context.Background()
	db := connection.ConnectionStart(ctx, cfg.DB)
	repo := usersDBMethods.New(db)

	defer db.Close()	

	router := server.NewRouter(repo)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("unable to start server: %v", err)
	}

}
