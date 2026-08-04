package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	h := make(http.Header)
	if _, err := GetAPIKey(h); err != ErrNoAuthHeaderIncluded {
		t.Errorf("Expected %v but got %v", ErrNoAuthHeaderIncluded, err)
	}
	key := "abcdef"
	headerValue := "ApiKey " + key
	h.Set("Authorization", headerValue)
	if val, err := GetAPIKey(h); val != headerValue {
		t.Errorf("Expected %v but got %v; err=%v", key, val, err)
	}
}
