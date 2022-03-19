package sites_structs

import (
	"parser/root_structs"
)

func Get_vg_paths() root_structs.Article_paths {
	VG_paths := root_structs.Article_paths{
		Links_xpath: "//div[@class = 'item-name type0']/a[1]/@href",
		Title_xpath: "//h1[@class ='news_item_title']/text()",
		// Content_xpath: "//p/text() or //h2/text() or  //div[@class = 'v12']/ul[not(@class)]/li/text() or //div[@class = 'v12']/ul[not(@class)]/li/a/text() or  //p/a[@class='l_ks' or @class='gatx']/text()",
		Content_xpath:    "//text()[ancestor::div/@class = 'v12' and((parent::p or parent::h2 or parent::a/@class='l_ks' or parent::a/@class='gatx' or parent::span[not[@class]]) or ancestor::ul[not(@class)])]",
		Pub_date_xpath:   "//div[@class='news_item_date']/meta[1]/@content",
		Image_url_xpath:  "//div[@class='news_item_image_img']/img/@data-src",
		Site_link:        "https://vgtimes.ru/tags/%D0%98%D0%B3%D1%80%D0%BE%D0%B2%D1%8B%D0%B5+%D0%BD%D0%BE%D0%B2%D0%BE%D1%81%D1%82%D0%B8/",
		Error_code_xpath: "//div[@class='radius']/div[@class='cup ln']/text()",
	}
	return VG_paths
}
