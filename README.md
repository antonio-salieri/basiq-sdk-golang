# Basiq.io Golang SDK

This is the documentation for the Golang SDK for Basiq.io API

## Introduction

To view the API docs, [click here](https://basiq.io/api/).

The SDKs to mirror the HTTP API's functionality and hierarchy.
The top level object needed for SDKs functionality is the Session
object which requires your API key to be instantiated.
You can create a new API key on the [dashboard](http://dashboard.basiq.io).

## Getting started

Install the SDK using:

```bash
go get -u https://github.com/basiqio/basiq-sdk-golang/
```

Import the package:
```go
import (
        "github.com/basiqio/basiq-sdk-golang/services"
)
```

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

#### Services

#### Session

##### Creating a new Session object

```go
var session *Services.Session = Services.NewSession("YOUR_API_KEY")
```

#### Entities

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

##### Updating a user instance [mut]

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
accounts, err := user.GetAccounts()
```

##### Get a user's single account

```go
account, err := user.GetAccount(accountId)
```

##### Get all of the user's transactions

```go
transactions, err := user.GetTransactions()
```

##### Get a user's single transaction

```go
transaction, err := user.GetTransaction(transactionId)
```

##### Create a new connection

```go
job, err := user.CreateConnection(&services.ConnectionData{
         Institution: &services.InstitutionData{
             Id: "AU00000",
         },
         LoginId:  "gavinBelson",
         Password: "hooli2018",
})
```

##### Refresh all connections

```go
err := user.RefreshAllConnections()
```

#### Connection

##### Refresh a connection

```go
job, err := connection.Refresh()
```

##### Update a connection

```go
job, err := connection.Update(password)
```

##### Delete a connection

```go
err := connection.Delete()
```

#### Job

##### Get the connection id (if available)

```go
connectionId := job.GetConnectionId()
```

##### Get the connection

```go
connection, err := job.GetConnection()
```

##### Get the connection after waiting for credentials step resolution
(interval is in milliseconds, timeout is in seconds)

```go
connection, err := job.WaitForCredentials(interval, timeout)
```

##### Get the connection after waiting for transactions step resolution
(interval is in milliseconds, timeout is in seconds)

```go
connection, err := job.WaitForTransactions(interval, timeout)
```

#### Transaction list

##### Getting the next set of transactions

```go
next, err := transactions.Next()
```

### Common usage examples

#### Fetching a list of institutions

```go
package main

import (
        "github.com/basiqio/basiq-sdk-golang/services"
        "log"
)

func main() {
        session, err := services.NewSession("YOUR_API_KEY")
        if err != nil {
            log.Printf("%+v", err)
        }

        institutions, err := session.GetInstitutions()
        if err != nil {
            log.Printf("%+v", err)
        }
}
```

#### Creating a new connection

```go
package main

import (
        "github.com/basiqio/basiq-sdk-golang/services"
        "log"
)

func main() {
        session, err := services.NewSession("YOUR_API_KEY")
        if err != nil {
            log.Printf("%+v", err)
        }

        user := session.ForUser(userId)

        job, err := user.CreateConnection(&services.ConnectionData{
                Institution: &services.InstitutionData{
                    Id: "AU00000",
                },
                LoginId:  "gavinBelson",
                Password: "hooli2018",
        })
        if err != nil {
                log.Printf("%+v", err)
        }

        // Poll our server to wait for the credentials step to be evaluated
        connection, err := job.WaitForCredentials(1000, 60)
        if err != nil {
                log.Printf("%+v", err)
        }
}
```

#### Fetching and iterating through transactions

```go
package main

import (
        "github.com/basiqio/basiq-sdk-golang/services"
        "log"
)

func main() {
        session, err := services.NewSession("YOUR_API_KEY")
        if err != nil {
            log.Printf("%+v", err)
        }

        user := session.ForUser(userId)

        fb := utilities.FilterBuilder{}
        fb.Eq("connection.id", "conn-id-213-id")
        transactions, err := user.GetTransactions(&fb)
        if err != nil {
                log.Printf("%+v", err)
        }

        for {
                next, err := transactions.Next()
                if err != nil {
                    log.Printf("%+v", err)
                    break
                }

                if next == false {
                    break
                }

                log.Println("Next transactions len:", len(transactions.Data))
        }
}
```