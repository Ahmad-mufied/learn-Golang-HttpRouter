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
==	Panic Handler

-	Apa yang terjadi jika terjadi panic pada logic Handler yang kita buat?
-	Secara otomatis akan terjadi error, dan web akan berhenti mengembalikan response
-	Kadang saat terjadi panic, kita ingin melakukan sesuatu, misal memberitahu jika terjadi kesalahan di web, atau bahkan mengirim informasi log kesalahan yang terjadi
-	Sebelumnya, seperti yang sudah kita bahas di materi Go-Lang Web, jika kita ingin menangani panic, kita harus membuat Middleware khusus secara manual
-	Namun di Router, sudah disediakan untuk menangani panic, caranya dengan menggunakan attribute PanicHandler : func(http.ResponseWriter, *http.Request, interface{})


*/

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, error interface{}) {
		fmt.Fprint(writer, "Panic : ", error)
	}

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		panic("Ups")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Panic : Ups", string(body))
}
