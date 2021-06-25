package models

type Message struct {
	Type string `json:"type"`
	ID   string `json:"id,omitempty"`
	Text string `json:"text"`
}
