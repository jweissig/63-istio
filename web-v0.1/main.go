package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	// disable cache
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	fmt.Fprintf(w, "Version: %v", version)

}

// used to dump headers for debugging
func debugHandler(w http.ResponseWriter, r *http.Request) {

	startTime := time.Now()

	// disable cache
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	// set hostname (used for demo)
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprint(w, "Error:", err)
	}

	// dump headers
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "%v", string(requestDump))
	fmt.Fprintf(w, "Served-By: %v\n", hostname)
	fmt.Fprintf(w, "Serving-Time: %s\n", time.Now().Sub(startTime))
	fmt.Fprintf(w, "Build-Date: %v\n", builddate)
	fmt.Fprintf(w, "Build-Revision: %v", buildrevision)
	fmt.Fprintf(w, "Version: %v", version)
	return

}

var (
	builddate     = "2019-03-11"
	buildrevision = ""
	version       = "v0.1"
)

// mux
var router = mux.NewRouter()

func main() {

	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/debug", debugHandler)
	http.Handle("/", router)

	fmt.Println("Listening on port 5005...")
	//http.ListenAndServe("localhost:5005", router)
	http.ListenAndServe(":5005", handlers.CompressHandler(router))

}
