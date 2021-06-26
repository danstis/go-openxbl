package main

import (
	"fmt"
	"os"

	"github.com/danstis/go-openxbl/openxbl"
)

var token string = os.Getenv("openxbltoken")

func main() {
	api := openxbl.NewClient(token)

	services, _, err := api.Services.List()
	if err != nil {
		panic(err)
	}

	for _, s := range *services {
		gs, _, _ := api.GameServers.Get(s.ID)
		fmt.Printf("GameServer for %q: %q\n", s.Details.Name, gs.GameHuman)
	}
}
