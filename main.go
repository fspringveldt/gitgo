package main

import (
	"fmt"
	flag "github.com/ogier/pflag"
	"os"
	"strings"
)

var (
	user string
)

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	users := strings.Split(user, ",")
	fmt.Printf("Searching for user(s): %s\n", users)

	fmt.Printf("%s", GetUser(user))
}
