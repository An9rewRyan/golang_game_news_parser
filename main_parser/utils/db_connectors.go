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
	insert_string_to_artcls := `insert into articles(title, content, pub_date, image_url, src_link, site_name)
					  			values(` + `'` + article.Title + `','` + article.Content + `','` + article.Pub_date + `','` + article.Image_url + `','` + article.Source_link + `','` + article.Site_alias + `')`
	insert_string_to_recent := `insert into recently_loaded_articles(pub_date, src_link, site_name)
								values('` + article.Pub_date + `','` + article.Source_link + `','` + article.Site_alias + `')`
	// delete_string_to_recent := `delete from recently_loaded_articles where site_name = '` + article.Site_alias + `'
	// 													   and pub_date = (select min(pub_date) from recently_loaded_articles)`
	rows, err := Db.Query(insert_string_to_artcls)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rows)
	rows, err = Db.Query(insert_string_to_recent)
	if err != nil {
		fmt.Println(err)
	}
	// rows, err = Db.Query(delete_string_to_recent)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	return "suceeded!"
}

func Check_if_article_in_db(link string) bool {
	Db := Connect_database()
	var found string
	defer Db.Close()
	check_string := `select src_link from recently_loaded_articles where src_link = '` + link + `'`
	if err := Db.QueryRow(check_string).Scan(&found); err != nil {
		return false
	}
	return true
}

func Connect_database() *sql.DB {
	db, err := sql.Open("postgres", config.Db_conn_str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Database connected succefully")
	}
	return db
}

func Create_articles_table() {
	Db := Connect_database()
	defer Db.Close()
	_, err := Db.Exec(`create table articles (
		article_id serial primary key,
		title varchar(300) not null,
		content text not null, 
		pub_date timestamp not null,
		image_url varchar(300) not null,
		src_link varchar(300) not null,
		site_name varchar(10) not null
	);`)
	if err != nil {
		fmt.Println(err)
	}
}

func Create_recently_loaded_articles_table() {
	Db := Connect_database()
	defer Db.Close()
	_, err := Db.Exec(`create table recently_loaded_articles (
		pub_date timestamp not null,
		src_link varchar(300) not null,
		site_name varchar(10) not null
	);`)
	if err != nil {
		fmt.Println(err)
	}
}
