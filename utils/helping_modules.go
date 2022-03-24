package utils

import (
	"database/sql"
	"fmt"
	"parser/config"
	"parser/root_structs"

	_ "github.com/lib/pq"
)

func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func Write_article_to_db(article root_structs.Article) string {
	Db := Connect_database()
	defer Db.Close()
	insert_string := `insert into articles(title, content, pub_date, image_url)
	values(` + `'` + article.Title + `','` + article.Content + `','` + article.Pub_date + `','` + article.Image_url + `'` + `);`
	// fmt.Println(insert_string)
	res, err := Db.Exec(insert_string)
	if err != nil {
		fmt.Println(err)
		fmt.Println(insert_string)
	}
	fmt.Println(res)
	return "suceeded!"
}

func Connect_database() *sql.DB {
	db, err := sql.Open("postgres", config.Db_conn_str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Подключение к базе данных было успешно")
	}
	// defer db.Close()
	return db
}
