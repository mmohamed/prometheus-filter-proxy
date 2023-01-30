package proxy

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/urfave/cli"
)

func Serve(c *cli.Context) error {
	prometheusServerURL, _ := url.Parse(c.String("prometheus-server"))
	serveAt := fmt.Sprintf(":%d", c.Int("port"))
	auth:= c.String("auth")

	http.HandleFunc("/", createHandler(prometheusServerURL, auth))
	if err := http.ListenAndServe(serveAt, nil); err != nil {
		log.Fatalf("Prometheus label-filter proxy can not start %v", err)
		return err
	}
	return nil
}

func createHandler(prometheusServerURL *url.URL, auth string) http.HandlerFunc {
	reverseProxy := httputil.NewSingleHostReverseProxy(prometheusServerURL)
	return LogRequest(BasicAuth(ReversePrometheus(reverseProxy, prometheusServerURL), auth))
}
