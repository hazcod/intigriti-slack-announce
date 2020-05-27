package intigriti

import (
	"time"
)

const (
	intigritiAPIAuthURL = "https://login.intigriti.com/connect/token"
	intigritiAPIFetchURL = "https://api.intigriti.com/external/submission"
)

type Endpoint struct {
	clientToken		string
	clientSecret 	string
	clientTag 		string

	authToken		string
	authTokenExp    time.Time
}

func NewEndpoint(clientToken string, clientSecret string, clientTag string) Endpoint {
	return Endpoint{
		clientToken: clientToken,
		clientSecret: clientSecret,
		clientTag: clientTag,
	}
}