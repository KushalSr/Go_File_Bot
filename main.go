package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}

	os.Setenv("SLACK_BOT_TOKEN", "SLACK_BOT_TOKEN")
	os.Setenv("CHANNEL_ID", "CHANNEL_ID")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"config/app.scss"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf(`%s\n`, err)
			return
		}
		fmt.Printf("Name : %s, URL : %s\n", file.Name, file.URL)
	}

}
