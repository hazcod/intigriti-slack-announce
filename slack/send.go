package slack

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (s *Endpoint) Send(message Message) error {
	jsonBody, err := json.Marshal(&message)
	if err != nil {
		return errors.Wrap(err, "could not encode to json")
	}

	logrus.WithField("message", string(jsonBody)).Debugf("sending to slack")

	req, err := http.NewRequest(http.MethodPost, s.webHook.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return errors.Wrap(err, "could not create http request to slack")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Client", s.clientTag)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "message sending to slack failed")
	}

	_ = resp.Body.Close()

	if resp.StatusCode > 399 {
		return errors.Errorf("message sent to slack returned status code: %d", req.Response.StatusCode)
	}

	logrus.Info("message delivered to slack")
	return nil
}

