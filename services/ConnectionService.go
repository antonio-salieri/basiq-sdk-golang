package services

import (
	"encoding/json"
	"fmt"
	"github.com/basiqio/basiq-sdk-golang/errors"
)

type ConnectionService struct {
	Session Session
	user    User
}

type Connection struct {
	Id          string       `json:"id"`
	Status      string       `json:"status"`
	LastUsed    string       `json:"lastUsed"`
	Institution Institution  `json:"institution"`
	Accounts    AccountsList `json:"accounts"`
	Service     *ConnectionService
}

func NewConnectionService(session *Session, user *User) *ConnectionService {
	return &ConnectionService{
		Session: *session,
		user:    *user,
	}
}

func (cs *ConnectionService) GetConnection(connectionId string) (Connection, *errors.APIError) {
	var data Connection

	data.Service = cs

	body, _, err := cs.Session.api.Send("GET", "users/"+cs.user.Id+"/connections/"+connectionId, nil)
	if err != nil {
		return data, err
	}


	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(string(body))
		return data, &errors.APIError{Message: err.Error()}
	}

	return data, nil
}

func (cs *ConnectionService) ForConnection(connectionId string) Connection {
	var data Connection

	data.Service = cs
	data.Id = connectionId

	return data
}

func (cs *ConnectionService) NewConnection(institutionId string, loginId string, password string, securityCode string) (Job, *errors.APIError) {
	var data Job
	data.Service = cs

	jsonBody := []byte(`{"institution": {"id": "` + institutionId + `"}, "loginId": "` + loginId + `", "password":"` + password + `"`)
	if securityCode != "" {
		jsonBody = append(jsonBody, []byte(`, "securityCode": "`+securityCode+`"}`)...)
	} else {
		jsonBody = append(jsonBody, []byte(`}`)...)
	}

	body, _, err := cs.Session.api.Send("POST", "users/"+cs.user.Id+"/connections", jsonBody)
	if err != nil {
		return data, err
	}


	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(string(body))
		return data, &errors.APIError{Message: err.Error()}
	}

	return data, nil
}

func (cs *ConnectionService) RefreshConnection(connectionId string) (Job, *errors.APIError) {
	var data Job
	data.Service = cs

	body, _, err := cs.Session.api.Send("POST", "users/"+cs.user.Id+"/connections/"+connectionId+"/refresh", nil)
	if err != nil {
		return data, err
	}


	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(string(body))
		return data, &errors.APIError{Message: err.Error()}
	}

	return data, nil
}

func (cs *ConnectionService) UpdateConnection(connectionId, password string) (Job, *errors.APIError) {
	var data Job
	data.Service = cs

	jsonBody := []byte(`{"password":"` + password + `"}`)

	body, _, err := cs.Session.api.Send("POST", "users/"+cs.user.Id+"/connections/"+connectionId, jsonBody)
	if err != nil {
		return data, err
	}


	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(string(body))
		return data, &errors.APIError{Message: err.Error()}
	}

	return data, nil
}

func (cs *ConnectionService) DeleteConnection(connectionId string) *errors.APIError {
	var data Job
	data.Service = cs

	_, _, err := cs.Session.api.Send("DELETE", "users/"+cs.user.Id+"/connections/"+connectionId, nil)
	if err != nil {
		return err
	}

	return nil
}

func (cs *ConnectionService) GetJob(jobId string) (Job, *errors.APIError) {
	var data Job
	data.Service = cs

	body, _, err := cs.Session.api.Send("GET", "jobs/"+jobId, nil)
	if err != nil {
		return data, err
	}


	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(string(body))
		return data, &errors.APIError{Message: err.Error()}
	}

	return data, nil
}

func (c *Connection) Refresh() (Job, *errors.APIError) {
	return c.Service.RefreshConnection(c.Id)
}

func (c *Connection) Update(password string) (Job, *errors.APIError) {
	return c.Service.UpdateConnection(c.Id, password)
}

func (c *Connection) Delete() *errors.APIError {
	return c.Service.DeleteConnection(c.Id)
}
