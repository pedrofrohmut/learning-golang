package json

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(writer http.ResponseWriter, status int, data any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(data)
}

func ReadRequest(request *http.Request, data any) error {
	var decoder = json.NewDecoder(request.Body)
	return decoder.Decode(data)
}
