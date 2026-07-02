package dto

type ShortenURLRequest struct {
	LongURL string `json:"long_url" binding:"required,url"`
	UserID  int    `json:"user_id,omitempty"`
}

type ShortenURLResponse struct {
	ID        int    `json:"id"`
	LongURL   string `json:"long_url"`
	ShortURL  string `json:"short_url"`
	UserID    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}