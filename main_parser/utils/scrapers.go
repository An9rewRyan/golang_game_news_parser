package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"parser/config"
	"parser/root_structs"
	"regexp"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func Get_js_genetated_page(link string) string {
	values := map[string]string{"link": link}
	jsonValue, _ := json.Marshal(values)
	var resp *http.Response
	var err error
	no_err := false
	for no_err == false {
		fmt.Println("Sent request to: " + link)
		resp, err = http.Post("http://js_parser:8000", "application/json", bytes.NewBuffer(jsonValue))
		fmt.Println("Got result from: " + link)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Retryng to request to js_parser with link " + link)
			time.Sleep(config.Links_loading_retry_time * time.Millisecond)
		} else {
			no_err = true
			fmt.Println(link + " loaded sucessfull!:)")
		}
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(bodyBytes)
}

func Add_domain_name(link string, site_paths root_structs.Article_paths) string {
	if !(strings.Contains(link, "https://")) {
		re := regexp.MustCompile("//.*?/")
		match := re.FindStringSubmatch(site_paths.Site_link)
		domain_name := match[0]
		return "https:" + domain_name + link[1:]
	}
	return link
}

func Get_element_by_xpath(page_html *html.Node, xpath string, elem_type string) string {
	var elems []*html.Node
	if elem_type == "image" || elem_type == "pub_date" {
		found_elem := htmlquery.FindOne(page_html, xpath)
		elems = append(elems, found_elem)
	} else {
		found_elems := htmlquery.Find(page_html, xpath)
		elems = found_elems
	}
	var elems_html []string
	for _, elem := range elems {
		if elem == nil && (elem_type == "image") {
			elems_html = append(elems_html, elem_type+" Not found")
			break
		}
		elems_html = append(elems_html, htmlquery.InnerText(elem))
	}
	elem_html := strings.Join(elems_html, "")
	return elem_html
}

func Switch_between_common_and_js_page(link string, site_paths root_structs.Article_paths) (*html.Node, error) {
	var article_html *html.Node
	var err error
	if site_paths.Use_js_generated_pages {
		article_plain := Get_js_genetated_page(link)
		article_html, err = htmlquery.Parse(strings.NewReader(article_plain))
		if err != nil {
			fmt.Println(err)
		}
	} else {
		article_html, err = htmlquery.LoadURL(link)
		if err != nil {
			fmt.Printf("An error occured while reading htmlfile on %s \n", link)
			fmt.Println(err)
		}
	}
	return article_html, err
}
