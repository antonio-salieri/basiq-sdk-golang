package Services

import (
	"encoding/json"
	"fmt"
	"github.com/basiqio/basiq-sdk-golang/Entities"
)

type ConnectionService struct {
	session Session
	user    Entities.User
}

func NewConnectionService(session *Session, user *Entities.User) *ConnectionService {
	return &ConnectionService{
		session: *session,
		user:    *user,
	}
}

func (cs *ConnectionService) GetConnection(connectionId string) (Entities.Connection, error) {
	body, err := cs.session.api.Send("GET", "users/"+cs.user.Id+"/connections/"+connectionId, nil)
	if err != nil {
		panic(err)
	}

	var data Entities.Connection

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(string(body))
		panic(err)
	}

	return data, err
}
