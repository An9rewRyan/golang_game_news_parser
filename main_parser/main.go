package main

import (
	"fmt"
	"parser/config"
	"parser/root_structs"
	"parser/scrappers"
	"parser/sites_structs"
	"parser/utils"
	"time"
)

func create_articles_table() {
	Db := utils.Connect_database()
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

func create_recently_loaded_articles_table() {
	Db := utils.Connect_database()
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

func main() {
	//uncomment when launching for first time and comment again afterwards
	// create_articles_table()
	// create_recently_loaded_articles_table()

	for {
		site_paths := []root_structs.Article_paths{
			// sites_structs.Get_igrm_paths(),
			sites_structs.Get_dtf_paths(),
			sites_structs.Get_sg_paths(),
			sites_structs.Get_pgd_paths(),
			sites_structs.Get_vg_paths(),
			sites_structs.Get_knb_paths(),
		}
		for _, site_path := range site_paths {
			scrappers.Get_articles(site_path.Site_link, site_path)
		}
		fmt.Println("Done!")
		time.Sleep(config.Sleep_time * time.Minute)
	}
}

//to do: handle i/o timeout error
