package main

import (
	"context"
	"go_discord_file_server/config"
	"go_discord_file_server/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"

	bot "go_discord_file_server/services"
)

func main() {
	config.LoadConfig()

	http.HandleFunc("/uploadImage", handlers.ImageUploadHandler)
	http.HandleFunc("/getImage", handlers.ImageGetHandler)
	http.HandleFunc("/deleteImage", handlers.ImageDeleteHandler)

	port := config.GetEnv("PORT", "8080")

	server := &http.Server{Addr: ":" + port}
	go func() {
		log.Printf("Starting server on port %s\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not start server: %s\n", err.Error())
		}
	}()
	defer server.Shutdown(context.Background())

	bot.BotToken = config.GetEnv("AUTH_TOKEN", "")
	bot.ChannelID = config.GetEnv("CHANNEL_ID", "")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Println("Shutting down the application...")
}
