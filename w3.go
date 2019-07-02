package main

import (
	"os"
	"fmt"
	"log"
	"flag"
    "net/http"
    "github.com/gorilla/mux"
    
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

	router:=mux.NewRouter()
	router.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w,r,"index.html")
	})
	router.HandleFunc("/signup",func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w,r,"signup.html")
	})
	router.HandleFunc("/login",func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w,r,"login.html")
	})
	router.HandleFunc("/File",func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w,r,"name.txt")
	})	
    router.HandleFunc("/user/{name}",func(w http.ResponseWriter, r *http.Request){
        vars:=mux.Vars(r)
        name:=vars["name"]
        age:=userDB[name]
        fmt.Fprintf(w,"My Name is %s %d year old", name,age)
    }).Methods("GET")
    
    flag.Parse()
    log.Printf("listening on %s\n", *listenAddr)

    http.ListenAndServe(":8080", router)
}


func user(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"User Request")
}