package service

import (
	"context"
	"net/http"
	"strconv"
	response "url-shortener/app/reponse"
	"url-shortener/config"
	"url-shortener/constant"
	"url-shortener/dto"
	"url-shortener/interfaces"
	"url-shortener/snowflakeIdGenerator"
)

type URLShortenerService struct {
	ctx context.Context
	repo interfaces.URLShortenerRepository
	snowflake *snowflakeIdGenerator.Snowflake
}

func NewURLShortenerService(repo interfaces.URLShortenerRepository) interfaces.URLShortenerService {
	return &URLShortenerService{
		ctx: context.Background(),
		repo: repo,
		snowflake: snowflakeIdGenerator.Generator,
	}
}

func (s *URLShortenerService) ShortenURL(r *http.Request, shortUrlRequest dto.ShortenURLRequest) (response.ApiResponse[dto.ShortenURLResponse]) {
	cfg := config.LoadEnvData()
	shortUrlDomain := cfg.SHORT_URL_DOMAIN
	id, _ := s.snowflake.Generate()
	shortUrl := shortUrlDomain + "short/" + strconv.FormatUint(id, 10)
	createdShortUrl, err := s.repo.ShortenURL(shortUrlRequest, shortUrl)
	if err != nil {
		return *response.StandardApiResponse(false, http.StatusBadRequest, constant.ShortUrlDoNotCreate, dto.ShortenURLResponse{})
	}
	return *response.StandardApiResponse(true, http.StatusCreated, constant.ShortURLCreatedSuccessfully, dto.ShortenURLResponse{ShortURL: createdShortUrl})
}

func (s *URLShortenerService) RedirectUrl(r *http.Request, shortURL string, userId int) (response.ApiResponse[dto.ShortenURLResponse]) {
	result, err := s.repo.RedirectUrl(shortURL, userId)
	if err != nil {
		return *response.StandardApiResponse(false, http.StatusNotFound, constant.URLNotFound, dto.ShortenURLResponse{})
	}
	return *response.StandardApiResponse(true, http.StatusOK, constant.URLFoundSuccessfully, result)
}