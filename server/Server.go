package server

import (
	"fmt"
	"net/http"

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

		switch command {
		case "get":
			result := mapdb.Get(key)
			fmt.Fprint(w, result)

		case "set":
			result := mapdb.Set(key, value)
			fmt.Fprint(w, result)
		}
	}
}

// Start db server
// e.g. port: 8080
func StartServer(port string) error {

	http.HandleFunc("/", handler)
	return http.ListenAndServe(":"+port, nil)
}
