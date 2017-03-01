package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang/glog"

	"goji.io"
	"goji.io/pat"
)

type Response struct {
	Header  *http.Request
	Message string
}

func resp(w http.ResponseWriter, r *http.Request) {

	name := pat.Param(r, "something")

	var response Response
	response.Header = r
	response.Message = "something " + r.Method + " : " + name

	glog.Info()
	log.Println(response.Header)
	log.Println(response.Message)
	fmt.Fprintf(w, "%s", response)

}

func main() {

	mux := goji.NewMux()
	mux.HandleFunc(pat.Put("/put/:something"), resp)
	mux.HandleFunc(pat.Get("/get/:something"), resp)
	mux.HandleFunc(pat.Post("/post/:something"), resp)
	mux.HandleFunc(pat.Delete("/get/:something"), resp)
	mux.HandleFunc(pat.Patch("/patch/:something"), resp)

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	const port = 8080
	log.Println("listening on : "+hostname+":", port)
	http.ListenAndServe(":8080", mux)
}
