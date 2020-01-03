package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/shu3-lab/postSlackService/postSlack"
)

type JsonSchema struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Url      string `json:"url"`
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
	slack.SetUsername(data.Username)
	slack.SetMessage(data.Message)
	slack.SetUrl(data.Url)
	err = postSlack.HttpPost(slack)
	if err != nil {
		fmt.Println("Error is occured!", err)
	} else {
		fmt.Printf("Success!")
	}
}
