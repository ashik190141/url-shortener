package app

import (
	"net/http"
	"url-shortener/handler"
	"url-shortener/interfaces"
	"github.com/gorilla/mux"
)

func Route(repo interfaces.URLShortenerRepository, service interfaces.URLShortenerService) *mux.Router {
	router := mux.NewRouter()
	h := handler.NewUrlShortenerHandler(service, repo)

	router.HandleFunc("/urlShortener/shorten", h.ShortenURLHandler).Methods("POST")
	router.HandleFunc("/urlShortener/redirect", h.RedirectUrlHandler).Methods("GET")

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Server is running"))
	}).Methods("GET")

	return router
}