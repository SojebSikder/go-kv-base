package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sojebsikder/go-kv-base/src/engine/mapdb"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	// u, err := url.Parse(r.URL.String())
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// params := u.Query()

	//
	switch r.Method {
	case "GET":

	case "POST":
		key := r.FormValue("key")
		value := r.FormValue("value")
		command := r.FormValue("command")
		command = strings.ToLower(command)

		_commands := map[string]string{
			"get":    "get",
			"set":    "set",
			"delete": "delete",
			"flush":  "flush",
		}

		if _commands[command] == command {
			switch command {
			case "get":
				result := mapdb.Get(key)
				if result == "" {
					fmt.Fprint(w, "'"+key+"' key not found")
					return
				}
				fmt.Fprint(w, result)

			case "set":
				result := mapdb.Set(key, value)
				fmt.Fprint(w, result)

			case "delete":
				result := mapdb.Delete(key)
				fmt.Fprint(w, result)

			case "flush":
				result := mapdb.Flush()
				fmt.Fprint(w, result)
			}
		} else {
			fmt.Fprint(w, "'"+command+"' command not supported")
		}

	}
}

// Start db server
// e.g. port: 8080
func StartServer(port string) error {

	http.HandleFunc("/", handler)
	return http.ListenAndServe(":"+port, nil)
}
