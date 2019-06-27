package utils

import "net/http"

//Message ...
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

//Respond ...
func Respond(w http.ResponseWriter, data map[string]interface{}) {

}
