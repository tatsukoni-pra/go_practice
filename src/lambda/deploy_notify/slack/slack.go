package slack

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
)

type slack struct {
	params params
	url    string
}

type params struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	Channel   string `json:"channel"`
}

func NewSlack(text string, userName string, iconEmoji string) *slack {
	p := params {
		Text:      text,
		Username:  userName,
		IconEmoji: iconEmoji,
		Channel:   os.Getenv("NOTIFY_CHANNEL")}

	return &slack {
		params: p,
		url:    os.Getenv("SLACK_HOCK_URL")}
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
