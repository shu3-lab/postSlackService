package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/shu3-lab/postSlackService/postSlack"
)

type JsonSchema struct {
	username string `json:"username"`
	message  string `json:"message"`
	url      string `json:"url"`
}

func main() {
	http.HandleFunc("/slack", slackHandler)
	http.ListenAndServe(":8080", nil)
}

func slackHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("io error")
		return
	}
	jsonBytes := ([]byte)(b)
	data := new(JsonSchema)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}
	slack := new(postSlack.Slack)
	slack.SetUsername(data.username)
	slack.SetMessage(data.message)
	slack.SetUrl(data.url)

	err = postSlack.HttpPost(slack)
	if err != nil {
		fmt.Printf("Error is occured!")
	} else {
		fmt.Printf("Success!")
	}
}
