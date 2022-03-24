package formatters

import (
	"regexp"
	"strings"
)

func Delete_whitespaces(title string) string {
	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	title_formatted := re_leadclose_whtsp.ReplaceAllString(title, "")
	title_formatted = re_inside_whtsp.ReplaceAllString(title_formatted, " ")
	return title_formatted
}

func Format_single_commas(text string) string {
	return strings.Replace(text, "'", "''", -1)
}

func Format_title(title string) string {
	title_formatted := Delete_whitespaces(title)
	title_formatted = Format_single_commas(title_formatted)
	return title_formatted
}

func Format_content(content string) string {
	content_formatted := Delete_whitespaces(content)
	content_formatted = Format_single_commas(content_formatted)
	return content_formatted
}
