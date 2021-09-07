package main

import (
	"fmt"
	"log"
	"net/http"
	c "plot-training/pkg/config"
)

func main() {
	port := 8086
	address := fmt.Sprintf("0.0.0.0:%d", port)
	log.Printf("address: %s", address)
	http.ListenAndServe(address, http.FileServer(http.Dir(c.StaticFolderName)))
}
