package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Tinyurl struct {
	bun.BaseModel `bun:"table:tiny_url"`

	Id        int64     `bun:"id,pk,autoincrement"`
	OrgUrl    string    `bun:"org_url,notnull"`
	ShortUrl  string    `bun:"short_url,unique,notnull"`
	Comment   string    `bun:"comment"`
	UserId    int       `bun:"user_id"`
	ExpiredAt time.Time `bun:"expired_at,notnull"`
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	CreatedBy string    `bun:"created_by,notnull"`
}
