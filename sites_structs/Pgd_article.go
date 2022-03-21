package sites_structs

import (
	"parser/root_structs"
)

func Get_pgd_paths() root_structs.Article_paths {
	PGD_paths := root_structs.Article_paths{
		Links_xpath:            "//div[@class='post-title']/a/@href",
		Title_xpath:            "//h1[@class = 'post-title']/text()",
		Content_xpath:          "//div[@class ='article-content js-post-item-content']/descendant::text()",
		Pub_date_xpath:         "//div[@class='post-metadata']/div/time/@datetime",
		Image_url_xpath:        "//figure/a/@href | //figure/img/@src | //div[@class='ytp-cued-thumbnail-overlay-image']/@style",
		Site_link:              "https://www.playground.ru/news/industry",
		Error_code_xpath:       "//div[@class='module-title']/text()",
		Error_message:          "Страница не найдена",
		Use_js_generated_pages: true,
	}
	return PGD_paths
}
