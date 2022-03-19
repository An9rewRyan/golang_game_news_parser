package sites_structs

import (
	"parser/root_structs"
)

func Get_igrm_paths() root_structs.Article_paths {
	DTF_paths := root_structs.Article_paths{
		Links_xpath: "//a[@class='aubli_img']/@href",
		Title_xpath: "//h1[@class ='page_news_ttl haveselect']/text()",
		// Content_xpath: `//div[@class = 'universal_content clearfix']/descendant::text()[(ancestor::div[not(@class)] or ancestor::li) and
		// 															not(ancestor::div[@class='uninote console'])]`,
		Content_xpath:   "//text()[ancestor::div/@class='universal_content clearfix' and not(ancestor::div[@class='uninote console'])]",
		Pub_date_xpath:  "//div[@class='page_news noselect']/meta[@itemprop = 'datePublished']/@content",
		Image_url_xpath: "//div[@class ='main_pic_container']/img/@src",
		Site_link:       "https://www.igromania.ru/news/game/",
	}
	return DTF_paths
}
