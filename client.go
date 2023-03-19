package main

import (
	"fmt" //Print package
	"os"
	"log"
	"io/ioutil"
	"net/http" //http start, get
	"time"
)
func sendGet(dest string) {
	fmt.Println("Sending get to:", dest)
	resp, err := http.Get(dest)
	if err != nil {
		log.Fatalln(err)
		fmt.Printf("Client got no response")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf(string(body))
}

const url = "http://localhost"
const portNum = ":8080"
const about = "/about"

func chkServerUp(dest string) bool{
	resp, err := http.Get(dest)
	if err != nil && resp != nil {
		fmt.Println("Server not up")
		return false
	} else {
		return true
	}
}
func main() {
	//Wait for server to be up
	time.Sleep(500 * time.Millisecond)
	//Issue get to base URL to check if Server is up
	sendGet(url + portNum + about)

	time.Sleep(5 * time.Second)
	fmt.Println("Client exiting")
	os.Exit(0)
}
