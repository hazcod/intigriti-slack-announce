package slack

import (
	"net/url"
)

type Endpoint struct {
	webHook		url.URL
	clientTag 	string
}

type Message struct {
	Text	string		`json:"text"`
}

func NewEndpoint(webHook url.URL, clientTag string) Endpoint {
	return Endpoint{
		webHook: webHook,
		clientTag: clientTag,
	}
}