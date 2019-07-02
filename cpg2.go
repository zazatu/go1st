package main 

import (
    _ "github.com/lib/pq"
    "database/sql"
    "fmt"
)
 
func main() {
    // Connect to the DB, panic if failed
    db, err := sql.Open("postgres", "postgres://user:password@localhost/dbName?sslmode=disable")
    if err != nil {
        fmt.Println(`Could not connect to db`)
        panic(err)
    }
    defer db.Close()
}import (
    _ "github.com/lib/pq"
    "database/sql"
    "fmt"
)
 
func main() {
    // Connect to the DB, panic if failed
    db, err := sql.Open("postgres", "postgres://postgres:eaxmple@localhost/postgres?sslmode=disable")
    if err != nil {
        fmt.Println(`Could not connect to db`)
        panic(err)
    }
    defer db.Close()
}