package models

import "encoding/json"

type ReplyMessage struct {
	ReplyToken string    `json:"replyToken"`
	Messages   []Message `json:"messages"`
}

func NewReplyMessage(replyToken, text string) *ReplyMessage {
	return &ReplyMessage{
		ReplyToken: replyToken,
		Messages: []Message{
			{
				Type: "text",
				Text: text,
			},
		},
	}
}

func (m *ReplyMessage) JSON() []byte {
	buf, _ := json.Marshal(m)
	return buf
}
