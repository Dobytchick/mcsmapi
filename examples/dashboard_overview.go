package main

import (
	"fmt"
	"log"

	"github.com/Dobytchick/mcsmapi"
)

func main() {
	client := mcsmapi.NewClient("your-api-key", "http://localhost:23333", nil)
	resp, err := client.Dashboard.GetOverview()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Panel overview: %+v\n", resp.Data)
}
