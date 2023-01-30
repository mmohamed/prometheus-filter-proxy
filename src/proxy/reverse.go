package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ReversePrometheus(reverseProxy *httputil.ReverseProxy, prometheusServerURL *url.URL) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		modifyRequest(r, prometheusServerURL)
		reverseProxy.ServeHTTP(w, r)
	}
}

func modifyRequest(r *http.Request, prometheusServerURL *url.URL) {
	r.URL.Scheme = prometheusServerURL.Scheme
	r.URL.Host = prometheusServerURL.Host
	r.Host = prometheusServerURL.Host
	
	query := r.URL.Query().Get("query")
	filter := r.URL.Query().Get("filter")
	
	if filter != "" && query != "" {
		newQuery := "("+query+") and "+filter
		values := r.URL.Query()
		values.Del("filter")
		values.Set("query", newQuery)
		r.URL.RawQuery = values.Encode()
	}
}
