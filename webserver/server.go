package webserver

import (
	"fmt"
	"github.com/alonana/goserver/logging"
	"github.com/alonana/goserver/service"
	"github.com/go-errors/errors"
	"net/http"
)

func makeUrlHandler(f func() (string, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		output, err := f()
		if err != nil {

			addErrorDetails := fmt.Sprintf("error in URL %v", r.URL)
			logging.AppLogger.Error(errors.WrapPrefix(err, addErrorDetails, 0))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, _ = fmt.Fprint(w, output)
	}
}
func Start() {
	http.HandleFunc("/start", makeUrlHandler(service.Start))
	http.HandleFunc("/stop", makeUrlHandler(service.Stop))
	logging.AppLogger.Error(http.ListenAndServe(":8080", nil))
}
