package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

/*
== Not Found Handler

-	Selain panic handler, Router juga memiliki not found handler
-	Not found handler adalah handler yang dieksekusi ketika client mencoba melakukan request URL yang memang tidak terdapat di Router
-	Secara default, jika tidak ada route tidak ditemukan, Router akan melanjutkan request ke http.NotFound, namun kita bisa mengubah nya
-	Caranya dengan mengubah router.NotFound = http.Handler


*/

func TestNotFound(t *testing.T) {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Gak ketemu")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Gak ketemu", string(body))
}
