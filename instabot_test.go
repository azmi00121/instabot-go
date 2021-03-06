package instabot

import "testing"
import "fmt"

func TestGetConfig(t *testing.T) {
	_, config := GetConfig()
	fmt.Println(config)
	if config.Username == "" {
		t.Error("Username is empty")
	}
	if config.Username != "" {
		t.Log("Username is not empty")
	}

	if config.Password == "" {
		t.Error("Password is empty")
	}
	if config.Password != "" {
		t.Log("Password is not empty")
	}
	if config.Tags != nil {
		t.Log("Tags not filled up")
	}

	if config.Tags == nil {
		t.Error("Tags empty")
	}
}
