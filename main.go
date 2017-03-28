package main

import (
	"fmt"
	"log"
	"net"
	"regexp"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}

	rep := regexp.MustCompile(`.*:(.*)$`)
	str := rep.ReplaceAllString(l.Addr().String(), "$1")
	fmt.Println(str)
}

