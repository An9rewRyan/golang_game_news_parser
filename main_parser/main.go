package main

import (
	"fmt"
	"parser/config"
	"parser/root_structs"
	"parser/scrappers"
	"parser/sites_structs"
	"time"
)

func main() {
	//uncomment when launching for first time and comment again afterwards
	// utils.create_articles_table()
	// utils.create_recently_loaded_articles_table()

	for {
		site_paths := []root_structs.Article_paths{
			// sites_structs.Get_igrm_paths(), //shut down cause this site is down
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
