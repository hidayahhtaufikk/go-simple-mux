package appHttpResponse

import (
	"encoding/json"
	"net/http"
)

type jsonResponder struct{}

func (j *jsonResponder) Write(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if data == nil {
		return
	}
	content, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = w.Write(content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func (j *jsonResponder) Data(w http.ResponseWriter, status int, message string, data interface{}) {
	content := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	j.Write(w, content)
}

func (j *jsonResponder) Error(w http.ResponseWriter, status int, error string) {
	content := ErrorResponse{
		ErrorID: status,
		Message: error,
	}
	j.Write(w, content)
}

func NewJSONResponder() IResponder {
	return &jsonResponder{}
}
