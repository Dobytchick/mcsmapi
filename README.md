# MCSManager API Go SDK

**MCSManager API Go SDK** is a client library for Go that provides convenient access to the [MCSManager](https://github.com/MCSManager/MCSManager) API (a panel for managing Minecraft and other servers). The SDK covers all major entities of the panel: users, daemons, instances, files, dashboard, and more.

## Features

- User management (create, update, delete, search)
- Daemon management (add, remove, connect)
- Instance management (create, start, stop, delete, send commands)
- File operations (list, read, write, copy, delete, archive)
- Retrieve panel and daemon information
- Flexible query parameter handling via Go structs and interfaces
- Fully typed request and response structures

## Installation

```sh
go get -u github.com/Dobytchick/mcsmapi
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/Dobytchick/mcsmapi"
)

func main() {
    client := mcsmapi.NewClient("your-api-key", "http://localhost:23333", nil)

    // Get user list
    params := &mcsmapi.UserQueryParams{Page: 1, PageSize: 10}
    users, err := client.User.GetList(params)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Users: %+v\n", users)
}
```

## Usage Examples

### Users

```go
// Get user list
params := &mcsmapi.UserQueryParams{Page: 1, PageSize: 20}
resp, err := client.User.GetList(params)

// Create a user
createReq := &mcsmapi.CreateUserRequest{
    Username: "newuser",
    Password: "securepassword",
    Permission: 1,
}
createResp, err := client.User.CreateUser(createReq)
```

### Daemons

```go
addReq := &mcsmapi.AddDaemonRequest{IP: "127.0.0.1", Port: 24444, APIKey: "daemon-key"}
resp, err := client.Daemon.Add(addReq)
```

### Instances

```go
listReq := &mcsmapi.ListInstancesQuery{DaemonID: "daemon-uuid", Page: 1, PageSize: 10}
instances, err
```

## More Examples

More usage examples can be found [here](https://github.com/Dobytchick/mcsmapi/tree/main/examples).