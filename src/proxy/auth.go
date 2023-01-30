package proxy

import (
	"crypto/subtle"
	"net/http"
)

const (
	realm = "Prometheus Filter Proxy"
)

func BasicAuth(handler http.HandlerFunc, auth string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		authorized := subtle.ConstantTimeCompare([]byte(user+":"+pass), []byte(auth)) == 1
		if auth != "" && (!ok || !authorized) {
			writeUnauthorisedResponse(w)
			return
		}
		handler(w, r)
	}
}

func writeUnauthorisedResponse(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
	w.WriteHeader(401)
	w.Write([]byte("Unauthorised\n"))
}
