package main

import (
	"fmt"

	"parser/root_structs"
	"parser/sites_structs"

	"golang.org/x/net/html"

	"github.com/antchfx/htmlquery"
)

func get_links(site_link string, site_paths root_structs.Article_paths) []string {
	var links []string
	page, err := htmlquery.LoadURL(site_link)
	if err != nil {
		panic(`not a valid XPath expression.`)
	}
	links_hex := htmlquery.Find(page, site_paths.Links_xpath)
	for _, link := range links_hex {
		links = append(links, htmlquery.SelectAttr(link, "href"))
	}
	return links
}

func get_articles(site_link string, site_paths root_structs.Article_paths) []root_structs.Article {
	var articles []root_structs.Article
	links := get_links(site_link, site_paths)
	for _, link := range links {
		articles = append(articles, get_article(link, site_paths))
	}
	return articles
}

func get_article(link string, site_paths root_structs.Article_paths) root_structs.Article {
	article_html, err := htmlquery.LoadURL(link)
	if err != nil {
		fmt.Println(err)
	}
	article := root_structs.Article{
		Title:     get_element_by_xpath(article_html, site_paths.Title_xpath, "title"),
		Content:   get_element_by_xpath(article_html, site_paths.Content_xpath, "content"),
		Image_url: get_element_by_xpath(article_html, site_paths.Image_url_xpath, "image"),
		Author:    get_element_by_xpath(article_html, site_paths.Author_xpath, "author"),
		Pub_date:  get_element_by_xpath(article_html, site_paths.Pub_date_xpath, "pub_date"),
	}
	return article
}

func get_element_by_xpath(page_html *html.Node, xpath string, elem_type string) []string {
	var elems []*html.Node
	if elem_type == "image" || elem_type == "author" || elem_type == "pub_date" {
		elems = append(elems, htmlquery.FindOne(page_html, xpath))
	} else {
		elems = htmlquery.Find(page_html, xpath)
	}
	var elems_html []string
	for _, elem := range elems {
		elems_html = append(elems_html, htmlquery.OutputHTML(elem, true))
	}
	return elems_html
}

func main() {
	site_paths := sites_structs.Get_dtf_paths()
	fmt.Println(get_articles(site_paths.Site_link, site_paths))
}
