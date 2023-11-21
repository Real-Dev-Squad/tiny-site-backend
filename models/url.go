package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Tinyurl struct {
	bun.BaseModel `bun:"table:tiny_url"`

	ID          int64     `bun:"id,pk,autoincrement" json:"id"`
	OriginalUrl string    `bun:"original_url,notnull" json:"originalUrl"`
	ShortUrl    string    `bun:"short_url,unique,notnull" json:"shortUrl"`
	Comment     string    `bun:"comment" json:"comment"`
	UserID      int64     `bun:"user_id,rel:belongs-to,join:id" json:"userId"`
	ExpiredAt   time.Time `bun:"expired_at,notnull" json:"expiredAt"`
	CreatedAt   time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"createdAt"`
	CreatedBy   string    `bun:"created_by,notnull" json:"createdBy"`
}
