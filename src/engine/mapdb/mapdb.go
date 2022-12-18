package mapdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var mapObject = map[string]any{}

// get commands
func Get(key string) any {
	return mapObject[key]
}

// set commands
func Set(key string, value any) int {
	mapObject[key] = value
	return 1
}

// delete commands
func Delete(key string) int {
	// check if key exists
	if _, ok := mapObject[key]; ok {
		delete(mapObject, key)
		return 1
	} else {
		return 0
	}
}

// flush commands
func Flush() int {
	kvlen := len(mapObject)
	mapObject = map[string]any{}
	return kvlen
}

func makeRequest(host string, postBody []byte) string {
	// connect to server
	// postBody, _ := json.Marshal(map[string]string{
	// 	"key":     "name",
	// 	"value":   "sojeb",
	// 	"command": "set",
	// })
	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(host, "application/json", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Print(sb)
	return sb
}

// client CLI
func Cli(host string) {

	var cmd string

	// load data from disk
	fmt.Println("Welcome to the simplest key-value memory database")
	for {
		fmt.Print(host + "> ")
		fmt.Scan(&cmd)

		if cmd == "set" {

			var key string
			var value string

			fmt.Print("Enter key: ")
			fmt.Scan(&key)

			// check if key already exists

			fmt.Print("Enter value: ")
			fmt.Scan(&value)

			postBody, _ := json.Marshal(map[string]string{
				"key":     "name",
				"value":   "sojeb",
				"command": "set",
			})
			result := makeRequest(host, postBody)
			if result == "1" {
				fmt.Println("Key added")
			} else {
				fmt.Println("res", result)
			}

		} else if cmd == "get" {

			var key string
			fmt.Print("Enter key: ")
			fmt.Scan(&key)
			postBody, _ := json.Marshal(map[string]string{
				"key":     "name",
				"command": "get",
			})
			result := makeRequest(host, postBody)
			fmt.Println("res", result)

		} else if cmd == "delete" {

			var key string
			fmt.Print("Enter key: ")
			fmt.Scan(&key)
			// check if key exists
			postBody, _ := json.Marshal(map[string]string{
				"key":     "name",
				"command": "delete",
			})
			result := makeRequest(host, postBody)
			fmt.Println(result)

		} else if cmd == "exit" {
			fmt.Println("Bye")
			break
		} else {
			fmt.Println("Invalid command")
		}
	}
}
