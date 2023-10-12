package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	Id           int64     `bun:"id,pk,autoincrement"`
	Username     string    `bun:"username,notnull"`
	Email        string    `bun:"email,unique,notnull"`
	Password     string    `bun:"password"`
	IsVerified   bool      `bun:"is_verified,default:false"`
	IsOnboarding bool      `bun:"is_onboarding,default:true"`
	CreatedAt    time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt    time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}
