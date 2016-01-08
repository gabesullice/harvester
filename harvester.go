package main

import (
	"bufio"
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"syscall"
)

var (
	domain, username, password string
)

func init() {
	flag.StringVar(&domain, "domain", "", "The domain name of your harvest app. <yourapp>.harvestapp.com.")
	flag.Parse()

	if len(domain) == 0 {
		fmt.Println("You must pass a domain with the -domain flag\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	getCredentials()
}

func main() {
	// Create a new request
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s/account/who_am_i", domain), strings.NewReader(""))
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Set authentication and add Accept header
	req.SetBasicAuth(username, password)
	req.Header.Add("Accept", "application/json")

	// Perform the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println(string(body))
}

func getCredentials() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Username: ")
	username, _ = reader.ReadString('\n')
	username = strings.Trim(username, "\n")

	fmt.Printf("Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password = string(bytePassword)

	fmt.Println() // Add a newline after the password prompt
}
