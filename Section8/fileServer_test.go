package Section8

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	//localhost:8080/static/index.html. strip prefix untuk menghapus static pada saat membaca directory
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources
var resources embed.FS

func TestFileServerGolangEmbed(t *testing.T) {
	directory, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(directory))
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	//localhost:8080/static/index.html. strip prefix untuk menghapus static pada saat membaca directory
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
