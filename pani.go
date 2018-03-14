/*
Parthasaradhi terugu
*/
package main

import (
	"fmt"
	"net/http"
)
func handlerFunc(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","text/html")
if r.URL.Path=="/"{
	fmt.Fprint(w,"<h1>welcome to my website</h1>")
} else if r.URL.Path=="/contact"{
	fmt.Fprint(w,"<h1>to get in touch please contact me at <a> href=\"mailto:terugup@gmail.com\">terugup@gmail.com</a>." )
} else {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w,"<h1>We could not find the path</h1>")
}
}

func main() {
	http.HandleFunc("/",handlerFunc)
	http.ListenAndServe("localhost:3000",nil)
}
