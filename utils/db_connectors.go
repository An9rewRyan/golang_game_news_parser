package utils

import (
	"database/sql"
	"fmt"
	"parser/config"
	"parser/root_structs"

	_ "github.com/lib/pq"
)

func Write_article_to_db(article root_structs.Article) string {
	Db := Connect_database()
	defer Db.Close()
	// insert_string_to_artcls := `Select src_link from articles where  src_link = "asdadadad";`
	insert_string_to_artcls := `insert into articles(title, content, pub_date, image_url, src_link, site_name)
					  			values(` + `'` + article.Title + `','` + article.Content + `','` + article.Pub_date + `','` + article.Image_url + `','` + article.Source_link + `','` + article.Site_alias + `')`
	insert_string_to_recent := `insert into recently_loaded_articles(pub_date, src_link, site_name)
								values('` + article.Pub_date + `','` + article.Source_link + `','` + article.Site_alias + `')`
	delete_string_to_recent := `delete from recently_loaded_articles where site_name = '` + article.Site_alias + `'
														   and pub_date = (select min(pub_date) from recently_loaded_articles)`
	rows, err := Db.Query(insert_string_to_artcls)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rows)
	rows, err = Db.Query(insert_string_to_recent)
	if err != nil {
		fmt.Println(err)
	}
	rows, err = Db.Query(delete_string_to_recent)
	if err != nil {
		fmt.Println(err)
	}
	return "suceeded!"
}

func Check_if_article_in_db(link string) bool {
	Db := Connect_database()
	var found string
	defer Db.Close()
	// fmt.Println("Вызывается фугкция!!!")
	check_string := `select src_link from recently_loaded_articles where src_link = '` + link + `'`
	if err := Db.QueryRow(check_string).Scan(&found); err != nil {
		fmt.Println(err, "Тут")
		return false
	}
	fmt.Println("Там")
	return true
}

func Connect_database() *sql.DB {
	db, err := sql.Open("postgres", config.Db_conn_str)
	if err != nil {
		fmt.Println(err)
	} else {
		// fmt.Println("Подключение к базе данных было успешно")
	}
	// defer db.Close()
	return db
}
