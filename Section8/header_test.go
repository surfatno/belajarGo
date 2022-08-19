package Section8

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprint(w,contentType)

}

func TestRequestHeader(t *testing.T){
	request:= httptest.NewRequest(http.MethodPost,"http://localhost:8080/", nil)
	request.Header.Add("Content-Type","application/json")
	recorder := httptest.NewRecorder()

	RequestHandler(recorder,request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}


func ResponseHeader(w http.ResponseWriter, r *http.Request){
	w.Header().Add("X-Powered-By","Surfatno's Team")
	fmt.Fprint(w,"OKE")
}
func TestResponseHeader(t *testing.T){
	request:= httptest.NewRequest(http.MethodPost,"http://localhost:8080/", nil)
	request.Header.Add("Content-Type","application/json")
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder,request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
	fmt.Println(response.Header.Get("x-pOwEred-by")) // < key insensitive
}