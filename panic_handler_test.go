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

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, i interface{}) {
		fmt.Fprint(writer,"Panic : ", i)
	}

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
		// fmt.Fprint(writer, "Hello World")
		panic("Uppss lalalala")
	})

	request := httptest.NewRequest("GET","http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	
	assert.Equal(t,"Panic : Uppss lalalala", string(body))
}