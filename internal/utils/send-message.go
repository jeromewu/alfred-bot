package utils

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"cloud.google.com/go/firestore"
	m "github.com/jeromewu/alfred-bot/internal/models"
)

func SendMessage(url *url.URL, token string, body []byte) error {
	req := &http.Request{
		Method: "POST",
		URL:    url,
		Header: map[string][]string{
			"Content-Type":  {"application/json"},
			"Authorization": {"Bearer " + token},
		},
		Body: ioutil.NopCloser(bytes.NewReader(body)),
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	respBody, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	log.Println("Reply response body", string(respBody))

	return nil
}

func NewFirestoreClient(conf m.Conf) (*firestore.Client, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, conf.ProjectID)
	if err != nil {
		return nil, err
	}
	return client, nil
}
