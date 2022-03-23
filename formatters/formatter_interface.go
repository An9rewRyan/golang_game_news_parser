package formatters

import (
	"parser/root_structs"
	"regexp"
)

// type Article_formatter interface {
// 	Format_title(title string) string
// 	Format_content(content string) string
// 	Format_pub_date(pub_data string) string
// 	Format_image_url(image_url string) string
// }

func Delete_whitespaces(title string) string {
	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	title_formatted := re_leadclose_whtsp.ReplaceAllString(title, "")
	title_formatted = re_inside_whtsp.ReplaceAllString(title_formatted, " ")
	return title_formatted
}

func Format_article(article root_structs.Article, article_info root_structs.Article_paths) root_structs.Article {
	// var title, content, image_url, pub_date string
	// site := article_info.Site_alias
	article_formatted := root_structs.Article{
		Title:     Format_title(article.Title),
		Content:   Format_content(article.Content),
		Image_url: Format_image_url(article.Image_url),
		Pub_date:  Format_pub_date(article.Pub_date),
	}
	// switch site {
	// case "dtf":
	// case "vg":
	// case "igrm":
	// case "sg":
	// case "knb":
	// case "pgd":
	// }
	return article_formatted
}

func Format_title(title string) string {
	title_formatted := Delete_whitespaces(title)
	return title_formatted
}

func Format_content(content string) string {
	content_formatted := Delete_whitespaces(content)
	return content_formatted
}

func Format_image_url(image_url string) string {
	return image_url
}

func Format_pub_date(pub_date string) string {

	return pub_date
}
