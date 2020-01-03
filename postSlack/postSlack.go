package postSlack

import (
	"bytes"
	"net/http"
)

type Slack struct {
	Username string
	Message  string
	Url      string
}

func (slack *Slack) SetUsername(username string) {
	slack.Username = username
}

func (slack *Slack) SetMessage(message string) {
	slack.Message = message
}

func (slack *Slack) SetUrl(url string) {
	slack.Url = url
}

func HttpPost(slackEntity *Slack) error {
	jsonStr := `"payload={"username": "` + slackEntity.Username + `","text": ` + slackEntity.Message + `}`
	url := slackEntity.Url
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
