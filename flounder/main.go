package main

import (
	"fmt"
	// "github.com/go-gitea/go-sdk/gitea"
	"code.gitea.io/sdk/gitea"
)

func main() {
	client := gitea.NewClient("http://192.168.200.20:3000", "3de174d4ee2773402438322e5ccf8add2d1c4033")
	a, _ := client.GetUserInfo("panzhengming")
	fmt.Println(a.Email)
	fmt.Println(a.AvatarURL)
	fmt.Println(a.FullName)
	if ( a.FullName == "") {
		fmt.Println("empty!")
	} else {
		fmt.Println("not empty!")
	}
	fmt.Println(a.UserName)
	username := "panzhengming"
	password := "123456"
	var accessToken = ""
	if accessToken == "" {
		token, terr := client.CreateAccessToken(
			username,
			password,
			gitea.CreateAccessTokenOption{Name: "drone"},
		)
		if terr != nil {
			fmt.Println(terr)
		}
		accessToken = token.Sha1
		fmt.Println(accessToken)
	}

	// fmt.Println(client.GetUserInfo("panzhengming").Email)
}


