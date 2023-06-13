package main

import "net/http"

func main() {

	//routes
	http.HandleFunc("/", HomeHanlder)
	http.HandleFunc("/conctat", ContactHanlder)

	//start server
	http.ListenAndServe(":3000", nil)

}

func HomeHanlder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
func ContactHanlder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Contact page"))
}
