package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Dobytchick/mcsmapi"
)

func main() {
	client := mcsmapi.NewClient(os.Getenv("MCSM_API_KEY"), "http://localhost:23333", nil)
	query := &mcsmapi.TryConnectDaemonQuery{UUID: "your-daemon-uuid"}
	resp, err := client.Daemon.TryConnect(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Daemon connect response: %+v\n", resp)
}
