intigriti-slack-announce
========================
Go bot that publishes (non-sensitive) new intigriti findings to Slack.

## Setup
1. Download [the latest isa release](https://github.com/hazcod/intigriti-slack-announce/releases).
2. Create a Slack app, create a [Slack incoming Webhook](https://api.slack.com/messaging/webhooks#getting_started) and add `chat:write` permissions.
3. Retrieve your [intigriti API token](https://intigriti.com/) and pass your (external) IP address for whitelisting.
4. Create your configuration file:
```yaml
# skip findings in audit, archived and closed
include_non_ready: false

# how often to check in minutes
check_interval_minutes: 15

# your slack webhook
slack_url: "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"

# your intigriti API credentials
intigriti_client_id: "XXXXXXXXXXX"
intigriti_client_secret: "XXXXXXXXXXX"
```
5. Run `isa` (preferably as a service) with arguments:
```shell
./isa -conf=my-conf.yaml
```
3. See new intigriti findings roll in on your Slack channel.
Any findings already sent to your Slack channel will be added to your YAML configuration file for portability.

## Building
This requires `make` and `go` to be installed.
Just run `make`.

## Customizing the Slack message
The [Slack BlockKit Builder](https://api.slack.com/tools/block-kit-builder?mode=message&blocks=%5B%7B%22type%22%3A%22section%22%2C%22text%22%3A%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22You%20have%20a%20new%20request%3A%5Cn*%3CfakeLink.toEmployeeProfile.com%7CFred%20Enriquez%20-%20New%20device%20request%3E*%22%7D%7D%2C%7B%22type%22%3A%22section%22%2C%22fields%22%3A%5B%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22*Type%3A*%5CnComputer%20(laptop)%22%7D%2C%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22*When%3A*%5CnSubmitted%20Aut%2010%22%7D%2C%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22*Last%20Update%3A*%5CnMar%2010%2C%202015%20(3%20years%2C%205%20months)%22%7D%2C%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22*Reason%3A*%5CnAll%20vowel%20keys%20aren%27t%20working.%22%7D%2C%7B%22type%22%3A%22mrkdwn%22%2C%22text%22%3A%22*Specs%3A*%5Cn%5C%22Cheetah%20Pro%2015%5C%22%20-%20Fast%2C%20really%20fast%5C%22%22%7D%5D%7D%2C%7B%22type%22%3A%22actions%22%2C%22elements%22%3A%5B%7B%22type%22%3A%22button%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22emoji%22%3Atrue%2C%22text%22%3A%22Approve%22%7D%2C%22style%22%3A%22primary%22%2C%22value%22%3A%22click_me_123%22%7D%2C%7B%22type%22%3A%22button%22%2C%22text%22%3A%7B%22type%22%3A%22plain_text%22%2C%22emoji%22%3Atrue%2C%22text%22%3A%22Deny%22%7D%2C%22style%22%3A%22danger%22%2C%22value%22%3A%22click_me_123%22%7D%5D%7D%5D)
can be used to replace [slack/message.go](slack/message.go).
