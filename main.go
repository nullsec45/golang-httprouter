package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
		fmt.Fprint(writer, "Hello Http Router")
	})

	server := http.Server {
		Handler : router,
		Addr : "localhost:3000",
	}

	server.ListenAndServe()
}

