package main

import (
	"log"
	"net/http"
	"simple-picture-server/controller"
)

func main(){
	http.Handle("/upload",controller.UploadContoller{})
	log.Fatal(http.ListenAndServe(":9898",nil))
}
