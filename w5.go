package main

import (
	"os" 
	"fmt"
	"log"
	"flag"
    "net/http"
    //"html/template"
    
)

var (
    listenAddr = flag.String("addr", getenvWithDefault("LISTENADDR", ":8080"), "HTTP address to listen on")
)

func getenvWithDefault(name, defaultValue string) string {
        val := os.Getenv(name)
        if val == "" {
                val = defaultValue
        }

        return val
}

type Product struct{
	Name string
	Price int
}

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/login",login)
    
    flag.Parse()
    log.Printf("listening on %s\n", *listenAddr)

    http.ListenAndServe(":8080", nil)
}


func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,"login.html")
}

func login(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,"login.html")
}

func user(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"User Request")
}