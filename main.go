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
	for {
		site_paths := []root_structs.Article_paths{
			sites_structs.Get_igrm_paths(),
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
