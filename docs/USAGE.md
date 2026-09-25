# Usage guide

This guide covers the request shapes that are easiest to get wrong. For the full API surface, see [Go Reference](https://pkg.go.dev/github.com/Dobytchick/mcsmapi) and the [MCSManager API docs](https://docs.mcsmanager.com/apis/get_apikey.html).

## Connect

Create an API key in the panel. It has the same permissions as the account that created it. The panel expects it in the `apikey` query parameter, so use HTTPS outside your own machine and keep request URLs out of logs. The SDK does not print the URL or key.

```go
client := mcsmapi.NewClient(os.Getenv("MCSM_API_KEY"), "https://panel.example.com", nil)
```

Pass a custom `*http.Client` as the third argument to configure transport, proxies, or timeout. The default timeout is 10 seconds. API and HTTP status failures are returned as errors.

## Instances

```go
instances, err := client.Instance.GetList(&mcsmapi.ListInstancesQuery{
    DaemonID: "daemon-uuid",
    Page: 1,
    PageSize: 10,
})
if err != nil { /* handle error */ }
_ = instances

command, err := client.Instance.SendCommandResult(
    "instance-uuid", "daemon-uuid", "say Hello",
)
if err != nil { /* handle error */ }
_ = command.Data // true when the API accepted the command
```

`SendCommand` remains available for source compatibility with earlier versions. Its response type predates the API's boolean command result; use `SendCommandResult` in new code.

## Files

Every file operation needs a daemon ID and instance UUID. File listing starts at page 0.

```go
files, err := client.File.GetFileList(&mcsmapi.GetFileListRequest{
    BaseRequest: mcsmapi.BaseRequest{DaemonID: "daemon-uuid", UUID: "instance-uuid"},
    Target: "/",
    Page: 0,
    PageSize: 20,
})
if err != nil { /* handle error */ }
_ = files.Data.Items
```

Update text files with `File.Update`. The SDK sends the target IDs in the URL and `target`/`text` in JSON.

```go
_, err := client.File.Update(&mcsmapi.UpdateFile{
    Target: &mcsmapi.UpdateFileRequest{DaemonID: "daemon-uuid", UUID: "instance-uuid"},
    FileData: &mcsmapi.UpdateFileRequestBody{Target: "/eula.txt", Text: "eula=true\n"},
})
```

To extract an archive, set `Targets` to one destination directory:

```go
_, err := client.File.UnzipTo(&mcsmapi.UnzipFile{
    Target: &mcsmapi.ZipFileRequest{BaseRequest: mcsmapi.BaseRequest{
        DaemonID: "daemon-uuid", UUID: "instance-uuid",
    }},
    FileData: &mcsmapi.UnzipFileRequestBody{
        Type: mcsmapi.CompressModeUnzip,
        Code: "utf-8",
        Source: "/backup.zip",
        Targets: "/restore",
    },
})
```

`File.Download` and `File.Upload` return one-time credentials and a daemon address. They do not transfer bytes. The next request goes directly to the daemon using the [documented download URL or multipart upload endpoint](https://docs.mcsmanager.com/apis/api_fileManager.html#upload-file). For upload configuration, set `UploadFileRequest.UploadDir`; the old `FileName` field remains as a compatibility alias.

## Examples and checks

Each directory under [`examples/`](../examples/) is a separate program. Set `MCSM_API_KEY`, replace placeholder IDs and endpoints, then run one with `go run ./examples/instance_list`. Run `go test ./...` and `go vet ./...` when changing the SDK.
