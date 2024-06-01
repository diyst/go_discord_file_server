package services_test

import (
	"go_discord_file_server/config"
	"go_discord_file_server/services"
	"log"
	"os"
	"testing"
)

func TestUploadImage(t *testing.T) {
	// Test successful upload
	file, err := os.Open("testdata/testImage.png")
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	config.LoadConfig()

	services.BotToken = config.GetEnv("AUTH_TOKEN", "")
	services.ChannelID = config.GetEnv("CHANNEL_ID", "")

	resp, err := services.UploadImage(file, file.Name())
	if err != nil {
		t.Errorf("Failed to upload image: %v", err)
	}

	if resp.MsgId == "" {
		t.Errorf("Expected non-empty response, but got empty")
	}

	log.Println(resp)
}

func TestDeleteImage(t *testing.T) {
	// Test successful deletion
	config.LoadConfig()
	services.BotToken = config.GetEnv("AUTH_TOKEN", "")
	services.ChannelID = config.GetEnv("CHANNEL_ID", "")

	file, err := os.Open("testdata/testImage.png")
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	respo, err := services.UploadImage(file, file.Name())

	if err != nil {
		t.Errorf("Failed to upload image: %v", err)
	}

	err = services.DeleteImage(respo.MsgId)
	if err != nil {
		t.Errorf("Failed to delete image: %v", err)
	}
}
