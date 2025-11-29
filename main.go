package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"socialai/backend"
	"socialai/handler"
	"socialai/util"
)

func main() {
	fmt.Println("started-service")

	config, err := util.LoadApplicationConfig("conf", "deploy.yaml")
	if err != nil {
		panic(err)
	}

	backend.InitElasticsearchBackend(config.ElasticsearchConfig)
	backend.InitGCSBackend(config.GCSConfig)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler.InitRouter(config.TokenConfig)))
}
