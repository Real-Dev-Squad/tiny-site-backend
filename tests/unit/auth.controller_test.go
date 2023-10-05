package controllers

import (
	"testing"

	"tiny-site-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func TestSignUpUser(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := controllers.SignUpUser(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SignUpUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSignInUser(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := controllers.SignInUser(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SignInUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
