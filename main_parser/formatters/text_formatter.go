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

func Basic_text_formatting(text string) string {
	text_formatted := Delete_whitespaces(text)
	text_formatted = Format_single_commas(text_formatted)
	text_formatted = strings.Replace(text_formatted, "\u00a0", " ", -1) //for deleting unbrakable space
	return text_formatted
}

func Format_title(title string) string {
	return Basic_text_formatting(title)
}

func Format_content(content string) string {
	return Basic_text_formatting(content)
}
