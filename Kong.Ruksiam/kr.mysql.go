package main


import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "user:passworde@tcp(999.999.999.999)/databasename")
	log.Println("DB: ", db)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("select id, branchCode,branchName,branchStatus from branch")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var id int
		var branchCode string
		var branchName string
		var branchStatus string
		err = rows.Scan(&id, &branchCode, &branchName, &branchStatus)
		fmt.Printf("id : %d branchCode %s branchName %s branchStatus %s\n", id, branchCode, branchName, branchStatus)
	}


}
