package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"io"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNotFound(t *testing.T) {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Not Found bre")	
	}) 

	router.GET("/fajarganteng", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
		fmt.Fprint(writer, "Hello World")
	})

	request := httptest.NewRequest("GET","http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	
	assert.Equal(t,"Not Found bre", string(body))
}