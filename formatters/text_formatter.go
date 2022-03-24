package formatters

import "regexp"

func Delete_whitespaces(title string) string {
	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	title_formatted := re_leadclose_whtsp.ReplaceAllString(title, "")
	title_formatted = re_inside_whtsp.ReplaceAllString(title_formatted, " ")
	return title_formatted
}

func Format_title(title string) string {
	title_formatted := Delete_whitespaces(title)
	return title_formatted
}

func Format_content(content string) string {
	content_formatted := Delete_whitespaces(content)
	return content_formatted
}
