package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"parser/root_structs"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"golang.org/x/net/html"
)

var Wg sync.WaitGroup
var Channel = make(chan int, MAX_amount_of_goroutines)

func get_js_genetated_page_geziyor(link string) string {
	var page_html string
	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered(link, g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			page_html = string(r.Body)
		},
	}).Start()
	return page_html
}

func get_js_genetated_page(link string) string {
	values := map[string]string{"link": link}
	jsonValue, _ := json.Marshal(values)
	resp, err := http.Post("http://localhost:8000", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(bodyBytes)
}

func add_domain_name(link string, site_paths root_structs.Article_paths) string {
	re := regexp.MustCompile("//.*?/")
	match := re.FindStringSubmatch(site_paths.Site_link)
	domain_name := match[0]
	return "https:" + domain_name + link
}

func Get_links(site_link string, site_paths root_structs.Article_paths) []string {
	var links []string
	var page *html.Node
	var err error
	if site_paths.Use_js_generated_pages {
		page_html := get_js_genetated_page(site_link)
		page, err = htmlquery.Parse(strings.NewReader(page_html))
		if err != nil {
			fmt.Println(err)
		}
	} else {
		page, err = htmlquery.LoadURL(site_link)
		if err != nil {
			fmt.Println(err)
		}
	}
	links_hex := htmlquery.Find(page, site_paths.Links_xpath)
	for _, link := range links_hex {
		links = append(links, htmlquery.SelectAttr(link, "href"))
	}
	links_no_duplicates := removeDuplicateStr(links)
	links_no_duplicates = links_no_duplicates[0:5]
	fmt.Println(links_no_duplicates)
	return links_no_duplicates
}

func Get_articles(site_link string, site_paths root_structs.Article_paths) []root_structs.Article {
	var articles []root_structs.Article
	var links []string
	amount_of_retries := 0
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Failed to load links for %s, tried for %d times. The link xpath is probably wrong.\n", site_link, MAX_amount_of_loading_retries)
		}
	}()
	if site_paths.Use_js_generated_pages {
		Channel = make(chan int, 1)
	}
	for {
		if amount_of_retries == MAX_amount_of_loading_retries {
			panic("1000") //status code for max retries while loading links
		}
		if len(links) == 0 {
			amount_of_retries++
			time.Sleep(Links_loading_retry_time * time.Millisecond)
			fmt.Printf("tryed to load %s for %d amount of times, retrying...\n", site_link, amount_of_retries)
			links = Get_links(site_link, site_paths)
		} else {
			break
		}
	}
	Wg.Add(len(links))
	for _, link := range links {
		Channel <- 1
		go Get_article(link, site_paths)
	}
	Wg.Wait()
	Wg_global.Done()
	return articles
}

func Get_article(link string, site_paths root_structs.Article_paths) root_structs.Article {
	amount_of_retries := 0
	var article_html *html.Node
	var err error
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Failed to load source for %s, tried for %d times. The link xpath is probably wrong.\n", link, MAX_amount_of_loading_retries)
			<-Channel
			Wg.Done()
		}
	}()
	if !(strings.Contains(link, "https://")) {
		link = add_domain_name(link[1:], site_paths)
	}
	if site_paths.Use_js_generated_pages {
		article_plain := get_js_genetated_page(link)
		article_html, err = htmlquery.Parse(strings.NewReader(article_plain))
		if err != nil {
			fmt.Println(err)
		}
	} else {
		article_html, err = htmlquery.LoadURL(link)
		if err != nil {
			fmt.Println(err)
		}
	}
	for Get_element_by_xpath(article_html, site_paths.Error_code_xpath, "text") == site_paths.Error_message {
		if amount_of_retries == MAX_amount_of_loading_retries {
			error_data := "1001"
			panic(error_data) //status code for max retries while loading article
		}
		article_html, err = htmlquery.LoadURL(link)
		amount_of_retries++
		if err != nil {
			fmt.Printf("An error occured while reading htmlfile on %s \n", link)
		}
		fmt.Println("retrying to get: " + link)
		time.Sleep(Article_loading_retry_time * time.Millisecond)
	}
	var article = root_structs.Article{
		Title:     Get_element_by_xpath(article_html, site_paths.Title_xpath, "title"),
		Content:   Get_element_by_xpath(article_html, site_paths.Content_xpath, "content"),
		Image_url: Get_element_by_xpath(article_html, site_paths.Image_url_xpath, "image"),
		Pub_date:  Get_element_by_xpath(article_html, site_paths.Pub_date_xpath, "pub_date"),
	}
	fmt.Printf("%+v\n", article)
	<-Channel
	Wg.Done()
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
		if elem == nil && (elem_type == "image") {
			elems_html = append(elems_html, elem_type+" Not found")
			break
		}
		elems_html = append(elems_html, htmlquery.InnerText(elem))
	}
	elem_html := strings.Join(elems_html, "")
	return elem_html
}
