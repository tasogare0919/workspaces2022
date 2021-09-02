package main

import (
	"context"
	"os"

	slack "./slack"
)

func HandleRequest(ctx context.Context) (string, error) {
	var slackURL = os.Getenv("slackURL")

	sl := slack.NewSlack(slackURL, "This is a line of test in a channel.\nAnd this is another line of text.", "user", ":100:", "", "#golang_test")
	sl.Send()
	return "", nil
}
