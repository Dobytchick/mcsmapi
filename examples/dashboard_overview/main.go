package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Dobytchick/mcsmapi"
)

func main() {
	client := mcsmapi.NewClient(os.Getenv("MCSM_API_KEY"), "http://localhost:23333", nil)
	resp, err := client.Dashboard.GetOverview()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Panel overview: %+v\n", resp.Data)
}
