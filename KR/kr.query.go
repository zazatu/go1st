package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
	"flag"
	"os"
	"net/http"
	"html/template"

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


type ResultData struct{
	Id int
	Code string
	Name string
	Status string
}

var templates = template.Must(template.ParseFiles("listsql.html"))

func main() {
	flag.Parse()
    log.Printf("listening on %s\n", *listenAddr)
    http.HandleFunc("/",result)
    http.HandleFunc("/delete",delete)

    http.ListenAndServe(":8080", nil)
	

}

func result(res http.ResponseWriter, r *http.Request) {
	var db, err = sql.Open("mysql", "user:password@tcp(999.999.999.999)/databasename")
	log.Println("DB: ", db)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
		fmt.Println(err)
		
	}
	
	db.SetMaxOpenConns(5)
	defer db.Close()

	rows, err := db.Query("select id, branchCode,branchName,branchStatus from branch")
	if err != nil {
		panic(err)
		fmt.Println(err)
	}

	tRes:=ResultData{}
	var results []ResultData
	for rows.Next() {
		var id int
		var branchCode,branchName,branchStatus string
		err = rows.Scan(&id,&branchCode,&branchName,&branchStatus)
		tRes.Id=id
		tRes.Code=branchCode
		tRes.Name=branchName
		tRes.Status=branchStatus
		results = append(results,tRes)
		if err != nil {
	    	panic(err)
		    fmt.Println(err)
     	}
	//	fmt.Printf("id : %d branchCode %s branchName %s branchStatus %s\n", id, branchCode, branchName, branchStatus)
	}
	templates.Execute(res,results)
	fmt.Println(results)
}

func delete(res http.ResponseWriter, req *http.Request) {
	var db, err = sql.Open("mysql", "user:password@tcp(999..999.999.999)/databasename")
	log.Println("DB: ", db)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	stmt,err:=db.Prepare("select * from branch where id=?")
	stmt.Exec(req.URL.Query().Get("id"))
	if err != nil {
		panic(err)
		fmt.Println(err)
		http.Redirect(res, req, "/",301)
	}
	fmt.Fprintf(res,"stmt")

}

