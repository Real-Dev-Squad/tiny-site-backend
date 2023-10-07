package controllers

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSignUpUser(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SignUpUser(tt.args.c)
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SignInUser(tt.args.c)
		})
	}
}
