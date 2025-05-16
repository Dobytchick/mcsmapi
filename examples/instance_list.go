package main

import (
	"fmt"
	"log"

	"github.com/Dobytchick/mcsmapi"
)

func main() {
	client := mcsmapi.NewClient("your-api-key", "http://localhost:23333", nil)
	query := &mcsmapi.ListInstancesQuery{
		DaemonID: "daemon-uuid",
		Page:     1,
		PageSize: 100,
	}
	resp, err := client.Instance.GetList(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Instances: %+v\n", resp)
}
