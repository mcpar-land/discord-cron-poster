package main

import "testing"

func TestParseConfiguration(t *testing.T) {
	s := `{
		"url": "https://youtube.com/neat",
		"jobs": [
			{
				"cron": "@every 10s",
				"webhook": {
					"content": "This is the cool content"
				}
			}
		]
	}`
	config, err := NewConfiguration(s)
	if err != nil {
		t.Errorf("Error parsing: %v", err)
	} else {
		t.Logf("%#v", config)
	}
}

func TestBadCronDefinition(t *testing.T) {
	s := `{
		"url": "https://youtube.com/neat",
		"jobs": [
			{
				"cron": "WOO YEA WOO YEA",
				"webhook": {
					"content": "This is the cool content"
				}
			}
		]
	}`
	config, err := NewConfiguration(s)
	if err != nil {
		t.Logf("%#v", config)
	} else {
		t.Error("bad cron definition test should have thrown")
	}
}
