package Services

import (
	"encoding/json"
	"fmt"
	"github.com/basiqio/basiq-sdk-golang/Utilities"
)

type UserService struct {
	Session Session
}

type UserData struct {
	Mobile string `json:"mobile,omitempty"`
	Email  string `json:"email,omitempty"`
}

func NewUserService(session *Session) *UserService {
	return &UserService{
		Session: *session,
	}
}

func (us *UserService) CreateUser(createData *UserData) (User, *Utilities.APIError) {
	var data User

	jsonBody, err := json.Marshal(createData)
	if err != nil {
		return data, &Utilities.APIError{Message: err.Error()}
	}

	body, statusCode, err := us.Session.api.Send("POST", "users", jsonBody)
	if err != nil {
		return data, &Utilities.APIError{Message: err.Error()}
	}
	if statusCode > 299 {
		response, err := Utilities.ParseError(body)
		if err != nil {
			return data, &Utilities.APIError{Message: err.Error()}
		}

		return data, &Utilities.APIError{
			Response:   response,
			Message:    response.GetMessages(),
			StatusCode: statusCode,
		}
	}

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(string(body))
		return data, &Utilities.APIError{Message: err.Error()}
	}

	data.Service = us

	return data, nil
}

func (us *UserService) GetUser(userId string) (User, *Utilities.APIError) {
	var data User

	body, statusCode, err := us.Session.api.Send("GET", "users/"+userId, nil)
	if err != nil {
		return data, &Utilities.APIError{Message: err.Error()}
	}
	if statusCode > 299 {
		response, err := Utilities.ParseError(body)
		if err != nil {
			return data, &Utilities.APIError{Message: err.Error()}
		}

		return data, &Utilities.APIError{
			Response:   response,
			Message:    response.GetMessages(),
			StatusCode: statusCode,
		}
	}

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(string(body))
		return data, &Utilities.APIError{Message: err.Error()}
	}

	data.Service = us

	return data, nil
}

func (us *UserService) UpdateUser(userId string, updateData *UserData) (User, *Utilities.APIError) {
	var data User

	jsonBody, err := json.Marshal(updateData)
	if err != nil {
		return data, &Utilities.APIError{Message: err.Error()}
	}

	body, statusCode, err := us.Session.api.Send("POST", "users/"+userId, jsonBody)
	if err != nil {
		return data, &Utilities.APIError{Message: err.Error()}
	}
	if statusCode > 299 {
		response, err := Utilities.ParseError(body)
		if err != nil {
			return data, &Utilities.APIError{Message: err.Error()}
		}

		return data, &Utilities.APIError{
			Response:   response,
			Message:    response.GetMessages(),
			StatusCode: statusCode,
		}
	}

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(string(body))
		return data, &Utilities.APIError{Message: err.Error()}
	}

	data.Service = us

	return data, nil
}

func (us *UserService) DeleteUser(userId string) *Utilities.APIError {
	body, statusCode, err := us.Session.api.Send("DELETE", "users/"+userId, nil)
	if err != nil {
		return &Utilities.APIError{Message: err.Error()}
	}
	if statusCode > 299 {
		response, err := Utilities.ParseError(body)
		if err != nil {
			return &Utilities.APIError{Message: err.Error()}
		}

		return &Utilities.APIError{
			Response:   response,
			Message:    response.GetMessages(),
			StatusCode: statusCode,
		}
	}

	return nil
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

func (u *User) CreateConnection(institutionId string, loginId string, password string, securityCode string) (Job, *Utilities.APIError) {
	return NewConnectionService(&u.Service.Session, u).NewConnection(institutionId, loginId, password, securityCode)
}

func (u *User) Update(update *UserData) *Utilities.APIError {
	user, err := u.Service.UpdateUser(u.Id, update)
	if err != nil {
		return err
	}

	*u = user

	return nil
}

func (u *User) Delete() *Utilities.APIError {
	return u.Service.DeleteUser(u.Id)
}
