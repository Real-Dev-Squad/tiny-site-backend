package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Tinyurl struct {
	bun.BaseModel `bun:"table:tiny_url"`

	ID             int64     `bun:"id,pk,autoincrement" json:"id"`
	OriginalUrl    string    `bun:"original_url,notnull" json:"originalUrl"`
	ShortUrl       string    `bun:"short_url,unique,notnull" json:"shortUrl"`
	Comment        string    `bun:"comment" json:"comment"`
	UserID         int64     `bun:"user_id"`
	User           *User     `bun:"rel:belongs-to,join:user_id=id"`
	ExpiredAt      time.Time `bun:"expired_at,notnull" json:"expiredAt"`
	CreatedAt      time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"createdAt"`
	CreatedBy      string    `bun:"created_by,notnull" json:"createdBy"`
	AccessCount    int64     `bun:"access_count,default:0" json:"accessCount"`
	LastAccessedAt time.Time `bun:"last_accessed_at,nullzero" json:"lastAccessedAt"`
	IsDeleted      bool      `bun:"is_deleted" json:"isDeleted"`
	DeletedAt      time.Time `bun:"deleted_at" json:"deletedAt"`
}
