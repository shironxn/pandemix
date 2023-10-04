package util

import (
	"encoding/json"
	"log"
	"net/http"
	"pandemix/entity/web"
)

func LogError(w http.ResponseWriter, message string, status int, err interface{}) {
	if err != nil {
		log.Println(err)
	}

	res := web.ResponseError{
		Message: message,
		Error:   err,
	}

	response, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
