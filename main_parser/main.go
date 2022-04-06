package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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
		site_paths_common := []root_structs.Article_paths{
			sites_structs.Get_dtf_paths(),
			sites_structs.Get_sg_paths(),
			sites_structs.Get_vg_paths(),
			sites_structs.Get_igrm_paths(),
			sites_structs.Get_knb_paths(),
			sites_structs.Get_pgd_paths(),
		}
		// site_paths_js := []root_structs.Article_paths{
		// 	sites_structs.Get_knb_paths(),
		// 	sites_structs.Get_pgd_paths(),
		// }
		config.Wg_main.Add(len(site_paths_common))
		// for _, site_path := range site_paths_js {
		// 	scrappers.Get_articles(site_path.Site_link, site_path)
		// }
		for _, site_path := range site_paths_common {
			go scrappers.Get_articles(site_path.Site_link, site_path)
		}
		config.Wg_main.Wait()
		message := "finished loading links, thank u, dear puppeteer server!"
		fmt.Println(message)
		//section below is used for telling js server that we loaded all links sucessfully and it should destroy previous browser instance for saving memory
		values := map[string]string{"link": message}
		jsonValue, _ := json.Marshal(values)
		_, _ = http.Post("http://js_parser:8000", "application/json", bytes.NewBuffer(jsonValue))
		/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		time.Sleep(config.Sleep_time * time.Minute)
	}
}
