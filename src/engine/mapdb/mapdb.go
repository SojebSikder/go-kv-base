package mapdb

var mapObject = map[string]any{}

// get commands
func Get(key string) any {
	return mapObject[key]
}

// set commands
func Set(key string, value string) int {
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
