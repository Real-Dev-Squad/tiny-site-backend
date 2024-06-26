package dtos

import "time"

type URLCreationResponse struct {
	Message   string    `json:"message"`
	ShortURL  string    `json:"shortUrl,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	URLCount  int       `json:"urlCount"`
}
type URLDeleteResponse struct{
	Message   string    `json:"message"`
	URLCount   int       `json:"urlCount"`
}
type UserURLsResponse struct {
	Message string       `json:"message"`
	URLs    []URLDetails `json:"urls,omitempty"`
}
type URLDetailsResponse struct {
	Message string     `json:"message"`
	URL     URLDetails `json:"url,omitempty"`
}
type URLDetails struct {
	ID             int64     `json:"id"`
	OriginalURL    string    `json:"originalUrl"`
	ShortURL       string    `json:"shortUrl"`
	Comment        string    `json:"comment,omitempty"`
	UserID         int64     `json:"userId"`
	CreatedBy      string    `json:"createdBy"`
	ExpiredAt      time.Time `json:"expiredAt,omitempty"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
	AccessCount    int64     `json:"accessCount"`
	LastAccessedAt time.Time `json:"lastAccessedAt"`
}

type URLNotFoundResponse struct {
	Message string `json:"message"`
}

type UserListResponse struct {
	Message string `json:"message"`
	Data    []User `json:"data,omitempty"`
}
type UserResponse struct {
	Message string `json:"message"`
	Data    User   `json:"data,omitempty"`
}

type User struct {
	ID        int64     `json:"id"`
	UserName  string    `json:"userName"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
