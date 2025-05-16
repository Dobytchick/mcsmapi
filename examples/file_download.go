package main

import (
	"fmt"
	"log"

	"github.com/Dobytchick/mcsmapi"
)

func main() {
	client := mcsmapi.NewClient("your-api-key", "http://localhost:23333", nil)
	req := &mcsmapi.DownloadFileRequest{
		BaseRequest: mcsmapi.BaseRequest{
			DaemonID: "daemon-uuid",
			UUID:     "instance-uuid",
		},
		FileName: "vstdlib.dll",
	}
	resp, err := client.File.Download(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Download info: %+v\n", resp.Data)
}
