package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	db()

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello"))
	})
	http.ListenAndServe(":8080", nil)
}

func db() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/member_service?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("DB 연동: %+v\n", db.Stats())

	conn, err := db.Query("SELECT * FROM tbl_member")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connection 생성: %+v\n", db.Stats())
	fmt.Println("")

	for conn.Next() {
		member_id := ""
		name := ""
		email := ""
		phone := ""
		address := ""
		if err := conn.Scan(&member_id, &name, &email, &phone, &address); err != nil {
			fmt.Println(err)
		}
		fmt.Println(member_id, name, email, phone, address)
	}
	fmt.Println("")
	
	conn.Close()
	fmt.Printf("Connection 연결 종료: %+v\n", db.Stats())

	db.Close()
	fmt.Printf("DB 연동 종료: %+v\n", db.Stats())
}