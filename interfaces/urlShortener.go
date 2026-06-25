package interfaces

import (
	"net/http"
	response "url-shortener/app/reponse"
	"url-shortener/dto"
)

type URLShortenerRepository interface {
	ShortenURL(shortUrlRequest dto.ShortenURLRequest, shortUrl string) (string, error)
	RedirectUrl(shortURL string, userId int) (dto.ShortenURLResponse, error)
}

type URLShortenerService interface {
	ShortenURL(r *http.Request, shortUrlRequest dto.ShortenURLRequest) (response.ApiResponse[dto.ShortenURLResponse])
	RedirectUrl(r *http.Request, shortURL string, userId int) (response.ApiResponse[dto.ShortenURLResponse])
}

type URLShortenerHandler interface {
	ShortenURLHandler(w http.ResponseWriter, r *http.Request)
	RedirectUrlHandler(w http.ResponseWriter, r *http.Request)
}