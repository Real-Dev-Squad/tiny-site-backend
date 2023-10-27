package unit

import (
	"testing"
	"time"

	"github.com/Real-Dev-Squad/tiny-site-backend/utils"
)

func TestGenerateMD5Hash(t *testing.T) {
	testCases := []struct {
		inputURL     string
		expectedHash string
	}{
		{
			inputURL:     "https://example.com",
			expectedHash: "a4433a82",
		},
		{
			inputURL:     "https://google.com",
			expectedHash: "afff4190",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.inputURL, func(t *testing.T) {
			actualHash := utils.GenerateMD5Hash(testCase.inputURL)
			if actualHash != testCase.expectedHash {
				t.Errorf("URL: %s, Expected Hash: %s, Actual Hash: %s", testCase.inputURL, testCase.expectedHash, actualHash)
			}
		})
	}
	urlWithTime := "https://example.com" + time.Now().Format("20060102150405")
	actualHash := utils.GenerateMD5Hash(urlWithTime)
	if actualHash == "" {
		t.Errorf("Hash for URL with time is empty")
	}
}
