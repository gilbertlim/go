package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	// sql.DB 객체 생성
	// 실제 db connection은 이루어지지 않고, 정보만 가지고 있는 객체
	// db connection은 query 등 실제 db에 연결이 필요한 시점에 이루어지게 됨
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/member_service?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// single row: QueryRow()
	// scan($address$): 쿼리 결과를 주소에 할당
	var name string
	err = db.QueryRow("SELECT name FROM tbl_member WHERE member_id='gilbert'").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	// multirow: Query()
	// Next() 메서드를 사용하여 다음 Row로 이동
	// Parameterized Query: SQL Injection 방지 목적으로 파라미터를 문자열 결합이 아닌 "별도 파라미터"로 대입시킴 (Placeholder) 사용
	var id string
	rows, err := db.Query("SELECT member_id, name FROM tbl_member WHERE phone LIKE $1", "010%")
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}

	// DML: Exec()
	result, err := db.Exec("INSERT INTO tbl_member VALUES ($1, $2, $3, $4, $5)",
		"zz",
		"zzkim",
		"zzkim@gmail.com",
		"010-1234-5678",
		"경기도 용인시")
	if err != nil {
		log.Fatal(err)
	}

	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println(n, "row inserted.")
	}

	dResult, err := db.Exec("DELETE FROM tbl_member WHERE member_id = $1", "zz")
	if err != nil {
		log.Fatal(err)
	}
	dn, err := dResult.RowsAffected()
	if n == 1 {
		fmt.Println(dn, "row deleted.")
	}

	// Prepared Statement
	// SQL문을 미리 준비시키는 기능 (추후 stmt가 호출될 때 준비된 SQL문을 빠르게 실행

	// stmt create
	stmt, err := db.Prepare("UPDATE tbl_member SET name=$1 WHERE member_id=$2")
	checkError(err)
	defer stmt.Close()

	// stmt run
	_, err = stmt.Exec("zzlim", "newbie2")
	checkError(err)
	_, err = stmt.Exec("zzyoon", "newbie3")
	checkError(err)
	_, err = stmt.Exec("zzzack", "newbie4")
	checkError(err)

	// Transaction
	tx, err := db.Begin()
	if err != nil {
		log.Panic(err)
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO tbl_member VALUES ($1, $2)", 15, "Jack")
	if err != nil {
		//에러메시지를 출력하고 panic() 호출.
		log.Panic(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Panic(err)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
