package main

import (
	"fmt"
	"log"

	"github.com/Dobytchick/mcsmapi"
)

func main() {
	client := mcsmapi.NewClient("your-api-key", "http://localhost:23333", nil)
	req := &mcsmapi.GetFileListRequest{
		BaseRequest: mcsmapi.BaseRequest{
			DaemonID: "your-daemon-uuid",
			UUID:     "your-instance-uuid",
		},
		Target:   "/",
		Page:     0,
		PageSize: 100,
		FileName: "",
	}
	resp, err := client.File.GetFileList(req)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range resp.Data.Items {
		fmt.Println("------------------------------")
		fmt.Printf("File: %+v\n", file.Name)
		fmt.Printf("Size: %+v\n", file.Size)
		fmt.Printf("Type: %+v\n", file.Type)
		fmt.Printf("Time: %+v\n", file.Time)
		fmt.Printf("Mode: %+v\n", file.Mode)
		fmt.Println("------------------------------")
	}
}
