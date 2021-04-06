package main

import (
	"log"
	"net/http"
	"simple-picture-server/controller"
)

func main(){
	http.Handle("/upload",controller.UploadContoller{})
	http.HandleFunc("/view",controller.ViewHandler)

	log.Fatal(http.ListenAndServe(":9898",nil))
}

