package main

type Embed struct {
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Url         string   `json:"url,omitempty"`
	Timestamp   string   `json:"timestamp,omitempty"`
	Color       int      `json:"color,omitempty"`
	Footer      Footer   `json:"footer,omitempty"`
	Image       Media    `json:"image,omitempty"`
	Thumbnail   Media    `json:"thumbnail,omitempty"`
	Video       Media    `json:"video,omitempty"`
	Provider    Provider `json:"provider,omitempty"`
	Author      Author   `json:"author,omitempty"`
	Fields      []Field  `json:"fields,omitempty"`
}

type Footer struct {
	Text         string `json:"text"`
	IconUrl      string `json:"icon_url,omitempty"`
	ProxyIconUrl string `json:"proxy_icon_url,omitempty"`
}

type Media struct {
	Url      string `json:"url,omitempty"`
	ProxyUrl string `json:"proxy_url,omitempty"`
	Height   string `json:"height,omitempty"`
	Width    string `json:"width,omitempty"`
}

type Provider struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type Author struct {
	Name         string `json:"name,omitempty"`
	Url          string `json:"url,omitempty"`
	IconUrl      string `json:"icon_url,omitempty"`
	ProxyIconUrl string `json:"proxy_icon_url,omitempty"`
}

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}
