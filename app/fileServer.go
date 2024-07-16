package app

import (
	"fmt"
	"net/http"
)

func RunFileServer(host, port, dir string) {
	http.Handle("/", http.FileServer(http.Dir(dir)))
	fmt.Printf("Server started on %s:%s, serving %s\n", host, port, dir)
	_ = http.ListenAndServe(host+":"+port, nil)
}
