package main

import (
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func assert(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s\r\n", err.Error())
	}
}

func main() {
	cost := flag.Int("cost", bcrypt.DefaultCost, "Bcrypt cost")
	format := flag.String("format", "base64", "Output format (available: base64, base64-url, hex)")
	flag.Parse()

	input, err := ioutil.ReadAll(os.Stdin)
	assert(err)

	hash, err := bcrypt.GenerateFromPassword(input, *cost)
	assert(err)

	switch strings.ToLower(*format) {
	case "base64":
		fmt.Println(base64.StdEncoding.EncodeToString(hash))
	case "base64-url":
		fmt.Println(base64.URLEncoding.EncodeToString(hash))
	case "hex":
		fmt.Println(hex.EncodeToString(hash))
	default:
		log.Fatalf("Unknown output format: %s\r\n", *format)
	}
}
