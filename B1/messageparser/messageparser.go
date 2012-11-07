// Package parses messages.
package messageparser

import (
	"strings"
)

// Message maps a string key to a list of values.
type Message map[string][]string

// ParseMessage parses the Message string and returns a map listing the
// values specified for each key.
func ParseMessage(message string) (m Message) {
	m = make(Message)
	parseMessage(m, message)
	return
}

func parseMessage(m Message, message string) {
	array := strings.Split(message, " ")
	for _, value := range array {
		elem := strings.Split(value, "=")
		m[elem[0]] = append(m[elem[0]], elem[1])
	}
	return
}
