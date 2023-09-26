package helpers

import "encoding/json"

//encode data
func JSONEncode(obj interface{}) string {
	json, _ := json.MarshalIndent(obj, "", "  ")
	return string(json)
}
