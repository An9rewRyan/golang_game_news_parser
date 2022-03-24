package main

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/lib/pq"
// )

// const Db_conn_str string = "user=postgres password=1234 dbname=gg sslmode=disable"

// func Connect_database() *sql.DB {
// 	db, err := sql.Open("postgres", Db_conn_str)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Подключение к базе данных было успешно")
// 	}
// 	return db
// }

// func main() {
// 	// https://pkg.go.dev/database/sql#DB.Exec
// 	db := Connect_database()
// 	rows, err := db.Query("Select link from recently_loaded_articles where src_link = " + "https://pkg.go.dev/database/sql#DB.Exec")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	cols, err := rows.Columns()
// 	if err != nil {
// 		fmt.Println("Failed to get columns", err)
// 		return
// 	}
// 	rawResult := make([][]byte, len(cols))
// 	result := make([]string, len(cols))
// 	dest := make([]interface{}, len(cols))
// 	for i, _ := range rawResult {
// 		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
// 	}
// 	for rows.Next() {
// 		err = rows.Scan(dest...)
// 		if err != nil {
// 			fmt.Println("Failed to scan row", err)
// 			return
// 		}

// 		for i, raw := range rawResult {
// 			if raw == nil {
// 				result[i] = "\\N"
// 			} else {
// 				result[i] = string(raw)
// 			}
// 		}

// 		fmt.Printf("%#v\n", result)
// 	}
// }
