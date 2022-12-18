// Author : sojebsikder<sojebsikder@gmail.com>
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sojebsikder/go-kv-base/server"
	"github.com/sojebsikder/go-kv-base/src/engine/mapdb"
)

func main() {
	// app info
	appName := "go-kv-base"
	version := "0.0.1"
	usage := "Welcome to go-base"
	description := "go-db is a simple database application"
	fmt.Printf("%s %s - %s\n", appName, version, usage)
	//

	if len(os.Args) < 2 {
		// run interactive mode
	} else {
		// args
		arg := os.Args[1]

		if arg == "version" {
			fmt.Println(appName + ": " + version)
		} else if arg == "help" {
			fmt.Println(description)
		} else if arg == "cli" {
			host := "http://localhost:8080"
			if len(os.Args) > 2 {
				host = os.Args[2]
			}
			if host == "" {
				fmt.Println("server url not provided")
				return
			}
			startCli(host)
		} else if arg == "start-server" {

			port := "8080"
			if len(os.Args) > 2 {
				port = os.Args[2]
			}
			startDBServer(port)
		} else {
			fmt.Println("Invalid command")
		}
	}

}

func startDBServer(port string) {
	fmt.Println("Server starting at " + port)
	error := server.StartServer(port)

	if error != nil {
		log.Fatal(error)
		return
	} else {
		fmt.Println("Server started at " + port)
	}
}

func startCli(_host string) {
	mapdb.Cli(_host)
}
