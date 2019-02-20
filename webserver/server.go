package webserver

import (
	"fmt"
	"github.com/alonana/goserver/service"
	"log"
	"net/http"
)

func makeUrlHandler(f func() (string, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		output, err := f()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, _ = fmt.Fprint(w, output)
	}
}
func Start() {
	http.HandleFunc("/start", makeUrlHandler(service.Start))
	http.HandleFunc("/stop", makeUrlHandler(service.Stop))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
