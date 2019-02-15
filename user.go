package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	base          = "https://api.github.com"
	pathSeperator = "/"
)

type User struct {
	Login     string `json:"login"`
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Company   string `json:"company"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
	Location  string `json:"location"`
}

// Gets the given users profile info from GitHub
func GetUser(user string) User {
	url := []string{base, "users", user}

	fmt.Println(strings.Join(url, pathSeperator))
	response, err := http.Get(strings.Join(url, pathSeperator))

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	var u User
	e := json.NewDecoder(response.Body).Decode(&u)

	if e != nil {
		panic(e)
	}
	return u
}
