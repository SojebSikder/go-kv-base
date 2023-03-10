package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sojebsikder/go-kv-base/src/engine/mapdb"
)

func typeof(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case map[string]interface{}:
		return "map"
	case []interface{}:
		return "[]map"
	//... etc
	default:
		return "unknown"
	}
}

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
		var bodyData map[string]any
		b, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal([]byte(b), &bodyData)
		if err != nil {
			panic(err)
		}
		r.Body.Close()

		var key string
		var command string
		if bodyData["key"] != nil {
			key = bodyData["key"].(string)
		}
		value := bodyData["value"]
		if bodyData["command"] != nil {
			command = bodyData["command"].(string)
		}
		// command := r.FormValue("command")
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
				if result == nil {
					fmt.Fprint(w, "'"+key+"' key not found")
					return
				}
				// fmt.Print(reflect.TypeOf(result))
				objType := typeof(result)

				switch objType {
				case "map":
					w.Header().Set("Content-Type", "application/json")
					// json.NewEncoder(w).Encode(result)
					jData, _ := json.Marshal(result)
					w.Write(jData)

				case "[]map":
					w.Header().Set("Content-Type", "application/json")
					jData, _ := json.Marshal(result)
					w.Write(jData)

				default:
					fmt.Fprint(w, result)
				}

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
