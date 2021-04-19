package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Webhook struct {
	Content   string  `json:"content,omitempty"`
	Username  string  `json:"username,omitempty"`
	AvatarUrl string  `json:"avatar_url,omitempty"`
	Tts       bool    `json:"tts,omitempty"`
	File      string  `json:"file,omitempty"`
	Embeds    []Embed `json:"embeds,omitempty"`
}

func (w Webhook) Description() string {
	var s []string

	if w.Content != "" {
		s = append(s, yellow.Sprint(truncateString(w.Content, 40)))
	}
	if w.File != "" {
		s = append(s, blue.Sprint(w.File))
	}
	if len(w.Embeds) > 0 {
		es := fmt.Sprint(len(w.Embeds), " embed")
		if len(w.Embeds) > 1 {
			es += "s"
		}
		s = append(s, cyan.Sprint(es))
	}

	return strings.Join(s, gray.Sprint(", "))
}

func (w Webhook) Post(c *Configuration) (*http.Response, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if w.File != "" {
		path := c.MediaFilePath(w.File)
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		part, err := writer.CreateFormFile("file", filepath.Base(path))
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, file)
		if err != nil {
			return nil, err
		}
	}
	payloadJson, err := json.Marshal(w)
	if err != nil {
		return nil, err
	}
	writer.WriteField("payload_json", string(payloadJson))
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.Url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := http.DefaultClient.Do(req)
	return res, err
}

func truncateString(str string, num int) string {
	bnoden := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		bnoden = str[0:num] + "..."
	}
	return bnoden
}
