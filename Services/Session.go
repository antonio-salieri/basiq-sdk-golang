package Services

import "github.com/basiqio/basiq-sdk-golang/Utilities/API"

type Session struct {
	api API
}

func NewSession() *Session {
	return &Session{
		api: API.NewAPI("https://au-api.basiq.io/"),
	}
}
