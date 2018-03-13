package Services

import (
	"encoding/json"
	"fmt"
	"github.com/basiqio/basiq-sdk-golang/Entities"
)

type UserService struct {
	session Session
}

func NewUserService(session *Session) *UserService {
	return &UserService{
		session: *session,
	}
}

func (us *UserService) GetUser(userId string) (Entities.User, error) {
	body, err := us.session.api.Send("GET", "users/"+userId, nil)
	if err != nil {
		panic(err)
	}

	var data Entities.User

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(string(body))
		panic(err)
	}

	return data, err
}

func (us *UserService) ForUser(userId string) *Entities.User {
	return &Entities.User{
		Id: userId,
	}
}
