/*
Parthasaradhi terugu
*/
package main

import (
	"fmt"
	"net/http"
)
func handlerFunc(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"<h1>welcome to my website</h1>")

}

func main() {
	http.HandleFunc("/home/",handlerFunc)
	http.ListenAndServe(":3000",nil)
}
