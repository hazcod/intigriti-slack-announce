package intigriti

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	expectedTokenType  = "bearer"
	expectedTokenScope = "external_api"
)

type authResponse struct {
	AccessToken			string	`json:"access_token"`
	ExpiresAtSeconds	int		`json:"expires_in"`
	TokenType			string 	`json:"token_type"`
	Scope 				string	`json:"scope"`
}

func (i *Endpoint) Authenticate() error {
	now := time.Now()
	if i.authTokenExp.Add( time.Second * 5 ).After( now) {
		logrus.WithField("auth_token_exp", i.authTokenExp).
			WithField("now", now).
			Debug("no need to refresh intigriti auth token")
		return nil
	}

	form := url.Values{}
	form.Add("grant_type", "client_credentials")
	form.Add("client_id", i.clientToken)
	form.Add("client_secret", i.clientSecret)

	req, err := http.NewRequest(http.MethodPost, intigritiAPIAuthURL, strings.NewReader(form.Encode()))
	if err != nil {
		return errors.Wrap(err, "could not create http request to intigriti")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-Client", i.clientTag)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "message sending to intigriti failed")
	}

	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return errors.Errorf("received error code: %d", resp.StatusCode)
	}

	var authResponse authResponse
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "could not read response")
	}

	if err := json.Unmarshal(respBytes, &authResponse); err != nil {
		logrus.Debugf("%+v", string(respBytes))
		return errors.Wrap(err, "could not decode intigriti auth response")
	}

	now = time.Now()
	newExpTime := now.Add(time.Second * time.Duration(authResponse.ExpiresAtSeconds))
	if newExpTime.Before(now) {
		return errors.Errorf("new expiration time %s is before %s", newExpTime, now)
	}

	i.authTokenExp = newExpTime
	i.authToken = authResponse.AccessToken

	if strings.ToLower(authResponse.TokenType) != expectedTokenType {
		logrus.WithField("token_type", authResponse.TokenType).Warn("unexpected token type")
	}

	if strings.ToLower(authResponse.Scope) != expectedTokenScope {
		logrus.WithField("token_scope", authResponse.Scope).Warn("unexpected token scope")
	}

	logrus.WithField("token_exp", newExpTime).Debug("new token expiration set")
	return nil
}
