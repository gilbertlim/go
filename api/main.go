package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db()

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello"))
	})
	http.ListenAndServe(":8080", nil)
}

func db() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/member_service")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("DB 연동: %+v\n", db.Stats())

	conn, err := db.Query("SELECT * FROM tbl_member")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connection 생성: %+v\n", db.Stats())

	for conn.Next() {
		member_num := ""
		member_id := ""
		member_name := ""
		if err := conn.Scan(&member_num, &member_id, &member_name); err != nil {
			fmt.Println(err)
		}
		fmt.Println(member_num, member_id, member_name)
	}

	conn.Close()
	fmt.Printf("Connection 연결 종료: %+v\n", db.Stats())

	db.Close()
	fmt.Printf("DB 연동 종료: %+v\n", db.Stats())
}