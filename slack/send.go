package slack

import (
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/hazcod/intigriti-slack-announce/intigriti"
)

func (s *Endpoint) Send(finding intigriti.Finding) []error {
	message := buildMessage(finding)
	return slack.Send(s.webHook.String(), "", message)
}

