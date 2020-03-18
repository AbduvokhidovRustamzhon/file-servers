package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var storagePathPtr = flag.String("storage", "serverdata", "folder for storing files")

func main() {
	flag.Parse()
	start()
}

func start() {
	createStorageFolder()
	mux := http.NewServeMux()
	fileSvc := files.NewFilesSvc(*storagePathPtr)
	server := app.NewServer(
		mux,
		fileSvc,
		*storagePathPtr,
	)

	server.InitRoutes()
	log.Fatal(http.ListenAndServe("0.0.0.0:9999", server))
}


func createStorageFolder() {
	err := os.Mkdir(*storagePathPtr, 0666)
	if err != nil {
		if !os.IsExist(err) {
			log.Fatalf("can't create directory: %s", err)
		}
	}
}
