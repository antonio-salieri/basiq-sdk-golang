package services

import (
	"encoding/json"
	"fmt"
	"github.com/basiqio/basiq-sdk-golang/utilities"
	"github.com/basiqio/basiq-sdk-golang/errors"
	"time"
)

type Session struct {
	apiKey string
	api    *utilities.API
	token  Token
}

type Token struct {
	value     string
	validity  time.Duration
	refreshed time.Time
}

type AuthorizationResponse struct {
	AccessToken string        `json:"access_token"`
	Type        string        `json:"type"`
	ExpiresIn   time.Duration `json:"expires_in"`
}

func NewSession(apiKey string) *Session {
	session := &Session{
		apiKey: apiKey,
		api:    utilities.NewAPI("https://au-api.basiq.io/"),
		token: Token{
			value:     "",
			validity:  0,
			refreshed: time.Now(),
		},
	}

	session.getToken()

	return session
}

func (s *Session) getToken() (Token, *errors.APIError) {
	var token Token

	if time.Now().Sub(s.token.refreshed) < s.token.validity {
		return s.token, nil
	}

	body, statusCode, err := s.api.SetHeader("Authorization", "Basic "+s.apiKey).
		SetHeader("basiq-version", "1.0").
		SetHeader("content-type", "application/json").
		Send("POST", "oauth2/token", nil)

	if err != nil {
		return token, &errors.APIError{Message: err.Error()}
	}
	if statusCode > 299 {
		response, err := errors.ParseError(body)
		if err != nil {
			return token, &errors.APIError{Message: err.Error()}
		}

		return token, &errors.APIError{
			Response: response,
			Message:  response.GetMessages(),
			StatusCode: statusCode,
		}
	}

	var data AuthorizationResponse

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(string(body))
		return token, &errors.APIError{Message: err.Error()}
	}

	s.api.SetHeader("Authorization", "Bearer "+data.AccessToken)

	return Token{
		value:     data.AccessToken,
		validity:  time.Duration(data.ExpiresIn) * time.Second,
		refreshed: time.Now(),
	}, nil
}
