package main

import (
	"github.com/stefankoop/dssh"
)

func main() {
	// Create Config instance with remote username, server address and path to private key.
	ssh := &dssh.MakeConfig{
		User:   "root",
		Server: "10.10.10.111",
		Key:    "./id_rsa_test_plain",
		Port:   "22",
	}

	// Call Run method with command you want to run on remote server.
	response, err := ssh.Run("ps ax")
	// Handle errors
	if err != nil {
		panic("Can't run remote command: " + err.Error())
	} else {
		print(response)
	}
}
