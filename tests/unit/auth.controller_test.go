package controllers

import (
	"testing"

	"tiny-site-backend/controllers"

	"github.com/gin-gonic/gin"
)

func TestSignUpUser(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controllers.SignUpUser(tt.args.c)
		})
	}
}

func TestSignInUser(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controllers.SignInUser(tt.args.c)
		})
	}
}