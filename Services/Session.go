package Services

import (
	"encoding/json"
	"fmt"
	"github.com/basiqio/basiq-sdk-golang/Utilities"
	"time"
)

type Session struct {
	apiKey string
	api    *Utilities.API
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
		api:    Utilities.NewAPI("https://au-api.basiq.io/"),
		token: Token{
			value:     "",
			validity:  0,
			refreshed: time.Now(),
		},
	}

	session.getToken()

	return session
}

func (s *Session) getToken() (Token, *Utilities.APIError) {
	var token Token

	if time.Now().Sub(s.token.refreshed) < s.token.validity {
		return s.token, nil
	}

	body, statusCode, err := s.api.SetHeader("Authorization", "Basic "+s.apiKey).
		SetHeader("basiq-version", "1.0").
		SetHeader("content-type", "application/json").
		Send("POST", "oauth2/token", nil)

	if err != nil {
		return token, &Utilities.APIError{Message: err.Error()}
	}
	if statusCode > 299 {
		response, err := Utilities.ParseError(body)
		if err != nil {
			return token, &Utilities.APIError{Message: err.Error()}
		}

		return token, &Utilities.APIError{
			Response: response,
			Message:  response.GetMessages(),
			StatusCode: statusCode,
		}
	}

	var data AuthorizationResponse

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(string(body))
		return token, &Utilities.APIError{Message: err.Error()}
	}

	s.api.SetHeader("Authorization", "Bearer "+data.AccessToken)

	return Token{
		value:     data.AccessToken,
		validity:  time.Duration(data.ExpiresIn) * time.Second,
		refreshed: time.Now(),
	}, nil
}
