package main

import "net/http"

func main() {
	http.HandleFunc("/product", Producthandler)
	http.ListenAndServe("9898", nil)
}

func Producthandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h3>Hello World</h3>"))
}
