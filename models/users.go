package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID           int64     `bun:"id,pk,autoincrement" json:"id"`
	UserName     string    `bun:"username,notnull" json:"userName"`
	Email        string    `bun:"email,unique,notnull" json:"email"`
	Password     string    `bun:"password" json:"-"`
	IsVerified   bool      `bun:"is_verified,default:false" json:"isVerified"`
	IsOnboarding bool      `bun:"is_onboarding,default:true" json:"isOnboarding"`
	CreatedAt    time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"createdAt"`
	UpdatedAt    time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updatedAt"`
}
