package mapdb

var mapObject = map[string]string{}

// commands
func Get(key string) string {
	return mapObject[key]
}
func Set(key string, value string) int {
	mapObject[key] = value
	return 1
}
func Delete(key string) int {
	// check if key exists
	if _, ok := mapObject[key]; ok {
		delete(mapObject, key)
		return 1
	} else {
		return 0
	}
}
func Flush() int {
	kvlen := len(mapObject)
	mapObject = map[string]string{}
	return kvlen
}
