package Services

import (
	"encoding/json"
	"fmt"
)

type UserService struct {
	Session Session
}

func NewUserService(session *Session) *UserService {
	return &UserService{
		Session: *session,
	}
}

func (us *UserService) GetUser(userId string) (User, error) {
	var data User

	body, err := us.Session.api.Send("GET", "users/"+userId, nil)
	if err != nil {
		return data, err
	}

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(string(body))
		return data, err
	}

	data.Service = us

	return data, err
}

func (us *UserService) ForUser(userId string) *User {
	return &User{
		Id:      userId,
		Service: us,
	}
}

type ConnectionList struct {
	Count int          `json:"count"`
	Data  []Connection `json:"data"`
}

type User struct {
	Id          string         `json:"id"`
	Email       string         `json:"email"`
	Mobile      string         `json:"mobile"`
	Connections ConnectionList `json:"connections"`
	Service     *UserService
}

func (u *User) CreateConnection(institutionId string, loginId string, password string, securityCode string) (Job, error) {
	return NewConnectionService(&u.Service.Session, u).NewConnection(institutionId, loginId, password, securityCode)
}
