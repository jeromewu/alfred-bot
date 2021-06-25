package models

import (
	"io/ioutil"
	"log"
	"net/url"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	Token       string `yaml:"token"`
	RawReplyURL string `yaml:"reply_url"`
	ProjectID   string `yaml:"project_id"`
}

func NewConf(path string) *Conf {
	conf := new(Conf)
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read config.yaml: %v", err)
	}

	if err := yaml.Unmarshal(yamlFile, conf); err != nil {
		log.Fatalf("failed to parse config.yaml: %v", err)
	}

	return conf
}

func (c *Conf) ReplyURL() *url.URL {
	u, _ := url.Parse(c.RawReplyURL)
	return u
}
