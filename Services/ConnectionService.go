package Services

import (
	"encoding/json"
	"fmt"
	"log"
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

func (cs *ConnectionService) GetConnection(connectionId string) (Connection, error) {
	var data Connection

	data.Service = cs

	body, err := cs.Session.api.Send("GET", "users/"+cs.user.Id+"/connections/"+connectionId, nil)
	if err != nil {
		return data, err
	}

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(string(body))
		return data, err
	}

	return data, err
}

func (cs *ConnectionService) NewConnection(institutionId string, loginId string, password string, securityCode string) (Job, error) {
	var data Job
	data.Service = cs

	jsonBody := []byte(`{"institution": {"id": "` + institutionId + `"}, "loginId": "` + loginId + `", "password":"` + password + `"`)
	if securityCode != "" {
		jsonBody = append(jsonBody, []byte(`, "securityCode": "`+securityCode+`"}`)...)
	} else {
		jsonBody = append(jsonBody, []byte(`}`)...)
	}

	body, err := cs.Session.api.Send("POST", "users/"+cs.user.Id+"/connections", jsonBody)
	if err != nil {
		return data, err
	}

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(string(body))
		return data, err
	}

	return data, err
}

func (cs *ConnectionService) RefreshConnection(connectionId string) (Job, error) {
	var data Job
	data.Service = cs

	body, err := cs.Session.api.Send("POST", "users/"+cs.user.Id+"/connections/"+connectionId+"/refresh", nil)
	if err != nil {
		return data, err
	}

	log.Println(string(body))

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(string(body))
		return data, err
	}

	return data, err
}

func (cs *ConnectionService) GetJob(jobId string) (Job, error) {
	var data Job
	data.Service = cs

	body, err := cs.Session.api.Send("GET", "jobs/"+jobId, nil)
	if err != nil {
		return data, err
	}

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(string(body))
		return data, err
	}

	return data, err
}

func (c *Connection) Refresh() (Job, error) {
	return c.Service.RefreshConnection(c.Id)
}
