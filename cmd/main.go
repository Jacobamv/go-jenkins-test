package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/test", Test)

	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
