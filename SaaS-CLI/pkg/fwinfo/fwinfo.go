package fwinfo

import (
	"fmt"
	"log"
	"os"
)

// FirewallInfo: represents firewall name and firewall API key
type FirewallInfo struct {
	Firewall string
	APIKey   []byte
}

//WriteHostname: takes the hostname as a string and writes to a file named host.data
func WriteHostname(s string) {
	err := os.WriteFile("./pkg/fwinfo/host.data", []byte(s), 0644)
	if err != nil {
		log.Fatal("Unable to write to host.data file")
	}
}

//WriteAPIKey: takes the API Key as an array of bytes and writes to a file named key.data
func WriteAPIKey(b []byte) {
	err := os.WriteFile("./pkg/fwinfo/key.data", b, 0644)
	if err != nil {
		log.Fatal("Unable to write to key.data file")
	}
	fmt.Println("Successfully wrote API key to key.data file!")
}

//Hostname: returns the hostname value stored in the host.data file
func Hostname() string {
	data, err := os.ReadFile("./pkg/fwinfo/host.data")
	if err != nil {
		log.Fatal("Unable to write to host.data file")
	}
	return string(data)
}