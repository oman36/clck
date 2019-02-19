package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"runtime/debug"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error occur: %#v\n", r)
			debug.PrintStack()
			_, _ = fmt.Fprintf(w, "error")
		}
	}()
	switch r.Method {
	case "POST":
		postHandler(w, r)
	case "GET":
		getHandler(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	initDB()
	defer db.Close()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
