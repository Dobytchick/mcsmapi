package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Dobytchick/mcsmapi"
)

func main() {
	client := mcsmapi.NewClient(os.Getenv("MCSM_API_KEY"), "http://localhost:23333", nil)
	page := 0
	pageSize := 50 // or 100, depending on your server's max page size
	var allUsers []mcsmapi.UserData

	for {
		params := &mcsmapi.UserQueryParams{Page: page, PageSize: pageSize}
		resp, err := client.User.GetList(params)
		if err != nil {
			log.Fatal(err)
		}
		allUsers = append(allUsers, resp.Data.Data...)
		if page >= resp.Data.MaxPage {
			break
		}
		page++
	}

	for _, user := range allUsers {
		fmt.Printf("UUID: %s, Username: %s, Permission: %d\n", user.UUID, user.UserName, user.Permission)
	}
}
