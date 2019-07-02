package main

import (
	"os"
	"fmt"
	"log"
	"flag"
    "net/http"
    
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

func main() {
	userDB:=map[string]int{
		"jing":20,
		"golang":30,
		"java":40,
		"php":50,
	}

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "jing")
	})
    http.HandleFunc("/product",product)
    http.HandleFunc("/user/",func(w http.ResponseWriter, r *http.Request){
        name:=r.URL.Path[len("/user/"):]
        age:=userDB[name]
        fmt.Fprintf(w,"My Name is %s %d year old", name,age)
    })
    
    flag.Parse()
    log.Printf("listening on %s\n", *listenAddr)

    http.ListenAndServe(":8080", nil)
}

func product(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Product Request")
}

func user(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"User Request")
}