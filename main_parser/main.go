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
	// utils.Create_articles_table()
	// utils.Create_recently_loaded_articles_table()

	for {
		site_paths := []root_structs.Article_paths{
			sites_structs.Get_knb_paths(),
			sites_structs.Get_pgd_paths(),
			sites_structs.Get_dtf_paths(),
			sites_structs.Get_sg_paths(),
			sites_structs.Get_vg_paths(),
			sites_structs.Get_igrm_paths(),
		}
		config.Wg_main.Add(len(site_paths))
		// workerChan := make(chan *root_structs.Worker, len(site_paths)*2)
		for _, site_path := range site_paths {
			// wk := &root_structs.worker{id: i}
			go scrappers.Get_articles(site_path.Site_link, site_path)
		}
		config.Wg_main.Wait()
		fmt.Println("Done!")
		time.Sleep(config.Sleep_time * time.Minute)
	}
}
