package main

import (
	"os"
	"log"
	"flag"
	"net/http"
	"fmt"
	"time"
)


var (
    listenAddr = flag.String("addr", getenvWithDefault("LISTENADDR", ":8080"), "HTTP address to listen on")
)


type Cookie struct{
	Name string
	Value string
	Expires time.Time
}

func getenvWithDefault(name, defaultValue string) string {
        val := os.Getenv(name)
        if val == "" {
                val = defaultValue
        }
        return val
}

func main() {
	http.HandleFunc("/", index)
		
	flag.Parse()
    log.Printf("listening on %s\n", *listenAddr)
    http.ListenAndServe(":8080", nil)
}


func index(w http.ResponseWriter, r *http.Request){
	expiration:=time.Now().Add(time.Hour*24*365)
	cookie:=http.Cookie{Name:"user",Value:"jing", Expires:expiration}
	http.SetCookie(w,&cookie)
	fmt.Fprintf(w,"Create Cokkie")
}