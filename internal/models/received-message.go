package models

type ReceivedMessage struct {
	Destination string `json:"destination"`
	Events      []struct {
		Type      string  `json:"type"`
		Message   Message `json:"message"`
		Timestamp int64   `json:"timestamp"`
		Source    struct {
			Type    string `json:"type"`
			GroupID string `json:"groupId,omitempty"`
			UserID  string `json:"userId"`
		} `json:"source"`
		ReplyToken string `json:"replyToken"`
		Mode       string `json:"mode"`
	} `json:"events"`
}

func (m *ReceivedMessage) ReplyToken() string {
	if len(m.Events) != 0 {
		return m.Events[0].ReplyToken
	}
	return ""
}

func (m *ReceivedMessage) MessageText() string {
	if len(m.Events) != 0 {
		return m.Events[0].Message.Text
	}
	return ""
}

func (m *ReceivedMessage) UserID() string {
	if len(m.Events) != 0 {
		return m.Events[0].Source.UserID
	}
	return ""
}
