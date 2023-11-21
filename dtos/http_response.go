package dtos

import "time"

type URLCreationResponse struct {
	Message  string `json:"message"`
	ShortURL string `json:"shortUrl,omitempty"`
}

type UserURLsResponse struct {
	Message string       `json:"message"`
	URLs    []URLDetails `json:"urls,omitempty"`
}

type URLDetailsResponse struct {
	Message string       `json:"message"`
	URL     []URLDetails `json:"urls,omitempty"`
}

type URLDetails struct {
	OriginalURL string    `json:"originalUrl"`
	ShortURL    string    `json:"shortUrl"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
}
