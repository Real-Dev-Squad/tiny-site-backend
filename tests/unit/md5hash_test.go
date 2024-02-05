package unit

import (
	"testing"

	"github.com/Real-Dev-Squad/tiny-site-backend/utils"
)

func TestGenerateMD5Hash(t *testing.T) {
	url := "https://example.com"
	hash := utils.GenerateMD5Hash(url)
	if len(hash) != 5 {
		t.Errorf("Expected length of 5, but got %d", len(hash))
	}
}

