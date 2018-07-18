package main

import (
	"fmt"
	"github.com/drone/drone-go/drone"
	"golang.org/x/oauth2"
)

const (
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZXh0IjoicGFuemhlbmdtaW5nIiwidHlwZSI6InVzZXIifQ.DEpNXDRqcmpvAHUL-y7xaiYIYqfwjZn8RwUYpYWxmec"
	host = "http://192.168.200.20:8000"
)

func main () {
	// create an http client with oauth authentication
	config := new(oauth2.Config)
	auther := config.Client(
		oauth2.NoContext,
		&oauth2.Token{
			AccessToken: token,
		},
	)

	// create the drone client with authenticator
	client := drone.NewClient(host, auther)

	// gets the current user
	user, err := client.Self()
	fmt.Println(user, err)

	// gets the named repository information
	repo, err := client.Repo("panzhengming", "demo")
	fmt.Println(repo, err)


}