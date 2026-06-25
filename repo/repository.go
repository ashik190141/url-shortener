package repo

import (
	"context"
	"url-shortener/dto"
	"url-shortener/interfaces"
	"url-shortener/model"
	"gorm.io/gorm"
)

type URLShortenerRepository struct {
	db  *gorm.DB
	ctx context.Context
}

func NewURLShortenerRepository(db *gorm.DB) interfaces.URLShortenerRepository {
	return &URLShortenerRepository{
		db:  db,
		ctx: context.Background(),
	}
}

func (r *URLShortenerRepository) ShortenURL(shortUrlRequest dto.ShortenURLRequest, shortUrl string) (string, error) {
	req := model.URL{
		LongURL:  shortUrlRequest.LongURL,
		ShortURL: shortUrl,
		UserID:   shortUrlRequest.UserID,
	}
	result := r.db.Create(&req)
	if result.Error != nil {
		return "", result.Error
	}
	return shortUrl, nil
}

func (r *URLShortenerRepository) RedirectUrl(shortURL string, userId int) (dto.ShortenURLResponse, error) {
	var url model.URL
	result := r.db.Where("short_url = ? AND user_id = ?", shortURL, userId).First(&url)
	if result.Error != nil {
		return dto.ShortenURLResponse{}, result.Error
	}
	return dto.ShortenURLResponse{ID: int(url.ID), LongURL: url.LongURL, ShortURL: url.ShortURL, UserID: url.UserID, CreatedAt: url.CreatedAt.String(), UpdatedAt: url.UpdatedAt.String()}, nil
}