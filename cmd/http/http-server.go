package main

import (
	c "dl-base/pkg/config"
	"net/http"
)

func main() {
	foo()
	http.ListenAndServe(":8086", http.FileServer(http.Dir(c.StaticFolderName)))
}
