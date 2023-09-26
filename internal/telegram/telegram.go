package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const (
	scheme string = "https"
	host   string = "api.telegram.org"
)

type Config interface {
	//GetChatID - return telegram chat id
	GetChatID() int64

	//GegToken - return telegram bot token
	GegToken() string
}

type telegram struct{ c Config }

func New(c Config) *telegram { return &telegram{c: c} }

func (t *telegram) SendMsg(msg string) error {
	u := url.URL{Scheme: scheme, Host: host, Path: fmt.Sprintf("bot%s/sendMessage", t.c.GegToken())}
	v := u.Query()
	v.Add("chat_id", strconv.FormatInt(t.c.GetChatID(), 10))
	v.Add("text", msg)
	u.RawQuery = v.Encode()
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}
	return do(req)
}

func (t *telegram) SendPhoto(photo string, caption string) error {
	file, err := os.Open(photo)
	if err != nil {
		return err
	}
	defer file.Close()
	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)
	writer.WriteField("chat_id", strconv.FormatInt(t.c.GetChatID(), 10))
	writer.WriteField("caption", caption)
	defer writer.Close()
	part, err := writer.CreateFormFile("photo", photo)
	if err != nil {
		return err
	}
	if _, err := io.Copy(part, file); err != nil {
		return err
	}
	writer.Close()
	u := url.URL{Scheme: scheme, Host: host, Path: fmt.Sprintf("bot%s/sendPhoto", t.c.GegToken())}
	req, err := http.NewRequest(http.MethodPost, u.String(), buffer)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return do(req)
}

func (t *telegram) SendDoc(doc string, caption string) error {
	file, err := os.Open(doc)
	if err != nil {
		return err
	}
	defer file.Close()
	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)
	writer.WriteField("chat_id", strconv.FormatInt(t.c.GetChatID(), 10))
	writer.WriteField("caption", caption)
	defer writer.Close()
	part, err := writer.CreateFormFile("document", doc)
	if err != nil {
		return err
	}
	if _, err := io.Copy(part, file); err != nil {
		return err
	}
	writer.Close()
	u := url.URL{Scheme: scheme, Host: host, Path: fmt.Sprintf("bot%s/sendDocument", t.c.GegToken())}
	req, err := http.NewRequest(http.MethodPost, u.String(), buffer)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return do(req)
}

func do(req *http.Request) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	resp.Close = true
	defer resp.Body.Close()
	type rterr struct {
		Ok          bool   `json:"ok,omitempty"`
		ErrorCode   int64  `json:"error_code,omitempty"`
		Description string `json:"description,omitempty"`
	}
	ert := rterr{}
	if err := json.NewDecoder(resp.Body).Decode(&ert); err != nil {
		return err
	}
	if !ert.Ok {
		return fmt.Errorf("something went wrong, response code %d, description %s", ert.ErrorCode, ert.Description)
	}
	return nil
}
