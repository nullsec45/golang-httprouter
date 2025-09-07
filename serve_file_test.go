package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"io"
	// "fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"embed"
	"io/fs"
)

//go:embed resources
var resources embed.FS



func TestServeFile(t *testing.T) {
	router := httprouter.New()
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	request := httptest.NewRequest("GET","http://localhost:3000/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	
	assert.Equal(t,"Hello Fajar Ganteng", string(body))
}