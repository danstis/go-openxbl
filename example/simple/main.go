package main

import (
	"fmt"
	"os"

	"github.com/danstis/go-openxbl/openxbl"
)

var token string = os.Getenv("openxbltoken")

func main() {
	api := openxbl.NewClient(token)

	user, err := api.FriendsService.Search("john")
	if err != nil {
		panic(err)
	}

	var gt string
	for _, v := range user.Settings {
		if v.ID == "Gamertag" {
			gt = v.Value
		}
	}

	fmt.Printf("Gamertag for user: %s", gt)
}
