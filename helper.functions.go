package main

import (
	"log"
	"strings"
)

func errorHandle(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}

func handleBadSpace(s string) string {
	var bad_space byte = 160
	return strings.ReplaceAll(s, string(bad_space), "")
}
