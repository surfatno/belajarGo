package Section8

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		writer.WriteHeader(http.StatusBadRequest) //400
		fmt.Fprint(writer, "name is empty")
	} else {
		writer.WriteHeader(http.StatusOK) //200
		fmt.Fprint(writer, "Hi ", name)
	}
}
func TestResponseCodeInvalid(t *testing.T) {
	request := httptest.NewRequest("Get", "http://localhost:8080", nil)

	recorder := httptest.NewRecorder()
	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}

func TestResponseCodeSuccess(t *testing.T) {
	request := httptest.NewRequest("Get", "http://localhost:8080/?name=EKO", nil)
	recorder := httptest.NewRecorder()
	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}
