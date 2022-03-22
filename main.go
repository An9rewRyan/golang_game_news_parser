package mains

import (
	"fmt"
	"parser/root_structs"
	"parser/sites_structs"
	"sync"
)

var Wg_global sync.WaitGroup

func main() {
	site_paths := []root_structs.Article_paths{
		sites_structs.Get_igrm_paths(),
		sites_structs.Get_dtf_paths(),
		sites_structs.Get_sg_paths(),
		sites_structs.Get_pgd_paths(),
		sites_structs.Get_vg_paths(),
	}
	Wg_global.Add(len(site_paths))
	for _, site_path := range site_paths {
		go Get_articles(site_path.Site_link, site_path)
	}
	Wg_global.Wait()
	fmt.Println("Done!")
}
