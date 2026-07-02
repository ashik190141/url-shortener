package handler

import (
	"encoding/json"
	"net/http"
	"url-shortener/dto"
	"url-shortener/interfaces"
)

type UrlShortenerHandler struct {
	service interfaces.URLShortenerService
	repo    interfaces.URLShortenerRepository
}

func NewUrlShortenerHandler(service interfaces.URLShortenerService, repo interfaces.URLShortenerRepository) interfaces.URLShortenerHandler {
	return &UrlShortenerHandler{
		service: service,
		repo:    repo,
	}
}

func (u *UrlShortenerHandler) RedirectUrlHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		UserID int `json:"user_id"`
		ShortUrl string `json:"short_url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	response := u.service.RedirectUrl(r, body.ShortUrl, body.UserID)
	json.NewEncoder(w).Encode(response)
}

func (u *UrlShortenerHandler) ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.ShortenURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	response := u.service.ShortenURL(r, req)
	json.NewEncoder(w).Encode(response)
}