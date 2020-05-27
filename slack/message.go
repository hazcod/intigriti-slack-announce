package slack

import (
	"fmt"
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/hazcod/intigriti-slack-announce/intigriti"
	"regexp"
)

var (
	regexRemove = regexp.MustCompile("[\\W_]+")
)

// remove all chars except alphanumeric
func cleanStr(str string) string {
	return string(regexRemove.ReplaceAll([]byte(str), []byte("")))
}

func buildMessage(f intigriti.Finding) slack.Payload {
	attach := slack.Attachment{}

	/*
	attach.AddField(slack.Field{
		Title: "Created",
		Value: f.Timestamp.Format("2006-01-02 15:04:05"),
	})
	*/

	attach.AddField(slack.Field{
		Title: "Severity",
		Value: f.Severity,
	})
	attach.AddField(slack.Field{
		Title: "Type",
		Value: f.Type,
	})

	if f.Endpoint != "" {
		attach.AddField(slack.Field{
			Title: "Endpoint",
			Value: f.Endpoint,
		})
	}

	return slack.Payload{
		Username:    "intigriti",
		IconUrl: "https://www.intigriti.com/assets/img/intigriti-kumkn.png",
		IconEmoji: ":ghost:",
		Text:        fmt.Sprintf("A new finding was published by *%s* for *%s*: <%s|%s>",
			cleanStr(f.Researcher), cleanStr(f.Program), f.URL, cleanStr(f.Title)),
		Attachments: []slack.Attachment{attach},
	}
}