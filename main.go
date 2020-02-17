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
		responseWithError(w, http.StatusBadRequest, "Invalid request.")
	} else {
		responseWithSucess(w)
	}
}

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func responseWithError(w http.ResponseWriter, code int, msg string) {
	responseWithJson(w, code, map[string]string{"error": msg})
}

func responseWithSucess(w http.ResponseWriter) {
	responseWithJson(w, http.StatusOK, nil)
}
