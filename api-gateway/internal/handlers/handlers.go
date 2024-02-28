package handlers

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

    loginServiceURL := "http://login-service:9090"

    url, err := url.Parse(loginServiceURL)
    if err != nil {
        log.Printf("Failed to parse target URL: %v", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    proxy := httputil.NewSingleHostReverseProxy(url)

    r.URL.Host = url.Host
    r.URL.Scheme = url.Scheme
    r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
    r.Host = url.Host

    // Serve the request using the reverse proxy
    proxy.ServeHTTP(w, r)
}
