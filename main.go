package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

/*
== Pengenalan HttpRouter

-	HttpRouter merupakan salah satu OpenSource Library yang populer untuk Http Handler di Go-Lang
-	HttpRouter terkenal dengan kecepatannya dan juga sangat minimalis
-	Hal ini dikarenakan HttpRouter hanya memiliki fitur untuk routing saja, tidak memiliki fitur apapun selain itu
-	https://github.com/julienschmidt/httprouter 

--	Menambah HttpRouter ke Project
-	go get github.com/julienschmidt/httprouter
-	go get github.com/stretchr/testify

*/

func main() {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello HttpRouter")
	})

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	server.ListenAndServe()
}
