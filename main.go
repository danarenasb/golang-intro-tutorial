package main

import (
	"fmt"
	"log"
)

func returnString(s string) (string, error) {
	return "", fmt.Errorf("This is an error message")
}

func main() {
	msg, err := returnString("This is a message")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(msg)
}
