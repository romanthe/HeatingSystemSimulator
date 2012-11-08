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

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (v Message) Get(key string) string {
	if v == nil {
		return ""
	}
	vs, ok := v[key]
	if !ok || len(vs) == 0 {
		return ""
	}
	return vs[0]
}

// Set sets the key to value. It replaces any existing
// values.
func (v Message) Set(key, value string) {
	v[key] = []string{value}
}
