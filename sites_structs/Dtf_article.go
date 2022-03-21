package sites_structs

import (
	"parser/root_structs"
)

func Get_dtf_paths() root_structs.Article_paths {
	DTF_paths := root_structs.Article_paths{
		Links_xpath:    "//a[@class= 'content-link']/@href",
		Title_xpath:    "//text()[not(ancestor::span/@class='content-editorial-tick') and ancestor::h1[@class='content-title']]",
		Content_xpath:  "//text()[ancestor::div/@class = 'content content--full ' and not(ancestor::div/@class='content content--embed') and not(ancestor::div/@class='content__thanks') and not(ancestor::div/@class='content-info__item ') and not(ancestor::div/@class='content-info content-info--full l-island-a')]",
		Pub_date_xpath: "//time[@class='time']/@title",
		Image_url_xpath: `//div[@class='content-image content-image--wide']/div[1]/@data-image-src|
						 div[@class='content-image']/div[1]/@data-image-src`,
		Site_link:              "https://dtf.ru/gameindustry",
		Error_code_xpath:       "//div[@class='error__code t-ff-1-700']/text()",
		Error_message:          "Ошибка 404",
		Use_js_generated_pages: false,
	}
	return DTF_paths
}
