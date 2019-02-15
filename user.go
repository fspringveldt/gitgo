package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const base = "https://api.github.com/users"

func GetUser(user string) []byte {
	url := []string{base, user}

	fmt.Println(strings.Join(url, "/"))
	response, err := http.Get(strings.Join(url, "/"))

	if err != nil {
		fmt.Printf("")
		panic(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("")
		os.Exit(1)
	}

	return body
}
