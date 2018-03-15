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

func (us *UserService) RefreshAllConnections(userId string) *Utilities.APIError {
	body, statusCode, err := us.Session.api.Send("POST", "users/"+userId+"/connections/refresh", nil)
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

func (us *UserService) ListAllConnections(userId string) (ConnectionList, *Utilities.APIError) {
	var data ConnectionList

	body, statusCode, err := us.Session.api.Send("GET", "users/"+userId+"/connections", nil)
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

	return data, nil
}

func (us *UserService) GetAccounts(userId string) (AccountsList, *Utilities.APIError) {
	var data AccountsList

	body, statusCode, err := us.Session.api.Send("GET", "users/"+userId+"/accounts", nil)
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

	return data, nil
}

func (us *UserService) GetAccount(userId string, accountId string) (Account, *Utilities.APIError) {
	var data Account

	body, statusCode, err := us.Session.api.Send("GET", "users/"+userId+"/accounts/"+accountId, nil)
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

	return data, nil
}

func (us *UserService) GetTransactions(userId string) (TransactionsList, *Utilities.APIError) {
	var data TransactionsList

	body, statusCode, err := us.Session.api.Send("GET", "users/"+userId+"/transactions", nil)
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

	return data, nil
}

func (us *UserService) GetTransaction(userId string, transactionId string) (Transaction, *Utilities.APIError) {
	var data Transaction

	body, statusCode, err := us.Session.api.Send("GET", "users/"+userId+"/transactions/"+transactionId, nil)
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

	return data, nil
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

func (u *User) RefreshAllConnections() *Utilities.APIError {
	return u.Service.RefreshAllConnections(u.Id)
}

func (u *User) ListAllConnections() (ConnectionList, *Utilities.APIError) {
	return u.Service.ListAllConnections(u.Id)
}

func (u *User) GetAccount(accountId string) (Account, *Utilities.APIError) {
	return u.Service.GetAccount(u.Id, accountId)
}

func (u *User) GetAccounts() (AccountsList, *Utilities.APIError) {
	return u.Service.GetAccounts(u.Id)
}

func (u *User) GetTransaction(transactionId string) (Transaction, *Utilities.APIError) {
	return u.Service.GetTransaction(u.Id, transactionId)
}

func (u *User) GetTransactions() (TransactionsList, *Utilities.APIError) {
	return u.Service.GetTransactions(u.Id)
}
