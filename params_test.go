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

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
		id := params.ByName("id")
		text := "Product "+id
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET","http://localhost:3000/products/1", nil)
	recorder := httptest.NewRecorder()

	
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	
	assert.Equal(t,"Product 1", string(body))
}