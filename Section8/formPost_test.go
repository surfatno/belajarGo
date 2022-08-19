package Section8
import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)
func FormPost(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil{
		panic(err)
	}

	firstName := r.PostForm.Get("first_name")
	lastName := r.PostForm.Get("last_name")
	fmt.Fprintf(w,"Hello %s %s",firstName,lastName)
}


func TestFormPost(t *testing.T){
	requestBody:= strings.NewReader("first_name=Eko&last_name=Khannedy")
	request:= httptest.NewRequest(http.MethodPost,"http://localhost:8080/", requestBody)
	request.Header.Add("Content-Type","application/x-www-form-urlencoded") // template form post

	recorder := httptest.NewRecorder()
	FormPost(recorder,request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}