package sites_structs

import (
	"parser/root_structs"
)

func Get_sg_paths() root_structs.Article_paths {
	SG_paths := root_structs.Article_paths{
		Links_xpath:            "//a[@class='article-image']/@href",
		Title_xpath:            "//h1[@class = 'article-title']/text()",
		Content_xpath:          "//text()[ancestor::section/@class='article']",
		Pub_date_xpath:         "//script[@type='application/ld+json']",
		Image_url_xpath:        "//section[@class='article']/descendant::img/@src | //div[@class='iframe_h']/data-iframe/@data-preroll-thumb",
		Site_link:              "https://stopgame.ru/news/industry",
		Error_code_xpath:       "//div[@class='error-code']/text()",
		Error_message:          "404",
		Use_js_generated_pages: false,
		Site_alias:             "sg",
	}
	return SG_paths
}
