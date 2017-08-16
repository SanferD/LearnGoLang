package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
// root:@127.0.0.1:3306:
var stmtIn *sql.Stmt
func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/Hema")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtIn, err = db.Prepare("SELECT id FROM Users WHERE username = ?")
	if err != nil {
		panic(err)
	}
	defer stmtIn.Close()
	
	var id int
	err = stmtIn.QueryRow("bob").Scan(&id)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("id: %d\n", id)
}
