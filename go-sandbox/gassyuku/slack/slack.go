package slack

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type slack struct {
	params params
	url    string
}

type params struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	IconURL   string `json:"icon_url"`
	Channel   string `json:"channel"`
}

func NewSlack(url string, text string, username string, iconEmoji string, iconURL string, channel string) *slack {

	p := params{
		Text:      text,
		Username:  username,
		IconEmoji: iconEmoji,
		IconURL:   iconURL,
		Channel:   channel}

	return &slack{
		params: p,
		url:    url}
}

func (s *slack) Send() {

	params, _ := json.Marshal(s.params)

	resp, err := http.PostForm(
		s.url,
		url.Values{"payload": {string(params)}},
	)
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()
}
