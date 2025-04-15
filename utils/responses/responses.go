package responses

import (
	"encoding/json"
	"net/http"
	"strings"
)

type ErrorHttpResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

// func SendOkResponse(w http.ResponseWriter, v interface{}) {
// 	responseBytes, _ := json.Marshal(v)
// 	w.Write(responseBytes)
// 	w.WriteHeader(http.StatusOK)
// }

func ParseJsonRequestBodyOrSendErrorResponse(w http.ResponseWriter, r *http.Request, v interface{}) error {
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "Bad Request", err)
		return err
	}

	return nil
}

func SendResponse(w http.ResponseWriter, code int, body interface{}) {
	responseBytes, _ := json.Marshal(body)
	w.WriteHeader(code)
	w.Write(responseBytes)
}

func SendErrorResponse(w http.ResponseWriter, code int, message string, err error) {
	var response ErrorHttpResponse
	response.Code = code
	response.Message = message
	response.Errors = strings.Split(err.Error(), "\n")

	responseBytes, _ := json.Marshal(response)
	w.WriteHeader(code)
	w.Write(responseBytes)
	w.Header().Set("Content-Type", "application/json")
}
