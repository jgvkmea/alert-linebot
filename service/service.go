package service

import (
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/sirupsen/logrus"
)

func SendAlertByLine(userID string, message string) error {
	log := logrus.New()
	log.Infoln("Receive message: ", message)

	client, err := linebot.New(
		os.Getenv("ALERT_LINEBOT_CHANNEL_SECRET"),
		os.Getenv("ALERT_LINEBOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Errorln("failed to create linebot client: ", err)
		return err
	}

	_, err = client.PushMessage(userID, linebot.NewTextMessage(message)).Do()
	if err != nil {
		log.Errorln("failed to push message: ", err)
		return err
	}
	log.Infoln("succeed to push message")
	return nil
}
