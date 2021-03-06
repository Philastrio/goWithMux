package api

import (
	"EricLau/api/auto"
	"EricLau/api/router"
	"EricLau/config"
	"fmt"

	"log"
	"net/http"
)

func Run() {
	config.Load()
	auto.Load()
	fmt.Printf("\n\tListening [::]:%d\n", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.NEW()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
