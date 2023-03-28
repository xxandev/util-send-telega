package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"util-send-telega/internal/utils"

	"gopkg.in/yaml.v2"
)

type FormatConfig int8

const (
	JSON FormatConfig = iota
	XML
	YAML
)

type Config struct {
	Telegram ConfigTelegram `json:"telegram,omitempty" xml:"telegram,omitempty" yaml:"telegram,omitempty"`
}

func (c *Config) Init() (bool, error) {
	switch {
	case utils.IsStatFile("send-telegram.json"):
		content, err := ioutil.ReadFile("send-telegram.json")
		if err != nil {
			return true, err
		}
		if err := c.Unmarshal(JSON, content); err != nil {
			return true, err
		}
	case utils.IsStatFile("send-telegram.xml"):
		content, err := ioutil.ReadFile("send-telegram.xml")
		if err != nil {
			return true, err
		}
		if err := c.Unmarshal(XML, content); err != nil {
			return true, err
		}
	case utils.IsStatFile("send-telegram.yaml"):
		content, err := ioutil.ReadFile("send-telegram.yaml")
		if err != nil {
			return true, err
		}
		if err := c.Unmarshal(YAML, content); err != nil {
			return true, err
		}
	}
	return false, errors.New("config file not found")
}

func (c *Config) Marshal(format FormatConfig) ([]byte, error) {
	switch format {
	case JSON:
		return json.MarshalIndent(c, "", "\t")
	case XML:
		return xml.MarshalIndent(c, "", "\t")
	case YAML:
		return yaml.Marshal(c)
	}
	return nil, errors.New("unknown configuration format")
}

func (c *Config) Unmarshal(format FormatConfig, data []byte) error {
	switch format {
	case JSON:
		return json.Unmarshal(data, c)
	case XML:
		return xml.Unmarshal(data, c)
	case YAML:
		return yaml.Unmarshal(data, c)
	}
	return errors.New("unknown configuration format")
}

func (c *Config) Example() {
	c.Telegram.ChatID = 123456789
	c.Telegram.Token = "987654321:xxxxxxxxx-xxx-xxxxxxxxxxxx"
}

func (c *Config) Check() error {
	if err := c.Telegram.Check(); err != nil {
		return fmt.Errorf("telegram: %v", err)
	}
	return nil
}

type ConfigTelegram struct {
	ChatID int64  `json:"chat_id,omitempty" xml:"chat_id,omitempty" yaml:"chat_id,omitempty"`
	Token  string `json:"bot_token,omitempty" xml:"bot_token,omitempty" yaml:"bot_token,omitempty"`
}

func (tlg *ConfigTelegram) Set(chatID int64, token string) {
	if chatID != 0 {
		tlg.ChatID = chatID
	}
	if len(token) > 0 {
		tlg.Token = token
	}
}

func (tlg *ConfigTelegram) GetChatID() int64 { return tlg.ChatID }

func (tlg *ConfigTelegram) GegToken() string { return tlg.Token }

func (tlg *ConfigTelegram) Check() error {
	if tlg.ChatID == 0 {
		return errors.New("chat id can't be 0")
	}
	if tlg.Token == "" {
		return errors.New("bot token can't be empty")
	}
	return nil
}
