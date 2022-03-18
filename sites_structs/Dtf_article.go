package sites_structs

import (
	"parser/root_structs"
)

func Get_dtf_paths() root_structs.Article_paths {
	DTF_paths := root_structs.Article_paths{
		// Title_xpath: "//h1[@class='content-title']/descendant::text()[not(ancestor::span/@class='content-editorial-tick')]",
		Title_xpath: "//text()[not(ancestor::span/@class='content-editorial-tick') and ancestor::h1[@class='content-title']]",
		// Content_xpath: `//div[@class = 'content content--full ']/descendant::text()[ancestor::div/@class='l-island-a' and
		// 											not(ancestor::div/@class='andropov_link andropov_link--with_image') and
		// 											not(ancestor::div/@class='content content--embed')]`,
		Content_xpath:  "//text()[ancestor::div/@class = 'content content--full ' and not(ancestor::div/@class='content content--embed') and not(ancestor::div/@class='content__thanks') and not(ancestor::div/@class='content-info__item ') and not(ancestor::div/@class='content-info content-info--full l-island-a')]",
		Pub_date_xpath: "//time[@class='time']/@title",
		Image_url_xpath: `//div[@class='content-image content-image--wide']/div[1]/@data-image-src|
						 //div[@class='content-image']/div[1]/@data-image-src`,
		// Image_url_xpath: "//image/src[parent::div/@class'andropov_image__inner']",
		Author_xpath: "//a[@class='content-header-author content-header-author--user content-header__item']/div[@class='content-header-author__name']/text()",
	}
	return DTF_paths
}
