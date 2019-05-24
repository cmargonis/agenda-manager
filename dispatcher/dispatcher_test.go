package dispatcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	userId = "sample_id"
	prefix = "<@" + userId + ">"
)

func TestMessageDirectedToUser(t *testing.T) {
	messageText := prefix + " Hello World"
	result := isMessageDirectedToUser(messageText, userId)

	assert.True(t, result, "Expected true, got false")
}

func TestMessageNotDirectedToUser(t *testing.T) {
	messageText := "Hello World"
	result := isMessageDirectedToUser(messageText, userId)

	assert.False(t, result, "Expected false, got true")
}
