package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	header := make(http.Header)
	header.Add("Authorization", "ApiKey 1234adsf")

	key, err := GetAPIKey(header)
	if err != nil {
		t.Error(err)
	}

	if key != "1234adsf" {
		t.Error("key did not match")
	}

}
