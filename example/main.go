package main

import (
	"fmt"
	"github.com/928799934/go-pprof"
	"net/http"
)

func main() {
	_ = pprof.InitByString([]string{":81"})
	defer pprof.Close()
	http.HandleFunc("/ccc", func(wr http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(wr, "ok")
	})
	http.ListenAndServe(":8080", nil)
}
