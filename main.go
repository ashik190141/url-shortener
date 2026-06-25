package main

import (
	"log"
	"net/http"
	"url-shortener/app/db"
	route "url-shortener/app/route"
	config "url-shortener/config"
	repo "url-shortener/repo"
	service "url-shortener/service"
	"url-shortener/snowflakeIdGenerator"
)

func main() {
	cfg := config.LoadEnvData()
	address := ":" + cfg.Port

	database:= db.ConnectPostgres()

	err := snowflakeIdGenerator.Init()
	if err != nil {
		panic(err)
	}

	log.Printf("HTTP Url shortener running on port %s", cfg.Port)

	urlShortenerRepo := repo.NewURLShortenerRepository(database)
	urlShorenerService := service.NewURLShortenerService(urlShortenerRepo)
	router := route.Route(urlShortenerRepo, urlShorenerService)

	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
	log.Printf("HTTP Url Shortener Service running on port %s", cfg.Port)
}