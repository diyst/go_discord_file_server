package services

import (
	"fmt"
	"log"
	"mime/multipart"

	"github.com/bwmarrin/discordgo"
)

var BotToken string
var ChannelID string

type UploadImageResponse struct {
	Url   string `json:"url"`
	MsgId string `json:"msgId"`
}

func UploadImage(file multipart.File, name string) (UploadImageResponse, error) {
	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		return UploadImageResponse{}, err
	}

	message := &discordgo.MessageSend{
		Files: []*discordgo.File{
			{
				Name:   name,
				Reader: file,
			},
		},
	}

	msgReturn, err := discord.ChannelMessageSendComplex(ChannelID, message)
	if err != nil {
		return UploadImageResponse{}, fmt.Errorf("fail to send file to discord: %v", err)
	}

	err = discord.Close()
	if err != nil {
		return UploadImageResponse{}, fmt.Errorf("fail to close discord session: %v", err)
	}

	return UploadImageResponse{
		Url:   msgReturn.Attachments[0].URL,
		MsgId: msgReturn.ID,
	}, nil
}

func DeleteImage(msgId string) error {
	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		return err
	}

	err = discord.ChannelMessageDelete(ChannelID, msgId)
	if err != nil {
		return fmt.Errorf("fail to delete message: %v", err)
	}

	err = discord.Close()
	if err != nil {
		return fmt.Errorf("fail to close discord session: %v", err)
	}

	return nil
}

func GetImage(msgId string) (string, error) {
	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		return "", err
	}

	message, err := discord.ChannelMessage(ChannelID, msgId)
	if err != nil {
		return "", fmt.Errorf("fail to get message: %v", err)
	}

	err = discord.Close()
	if err != nil {
		return "", fmt.Errorf("fail to close session: %v", err)
	}

	log.Println(message.Attachments[0].URL)

	return message.Attachments[0].URL, nil
}
