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

	"github.com/antchfx/htmlquery"
	"github.com/tebeka/selenium"
	"golang.org/x/net/html"
)

func Get_js_genetated_page_node(link string) string {
	values := map[string]string{"link": link}
	jsonValue, _ := json.Marshal(values)
	resp, err := http.Post("http://js_parser:8000", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(bodyBytes)
}

func Get_js_genetated_page(link string) string {
	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(config.Selenium_path, config.Selenium_port, config.Selenium_opts...)
	if err != nil {
		fmt.Println(err)
	}
	defer service.Stop()
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", config.Selenium_port))
	if err != nil {
		fmt.Println(err)
	}

	defer wd.Quit()
	if err := wd.Get(link); err != nil {

		panic(err)
	}
	page, err := wd.PageSource()
	if err != nil {
		fmt.Println(err)
	}
	return page
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
