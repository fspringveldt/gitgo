package main

import "fmt"

func moo () string {
	return "Moo"
}

func swap(x,y string) (string, string){
	return y, x
}

func main() {
	fmt.Println(swap("Hello","World"))
	fmt.Println(moo())
}

