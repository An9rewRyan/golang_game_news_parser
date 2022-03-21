package sites_structs

import (
	"parser/root_structs"
)

func Get_knb_paths() root_structs.Article_paths {
	KNB_paths := root_structs.Article_paths{
		Links_xpath: `//a[@class='fabric_absolute_body__OxcZb' or
									 @class='style_body__r9G3R' or 
									 @class='style_body__H1Cks' or
									 @class='SmallTextCard_body__f7Khf']`,
		Title_xpath:            "//h1/text()",
		Content_xpath:          "//text()[ancestor::div[starts-with(@class, 'material-content']]",
		Pub_date_xpath:         "//script[@type='application/ld+json']/text()",
		Image_url_xpath:        "//link[@rel='preload'][@as='image']/@href",
		Site_link:              "https://kanobu.ru/videogames/",
		Error_code_xpath:       "//div[@class='knb-404_title']/text()",
		Error_message:          "Страница не найдена!",
		Use_js_generated_pages: true,
	}
	return KNB_paths
}
