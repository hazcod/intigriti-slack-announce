package intigriti

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

type fetchResponse []struct {
	Code              string `json:"code"`
	InternalReference struct {
		Reference string `json:"reference"`
		URL       string `json:"url"`
	} `json:"internalReference"`
	Title         string `json:"title"`
	OriginalTitle string `json:"originalTitle"`
	Program       struct {
		Name                 string      `json:"name"`
		Handle               string      `json:"handle"`
		LogoURL              interface{} `json:"logoUrl"`
		ConfidentialityLevel struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"confidentialityLevel"`
		Status struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"status"`
	} `json:"program"`
	Type struct {
		Name     string      `json:"name"`
		Category interface{} `json:"category"`
		Cwe      interface{} `json:"cwe"`
	} `json:"type"`
	Severity struct {
		ID     int         `json:"id"`
		Vector interface{} `json:"vector"`
		Value  string      `json:"value"`
	} `json:"severity"`
	Domain struct {
		Type struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"type"`
		Value          string `json:"value"`
		Motivation     string `json:"motivation"`
		BusinessImpact struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"businessImpact"`
	} `json:"domain"`
	EndpointVulnerableComponent string `json:"endpointVulnerableComponent"`
	State                       struct {
		Status struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"status"`
		CloseReason struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"closeReason"`
	} `json:"state"`
	TotalPayout      float64     `json:"totalPayout"`
	CreatedAt        int         `json:"createdAt"`
	LastUpdatedAt    int         `json:"lastUpdatedAt"`
	ValidatedAt      int         `json:"validatedAt"`
	AcceptedAt       int         `json:"acceptedAt"`
	ClosedAt         int         `json:"closedAt"`
	ArchivedAt       interface{} `json:"archivedAt"`
	AwaitingFeedback bool        `json:"awaitingFeedback"`
	Assignee         struct {
		UserName  string      `json:"userName"`
		AvatarURL interface{} `json:"avatarUrl"`
		Email     string      `json:"email"`
	} `json:"assignee"`
	Researcher struct {
		UserName  string      `json:"userName"`
		AvatarURL interface{} `json:"avatarUrl"`
		Ranking   struct {
			Rank       int `json:"rank"`
			Reputation int `json:"reputation"`
			Streak     struct {
				ID    int    `json:"id"`
				Value string `json:"value"`
			} `json:"streak"`
		} `json:"ranking"`
		IdentityChecked bool `json:"identityChecked"`
	} `json:"researcher"`
	LastUpdater struct {
		UserName  string      `json:"userName"`
		AvatarURL interface{} `json:"avatarUrl"`
		Email     string      `json:"email"`
	} `json:"lastUpdater"`
}

type Finding struct {
	ID			string
	URL			string
	Title		string
	Researcher	string
	Severity	string
	Timestamp	time.Time
	Endpoint	string
}

func (i *Endpoint) Get() ([]Finding, error) {
	var findings []Finding

	if err := i.Authenticate(); err != nil {
		return  findings, errors.Wrap(err, "could not authenticate to intigriti API")
	}

	req, err := http.NewRequest(http.MethodGet, intigritiAPIFetchURL, nil)
	if err != nil {
		return findings, errors.Wrap(err, "could not create http request to intigriti")
	}

	req.Header.Set("X-Client", i.clientTag)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer " + i.authToken))

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return findings, errors.Wrap(err, "message sending to intigriti failed")
	}

	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return findings, errors.Errorf("message fetch from intigriti returned status code: %d", resp.StatusCode)
	}

	var fetchResp fetchResponse
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return findings, errors.Wrap(err, "could not read response")
	}

	if err := json.Unmarshal(respBytes, &fetchResp); err != nil {
		return findings, errors.Wrap(err, "could not decode intigriti fetch response")
	}

	for _, entry := range fetchResp {
		findings = append(findings, Finding{
			ID:			entry.InternalReference.Reference,
			URL:		entry.InternalReference.URL,
			Title:      entry.Title,
			Researcher: entry.Researcher.UserName,
			Severity:   entry.Severity.Value,
			Timestamp: 	time.Unix(int64(entry.CreatedAt), 0),
			Endpoint:   entry.Domain.Value,
		})
	}

	logrus.WithField("findings_size", len(fetchResp)).Info("found findings on intigriti")
	return findings, nil
}