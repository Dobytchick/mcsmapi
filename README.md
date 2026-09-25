# mcsmapi

[![Go CI](https://github.com/Dobytchick/mcsmapi/actions/workflows/ci.yml/badge.svg)](https://github.com/Dobytchick/mcsmapi/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/Dobytchick/mcsmapi.svg)](https://pkg.go.dev/github.com/Dobytchick/mcsmapi)
[![License: Apache 2.0](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)

A Go client for the [MCSManager](https://github.com/MCSManager/MCSManager) panel API. Manage users, daemons, instances, files, Docker images, and dashboard data from Go.

## Install

Requires Go 1.23.4 or newer.

```sh
go get github.com/Dobytchick/mcsmapi
```

## Quick start

Create an API key in MCSManager, then set `MCSM_API_KEY` in your environment. Use HTTPS when connecting to a remote panel: the MCSManager API sends the key as a query parameter.

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/Dobytchick/mcsmapi"
)

func main() {
    client := mcsmapi.NewClient(os.Getenv("MCSM_API_KEY"), "http://localhost:23333", nil)
    users, err := client.User.GetList(&mcsmapi.UserQueryParams{Page: 1, PageSize: 10})
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%+v\n", users.Data)
}
```

`NewClient` accepts an optional `*http.Client`; passing `nil` uses a client with a 10-second timeout. API and HTTP failures are returned as errors.

## API areas

| Client | Typical operations |
| --- | --- |
| `client.User` | Search, create, update, and delete users |
| `client.Daemon` | Add, remove, and connect daemons |
| `client.Instance` | List, create, control, and send commands to instances |
| `client.File` | List, read, update, copy, move, archive, and delete files |
| `client.Image` | Inspect images, containers, networks, and build progress |
| `client.Dashboard` | Read the panel overview |

### Example: send a command

```go
result, err := client.Instance.SendCommandResult("instance-uuid", "daemon-uuid", "say Hello from Go")
if err != nil {
    log.Fatal(err)
}
fmt.Println(result.Data)
```

### File transfers

`File.Download` and `File.Upload` request one-time transfer credentials from the panel. The actual download or multipart upload takes place against the daemon address returned by the API. See the [MCSManager file API](https://docs.mcsmanager.com/apis/api_fileManager.html) for that second step.

Browse the [usage guide](docs/USAGE.md), [runnable examples](examples/), and the [Go package reference](https://pkg.go.dev/github.com/Dobytchick/mcsmapi) for request and response types.

## Development

```sh
go test ./...
go vet ./...
```

Each example has its own directory and can be run with `go run ./examples/<name>` after setting `MCSM_API_KEY` and replacing placeholder IDs.

## License

[Apache 2.0](LICENSE).
