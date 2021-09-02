package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"./slack"
	"github.com/aws/aws-lambda-go/events"

	"github.com/aws/aws-lambda-go/lambda"
)

type Message struct {
	AlarmName        string `json:"AlarmName"`
	AlarmDescription string `json:"AlarmDescription"`
}

func HandleRequest(ctx context.Context, snsEvent events.SNSEvent) (string, error) {
	var slackURL = os.Getenv("slackURL")
	if len(snsEvent.Records) == 0 {
		return "error", errors.New("SNSRecordがない")
	}

	snsRecord := snsEvent.Records[0].SNS

	message := new(Message)
	jsonBytes := []byte(snsRecord.Message)

	if err := json.Unmarshal(jsonBytes, message); err != nil {
		log.Fatal(err)
	}

	postMessage := fmt.Sprintf(
		"[アラート発報中]%s %s",
		message.AlarmName,
		message.AlarmDescription,
	)

	sl := slack.NewSlack(slackURL, postMessage, "EC2alert", ":rotating_light:", "", "#golang_test")
	sl.Send()

	return "", nil
}

func main() {
	lambda.Start(HandleRequest)
}
