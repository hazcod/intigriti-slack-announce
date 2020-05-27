package slack

import (
	"fmt"
	"github.com/hazcod/intigriti-slack-announce/intigriti"
)

func buildMessage(f intigriti.Finding) string {
	return fmt.Sprintf(`
{
	"blocks": [
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "A new *%s* finding was published:\n*<%s|%s>*"
			}
		},
        {
         	"type": "context",
            "elements": [
                {
					"type": "mrkdwn",
					"text": "*Timestamp:*\n %s"
				},
                {
					"type": "mrkdwn",
					"text": "*Severity:*\n *%s*"
				} 
            ]
        },
		{
			"type": "context",
			"elements": [
				{
					"type": "mrkdwn",
					"text": "*Researcher:*\n %s"
				},
                {
					"type": "mrkdwn",
					"text": "*Endpoint:*\n%s"
				}
			]
		}
	]
}
`, f.Severity, f.URL, f.Title,
	f.Timestamp,
	f.Severity,
	f.Researcher,
	f.Endpoint)
}
