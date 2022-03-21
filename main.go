package main

import (
	"fmt"
	"parser/sites_structs"
)

func main() {
	site_paths := sites_structs.Get_sg_paths()
	articles := Get_articles(site_paths.Site_link, site_paths)
	fmt.Println(articles)
}
