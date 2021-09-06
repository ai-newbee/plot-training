package main

import (
	c "dl-base/pkg/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8086
	address := fmt.Sprintf("localhost:%d", port)
	log.Printf("address: %s", address)
	http.ListenAndServe(address, http.FileServer(http.Dir(c.StaticFolderName)))
}
