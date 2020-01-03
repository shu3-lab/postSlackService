package postSlack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetUsername(t *testing.T) {
	slack := new(Slack)
	slack.SetUsername("user")
	result := slack.Username
	expected := "user"
	if result != expected {
		t.Error("\n結果:", result, "\n期待:", expected)
	}
	t.Log("終了")
}

func TestSetMessage(t *testing.T) {
	slack := new(Slack)
	slack.SetMessage("hello")
	assert.Equal(t, slack.Message, "hello")
}
