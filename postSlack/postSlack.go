package postSlack

import (
	"bytes"
	"fmt"
	"io/ioutil"
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
	jsonStr := `"payload={"username": "` + slackEntity.Username + `","text": "` + slackEntity.Message + `"}"`
	println(jsonStr)
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
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, bErr := ioutil.ReadAll(res.Body)
	if bErr != nil {
		fmt.Println(bErr)
	}
	fmt.Println("Slack's response is " + string(body))
	return err
}
