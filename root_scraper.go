package main

import (
	"fmt"
	"parser/root_structs"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

const MAX = 10

var wg sync.WaitGroup
var channel = make(chan int, MAX)

func add_domain_name(link string, site_paths root_structs.Article_paths) string {
	re := regexp.MustCompile("//.*?/")
	match := re.FindStringSubmatch(site_paths.Site_link)
	domain_name := match[0]
	return "https:" + domain_name + link
}

func Get_links(site_link string, site_paths root_structs.Article_paths) []string {
	var links []string
	page, err := htmlquery.LoadURL(site_link)
	if err != nil {
		fmt.Println(`not a valid XPath expression.`)
	}
	links_hex := htmlquery.Find(page, site_paths.Links_xpath)
	for _, link := range links_hex {
		links = append(links, htmlquery.SelectAttr(link, "href"))
	}
	fmt.Println("links!!", links, len(links))
	return links
}

func Get_articles(site_link string, site_paths root_structs.Article_paths) []root_structs.Article {
	var articles []root_structs.Article
	var links []string

	for {
		if len(links) == 0 {
			links = Get_links(site_link, site_paths)
		} else {
			break
		}
	}
	wg.Add(len(links))
	for i, link := range links {
		channel <- 1
		go Get_article(link, site_paths)
		fmt.Println(i)
	}
	wg.Wait()
	return articles
}

func Get_article(link string, site_paths root_structs.Article_paths) root_structs.Article {
	defer wg.Done()
	if !(strings.Contains(link, "https://")) {
		link = add_domain_name(link[1:], site_paths)
	}
	article_html, err := htmlquery.LoadURL(link)
	if err != nil {
		fmt.Println("An error occured while reading htmlfile on " + link)
	}
	for Get_element_by_xpath(article_html, site_paths.Error_code_xpath, "text") != "" {
		article_html, err = htmlquery.LoadURL(link)
		if err != nil {
			fmt.Println("An error occured while reading htmlfile on " + link)
		}
		time.Sleep(2 * time.Second)
	}
	var article = root_structs.Article{}
	article = root_structs.Article{
		Title:     Get_element_by_xpath(article_html, site_paths.Title_xpath, "title"),
		Content:   Get_element_by_xpath(article_html, site_paths.Content_xpath, "content"),
		Image_url: Get_element_by_xpath(article_html, site_paths.Image_url_xpath, "image"),
		Pub_date:  Get_element_by_xpath(article_html, site_paths.Pub_date_xpath, "pub_date"),
	}
	fmt.Printf("%+v\n", article)
	fmt.Println(len(article.Title))
	<-channel
	return article
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
		if elem == nil && (elem_type == "image" || elem_type == "pub_date") {
			elems_html = append(elems_html, elem_type+" Not found")
			break
		}
		elems_html = append(elems_html, htmlquery.InnerText(elem))
	}
	elem_html := strings.Join(elems_html, "")
	return elem_html
}
