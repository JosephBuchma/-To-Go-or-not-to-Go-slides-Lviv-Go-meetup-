package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
)

// FBB OMIT
func fooBarBaz(n int) (string, bool) {
	var fbb string
	var ok = true

	switch {
	case n%15 == 0:
		fbb = "baz"
	case n%5 == 0:
		fbb = "bar"
	case n%3 == 0:
		fbb = "foo"
	default:
		ok = false
	}
	return fbb, ok
}

// FBB OMIT

// checkFBB OMIT

func fooBarBazDecorated(n int) string {
	if x, ok := fooBarBaz(n); ok {
		return fmt.Sprintf("%s (%d)", x, n)
	}
	return fmt.Sprintf("%d", n)
}

func checkFooBarBaz(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(fooBarBazDecorated(i))
	}
}

// checkFBB OMIT

// checkFBBClosure OMIT
func checkFBBFunc(n int) func() {
	return func() {
		checkFooBarBaz(n)
	}
}

// checkFBBClosure OMIT

// FBBSrv OMIT
func StartFooBarBazSrv() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request %s", r.URL)
		w.Header().Set("Connection", "close")
		i, err := strconv.Atoi(html.EscapeString(r.URL.Path[1:]))
		if err != nil {
			fmt.Fprintf(w, "Invalid number")
		} else {
			fmt.Fprintf(w, fooBarBazDecorated(i))
		}
	})
	log.Print("Serving at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// FBBSrv OMIT

// MAIN OMIT
func main() {
	// CHECKFBB OMIT
	//checkFooBarBaz(100)
	// CHECKFBB OMIT

	// STARTFBBSRV OMIT
	//StartFooBarBazSrv()
	// STARTFBBSRV OMIT

	// MFBBClosure OMIT
	//fbb30 := checkFBBFunc(30)
	//fbb30()
	// MFBBClosure OMIT
}

// MAIN OMIT
