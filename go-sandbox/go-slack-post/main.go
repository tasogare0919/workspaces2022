package main

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type payload struct {
	UserName string `json:"username"`
	Text     string `json:"text"`
}

func sendSlack(webhookUrl string) error {
	slackPayload, err := json.Marshal(
		payload{
			UserName: "TestEvent",
			Text:     "Hello, World!",
		},
	)
	if err != nil {
		return err
	}
	res, err := http.PostForm(
		webhookUrl,
		url.Values{"payload": {
			string(slackPayload),
		},
		})
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil

}

func main() {
	webhookUrl := ""
	sendSlack(webhookUrl)
}
