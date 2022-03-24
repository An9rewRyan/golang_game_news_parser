package formatters

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func Format_pub_date(pub_date string) string {
	if strings.Contains(pub_date, "{") {
		var result map[string]string
		json.Unmarshal([]byte(pub_date), &result)
		pub_date = result["datePublished"]
	}
	var pub_date_formatted string
	pub_date_type_1 := `\d+\-\d+\-\d+\w+\d+\:\d+\:\d+\+\d+\:\d\d`
	pub_date_type_2 := `\d+\.\d+\.\d+\s\d+\:\d+\:\d+\s\(.+\)?`
	is_type_1, err := regexp.Match(pub_date_type_1, []byte(pub_date))
	if err != nil {
		fmt.Println(err)
	}
	is_type_2, err := regexp.Match(pub_date_type_2, []byte(pub_date))
	if err != nil {
		fmt.Println(err)
	}
	if is_type_1 {
		pub_date_formatted = Format_pub_date_type_1(pub_date)
	}
	if is_type_2 {
		pub_date_formatted = Format_pub_date_type_2(pub_date)
	}
	return pub_date_formatted
}

func Format_pub_date_type_1(pub_date string) string {
	symbol := regexp.MustCompile(`[a-zA-Z]`)
	time_code := regexp.MustCompile(`\+.+`)
	pub_date = symbol.ReplaceAllString(pub_date, " ")
	pub_date = time_code.ReplaceAllString(pub_date, "")
	return pub_date
}

func Format_pub_date_type_2(pub_date string) string {
	symbol := regexp.MustCompile(`\s`)
	symbol_splitter := symbol.FindString(pub_date)
	pub_date_splitted := strings.Split(pub_date, symbol_splitter)
	date_splitted := strings.Split(pub_date_splitted[0], ".")
	date_formatted := date_splitted[2] + "-" + date_splitted[1] + "-" + date_splitted[0]
	time := pub_date_splitted[1]
	return date_formatted + " " + time
}
