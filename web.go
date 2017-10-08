package main

import (
	"net/http"
	"os"

	"github.com/kpango/glg"
	"github.com/ruybrito106/GeovisService/src/router"
)

func main() {

	router := router.NewRouter()
	serverAddr := ":8001"

	glg.Info("Starting server on address: " + os.Getenv("SERVER_HOST") + ":8001")
	err := http.ListenAndServe(serverAddr, router)

	if err != nil {
		glg.Error("Unable to listen on given address: " + err.Error())
	}

}
