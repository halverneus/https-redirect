package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/halverneus/https-redirect/lib/config"
)

var (
	// Values to be overridden to simplify unit testing.
	redirectHandler = redirect
)

// Run server.
func Run() error {
	if config.Get.Debug {
		config.Log()
	}

	binding := fmt.Sprintf("%s:%d", config.Get.Host, config.Get.Port)
	httpServer := &http.Server{
		Addr:    binding,
		Handler: http.HandlerFunc(redirectHandler),
	}
	return httpServer.ListenAndServe()
}

// redirect a request from http://some.where/out/there to
// https://some.where/out/there .
func redirect(w http.ResponseWriter, r *http.Request) {
	// Build URL and apply redirect.
	target := url.URL{
		Scheme:   "https",
		Host:     r.Host,
		Path:     r.URL.Path,
		RawQuery: r.URL.RawQuery,
		Fragment: r.URL.Fragment,
	}
	if config.Get.Debug {
		log.Printf("Redirected %s to %s\n", r.URL.String(), target.String())
	}
	http.Redirect(w, r, target.String(), http.StatusTemporaryRedirect)
}
