package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

type Configuration struct {
	Url       string `json:"url,omitempty"`
	MediaPath string `json:"media,omitempty"`
	Jobs      []Job  `json:"jobs,omitempty"`
	Timezone  string `json:"tz,omitempty"`
}

func NewConfiguration(s string) (Configuration, error) {
	var c Configuration
	err := json.Unmarshal([]byte(s), &c)

	if err != nil {
		return c, err
	}

	for _, job := range c.Jobs {
		err := job.Validate()
		if err != nil {
			return c, fmt.Errorf("error processing cron schedule \"%v\": %v", job.Cron, err)
		}
	}

	return c, nil
}

// Get path to media file relative to MediaPath
func (c Configuration) MediaFilePath(mediaFile string) string {
	return path.Join(c.MediaPath, mediaFile)
}

func (c Configuration) TryString() (string, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c Configuration) JobsString() string {

	s := make([]string, len(c.Jobs))
	for i, job := range c.Jobs {
		s[i] = gray.Sprint("• ", job.Description())
	}
	return strings.Join(s, "\n")
}

type Job struct {
	Cron    string  `json:"cron"`
	Webhook Webhook `json:"webhook"`
}

func (j Job) Validate() error {
	_, err := cron.ParseStandard(j.Cron)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (j Job) Description() string {
	return fmt.Sprint(
		purp.Sprint(j.Cron),
		gray.Sprint(" - "),
		j.Webhook.Description(),
	)
}

func (j Job) CronFunction(c *Configuration) func() {

	return func() {
		res, err := j.Webhook.Post(c)
		if err != nil {
			printErr("Error in response:", err)
		}
		body := &bytes.Buffer{}
		_, err = body.ReadFrom(res.Body)
		if err != nil {
			printErr("Error parsing response:", err)
		}
		res.Body.Close()

		t := time.Now()

		timeFormatted := t.Format("01.02.06 03:04pm")

		gray.Printf(
			"%s %s %s\n",
			green.Sprint("✓"),
			gray.Sprint(timeFormatted),
			j.Description(),
		)
	}
}
