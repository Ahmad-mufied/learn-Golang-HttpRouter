package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
==	Router

-	Inti dari library HttpRouter adalah struct Router
-	Router ini merupakan implementasi dari http.Handler, sehingga kita bisa dengan mudah menambahkan ke dalam http.Server
-	Untuk membuat Router, kita bisa menggunakan function httprouter.New(), yang akan mengembalikan Router pointer


== HTTP Method

-	Router mirip dengan ServeMux, dimana kita bisa menambahkan route ke dalam Router
-	Kelebihan dibandingkan dengan ServeMux adalah, pada Router, kita bisa menentukan HTTP Method yang ingin kita gunakan, misal GET, POST, PUT, dan lain-lain
-	Cara menambahkan route ke dalam Router adalah gunakan function yang sama dengan HTTP Method nya, misal router.GET(), router.POST(), dan lain-lain


*/

func TestRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello World")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello World", string(body))
}
