package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	m "github.com/jeromewu/alfred-bot/internal/models"
	u "github.com/jeromewu/alfred-bot/internal/utils"
)

func raise(w http.ResponseWriter, prefix string, err error) {
	msg := fmt.Sprintf("%s: %v", prefix, err)
	log.Println(msg)
	http.Error(w, msg, http.StatusBadRequest)
}

func post(w http.ResponseWriter, r *http.Request, conf *m.Conf) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		raise(w, "Error reading body", err)
		return
	}

	_, err = base64.StdEncoding.DecodeString(r.Header.Get("x-line-signature"))
	if err != nil {
		raise(w, "Error decoding signature", err)
		return
	}

	// hash := hmac.New(sha256.New, []byte(conf.Token))
	// hash.Write(body)
	// if !hmac.Equal(decoded, hash.Sum(nil)) {
	//   log.Printf("Error matching signature")
	//   http.Error(w, "can't match signature", http.StatusBadRequest)
	//   return
	// }

	log.Println("Request body: ", string(body))

	msg := new(m.ReceivedMessage)
	if err := json.Unmarshal(body, msg); err != nil {
		raise(w, "Error parsing body", err)
		return
	}

	if cmd := m.NewCommand(msg.MessageText()); cmd != nil {
		ret, err := cmd.Execute(conf, msg)
		if err != nil {
			raise(w, "Error executing command", err)
		}
		if err := u.SendMessage(
			conf.ReplyURL(),
			conf.Token,
			m.NewReplyMessage(msg.ReplyToken(), ret).JSON(),
		); err != nil {
			raise(w, "Error replying message", err)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func Webhook(conf *m.Conf) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			post(w, r, conf)
		} else {
			http.Error(w, "405 METHOD NOT ALLOWED", http.StatusMethodNotAllowed)
		}
	}
}
