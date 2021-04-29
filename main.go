package main

import (
	"fmt"
"github.com/gorilla/mux"
	"net/http"
)

func main(){
	r := mux.NewRouter()
	r.Handle("/", health)
	fmt.Print("hello")
	err := http.ListenAndServe("", r)
	if err != nil {
		return
	}
}

func health(w http.ResponseWriter, r *http.Request){
	fmt.Print("ok")
}