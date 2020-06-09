package slack

import (
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/hazcod/go-intigriti"
)

func (s *Endpoint) Send(finding intigriti.Submission) []error {
	message := buildMessage(finding)
	return slack.Send(s.webHook.String(), "", message)
}

