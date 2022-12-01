package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSONThis(w http.ResponseWriter, statusCode int, data interface{}) {
	// tell writer to put a status code in the response
	w.WriteHeader(statusCode)
	// encode the data to be presented in json format
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSONThis(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSONThis(w, http.StatusBadRequest, nil)

}
