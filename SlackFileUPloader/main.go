package main

import (
	"fmt"
	"os"

	"github.com/Edu58/GoSlackFileUploader/utils"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	err := godotenv.Load(".env")
	utils.ErrorHandler(err)

	api := slack.New(os.Getenv("BOT_OAUTH_TOKEN"), slack.OptionDebug(true))
	channels := []string{os.Getenv("CHANNEL_ID")}
	files := []string{"./assets/files/verify_phone.txt", "./assets/files/signature.png"}

	for _, file := range files {
		params := slack.FileUploadParameters{
			Channels: channels,
			File:     file,
		}

		file, err := api.UploadFile(params)
		utils.ErrorHandler(err)

		fmt.Printf("File %s uploaded to: %s", file.Name, file.Permalink)
	}
}