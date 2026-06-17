// Package main is a tiny sample service used as a scan target. It depends on
// golang.org/x/text v0.3.6 and reaches language.Parse from an HTTP handler, so a
// reachability-aware dependency scan has a real, populated SBOM and a genuine
// candidate finding to report.
package main

import (
	"io"
	"net/http"

	"golang.org/x/text/language"
)

// parseHandler passes the untrusted Accept-Language header to language.Parse,
// giving the scanner a reachable path from an entrypoint to the dependency symbol.
func parseHandler(w http.ResponseWriter, r *http.Request) {
	tag := r.Header.Get("Accept-Language")
	parsed, _ := language.Parse(tag)
	_, _ = io.WriteString(w, parsed.String())
}

func main() {
	http.HandleFunc("/", parseHandler)
	_ = http.ListenAndServe("127.0.0.1:8080", nil)
}
