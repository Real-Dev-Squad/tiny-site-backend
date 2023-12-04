package dtos

type URLCreationRequest struct {
	OriginalURL string `json:"original_url"`
	UserID      int64  `json:"user_id"`
	CreatedBy   string `json:"created_by,omitempty"`
	Comment     string `json:"comment,omitempty"`
}

type URLRedirectionRequest struct {
	ShortURL string `json:"short_url"`
}

type UserURLsFetchRequest struct {
	UserID string `uri:"id" binding:"required"`
}

type URLDetailsFetchRequest struct {
	ShortURL string `uri:"shortURL" binding:"required"`
}

type UserByIDFetchRequest struct {
	ID string `uri:"id" binding:"required"`
}
