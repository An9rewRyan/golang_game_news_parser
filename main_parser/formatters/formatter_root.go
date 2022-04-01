package formatters

import (
	"parser/root_structs"
)

func Format_article(article root_structs.Article, article_info root_structs.Article_paths) root_structs.Article {
	article_formatted := root_structs.Article{
		Title:       Format_title(article.Title),
		Content:     Format_content(article.Content),
		Image_url:   Format_image_url(article.Image_url),
		Pub_date:    Format_pub_date(article.Pub_date),
		Source_link: Format_single_commas(article.Source_link),
		Site_alias:  article.Site_alias,
	}

	return article_formatted
}
