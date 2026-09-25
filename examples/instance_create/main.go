package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Dobytchick/mcsmapi"
)

func main() {
	client := mcsmapi.NewClient(os.Getenv("MCSM_API_KEY"), "http://localhost:23333", nil)
	config := &mcsmapi.InstanceConfig{
		Nickname:     "MyServer",
		StartCommand: "./start.sh",
	}
	resp, err := client.Instance.Create("daemon-uuid", config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created instance: %+v\n", resp.Data)
}
