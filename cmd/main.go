package main

import (
	"flag"

	"github.com/jgvkmea/alert-linebot/service"
)

var (
	userID       = flag.String("userid", "", "send to userid")
	alertMessage = flag.String("message", "", "send alert message")
)

func main() {
	flag.Parse()
	service.SendAlertByLine(*userID, *alertMessage)
}
