package main

import (
	"fmt"
)

func main() {
	client := Github{}
	client.WithAccessToken("fake_token").WithTeam("my fake team")

	fmt.Println(client)
}
