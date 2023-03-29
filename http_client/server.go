package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
)


func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("before request")
		next.ServeHTTP(w, r)
        log.Println("after request")
	})
}


func main() {
	
    r := mux.NewRouter()


    
    r.Use(commonMiddleware)

    r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, world!"))
    })

    log.Fatal(http.ListenAndServe(":3333", r))
}
