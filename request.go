package main

import "net/http"

func getPath(r *http.Request) string {
	return r.URL.Path[1:]
}
