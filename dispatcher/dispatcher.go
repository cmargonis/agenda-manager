/*
Package dispatcher encompases the controller logic where incoming messages
get parsed to determine whether the program is being instructed to perform
a known operation.
*/
package dispatcher

import (
	"fmt"
	"strings"
)

// CheckForCommand receives a text and determines if it can start an operation
// based on the contents of the text
func CheckForCommand(text string, user string, channel string, currentUserId string) {

}

func isMessageDirectedToUser(text string, userId string) bool {
	prefix := fmt.Sprintf("<@%s> ", userId)
	return strings.HasPrefix(text, prefix)
}
