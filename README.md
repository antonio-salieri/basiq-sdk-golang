# Basiq.io Golang SDK

This is the documentation for the Golang SDK for Basiq.io API

## Introduction

To view the API docs, [click here](https://basiq.io/api/).

The SDKs to mirror the HTTP API's functionality and hierarchy.
The top level object needed for SDKs functionality is the Session
object which requires your API key to be instantiated.
You can create a new API key on the [dashboard](http://dashboard.basiq.io).

## API

The API of the SDK is manipulated using Services and Entities. Different
services return different entities, but the mapping is not one to one.

### Errors

If an action encounters an error, you will receive an APIError instance.
The struct contains all available data which you can use to act accordingly.

##### APIError struct
```go
type ResponseError struct {
        CorrelationId string `json:"correlationId"`
        Data          []ResponseErrorItem
}

type ResponseErrorItem struct {
        Code   string      `json:"code"`
        Title  string      `json:"title"`
        Detail string      `json:"detail"`
        Source ErrorSource `json:"source"`
}

type APIError struct {
        Data       map[string]interface{}
        Message    string
        Response   ResponseError
        StatusCode int
}
```

Check the [docs](https://basiq.io/api/) for more information about relevant
fields in the error object.

### SDK API List

#### Session

##### Creating a new Session object

```go
var session *Services.Session = Services.NewSession("YOUR_API_KEY")
```

#### User

The following are APIs available for the User service and entity

##### Creating a new UserService

```go
userService := Services.NewUserService(session)
```

##### Referencing a user
*Note: The following action will not send an HTTP request, and can be used
to perform additional actions for the instantiated user.*

```go
user := userService.ForUser(userId)
```

##### Creating a new User

```go
user, err := userService.CreateUser(&Services.UserData{
        Mobile: "+61410888555",
})
```

##### Updating a user instance

```go
err := user.Update(&Services.UserData{
        Mobile: "+61410888665",
})
```

##### Deleting a user

```go
err := user.Delete()
```

##### Get all of the user's accounts

```go
err := user.GetAccounts()
```

##### Get a user's single account

```go
err := user.GetAccount(accountId)
```

##### Get all of the user's transactions

```go
err := user.GetTransactions()
```

##### Get a user's single transaction

```go
err := user.GetTransaction(transactionId)
```



