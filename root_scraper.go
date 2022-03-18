package main

import (
	"fmt"
	"parser/root_structs"
	"strings"
	"sync"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

var wg sync.WaitGroup

func Get_links(site_link string, site_paths root_structs.Article_paths) []string {
	var links []string
	page, err := htmlquery.LoadURL(site_link)
	if err != nil {
		panic(`not a valid XPath expression.`)
	}
	links_hex := htmlquery.Find(page, site_paths.Links_xpath)
	for _, link := range links_hex {
		links = append(links, htmlquery.SelectAttr(link, "href"))
	}
	wg.Add(len(links))
	return links
}

func Get_articles(site_link string, site_paths root_structs.Article_paths) []root_structs.Article {
	var articles []root_structs.Article
	links := Get_links(site_link, site_paths)
	for _, link := range links {
		go Get_article(link, site_paths)
	}
	wg.Wait()
	return articles
}

func Get_article(link string, site_paths root_structs.Article_paths) root_structs.Article {
	defer wg.Done()
	article_html, err := htmlquery.LoadURL(link)
	if err != nil {
		fmt.Println(err)
	}
	article := root_structs.Article{
		Title:     Get_element_by_xpath(article_html, site_paths.Title_xpath, "title"),
		Content:   Get_element_by_xpath(article_html, site_paths.Content_xpath, "content"),
		Image_url: Get_element_by_xpath(article_html, site_paths.Image_url_xpath, "image"),
		Author:    Get_element_by_xpath(article_html, site_paths.Author_xpath, "author"),
		Pub_date:  Get_element_by_xpath(article_html, site_paths.Pub_date_xpath, "pub_date"),
	}
	fmt.Println(article)
	return article
}

func Get_element_by_xpath(page_html *html.Node, xpath string, elem_type string) string {
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
	elem_html := strings.Join(elems_html, "")
	return elem_html
}
