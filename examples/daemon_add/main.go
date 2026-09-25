package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Dobytchick/mcsmapi"
)

func main() {
	client := mcsmapi.NewClient(os.Getenv("MCSM_API_KEY"), "http://localhost:23333", nil)
	req := &mcsmapi.AddDaemonRequest{
		IP:        "localhost",
		Port:      25555,
		Prefix:    "",
		Title:     "Test Daemon",
		AccessKey: "your-access-key",
	}
	resp, err := client.Daemon.Add(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Daemon add response: %+v\n", resp)
}
