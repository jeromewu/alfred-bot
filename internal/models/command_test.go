package models

import "testing"

func TestNewCommand(t *testing.T) {
	cases := map[string]string{
		"foo ":                     "",
		"":                         "",
		"af ":                      "  []",
		"af help":                  "help  []",
		"af todo":                  "todo  []",
		"af todo add":              "todo add []",
		"af todo add 1":            "todo add [1]",
		"af todo add buy the milk": "todo add [buy the milk]",
	}

	for k, v := range cases {
		cmdStr := NewCommand(k).String()
		if cmdStr != v {
			t.Errorf("Failed to parse command: `%s` != `%s`", cmdStr, v)
		}
	}
}
