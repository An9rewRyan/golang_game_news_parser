package scrappers

import (
	"fmt"
	"parser/config"
	"parser/formatters"
	"parser/root_structs"
	"parser/utils"
	"strings"
	"sync"
	"time"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

var Wg sync.WaitGroup
var Channel = make(chan int, config.MAX_amount_of_goroutines)

func Get_links(site_link string, site_paths root_structs.Article_paths) []string {
	no_links_msg := []string{"no new links"}
	var links []string
	var page *html.Node
	var err error
	var new_links []string
	if site_paths.Use_js_generated_pages {
		page_html := utils.Get_js_genetated_page(site_link)
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
		links = append(links, utils.Add_domain_name(htmlquery.SelectAttr(link, "href"), site_paths))
	}
	links_no_duplicates := utils.RemoveDuplicateStr(links)
	if len(links_no_duplicates) != 0 {
		links_no_duplicates = links_no_duplicates[0:2]
	}
	fmt.Println(links_no_duplicates)

	for _, link := range links_no_duplicates {
		if utils.Check_if_article_in_db(link) == false {
			fmt.Println(utils.Check_if_article_in_db(link))
			new_links = append(new_links, link)
		}
	}
	fmt.Println(new_links)
	if len(links_no_duplicates) > 0 && len(new_links) == 0 { //it basically means that there are no new posts published
		return no_links_msg
	}
	return new_links
}

func Get_articles(site_link string, site_paths root_structs.Article_paths) []root_structs.Article {
	var articles []root_structs.Article
	amount_of_retries := 0
	links := Get_links(site_link, site_paths)
	//it means that there are no new articles
	if links[0] == "no new links" {
		return articles
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Failed to load links for %s, tried for %d times. The link xpath is probably wrong.\n", site_link, amount_of_retries)
			fmt.Println(err)
		}
	}()
	//reducing channel capacity, because js rendering server has troubles with concurrency (yet)
	if site_paths.Use_js_generated_pages {
		Channel = make(chan int, 1)
	}
	for {
		if amount_of_retries == config.MAX_amount_of_loading_retries {
			panic("1000") //status code for max retries while loading links
		}
		if len(links) == 0 {
			amount_of_retries++
			time.Sleep(config.Links_loading_retry_time * time.Millisecond)
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
	return articles
}

func Get_article(link string, site_paths root_structs.Article_paths) root_structs.Article {
	amount_of_retries := 0
	var article_html *html.Node
	var err error
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Failed to load source for %s, tried for %d times. The link xpath is probably wrong.\n", link, amount_of_retries)
			fmt.Println(err)
			<-Channel
			Wg.Done()
		}
	}()

	if site_paths.Use_js_generated_pages {
		article_plain := utils.Get_js_genetated_page(link)
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
	//checking if parser got proper response (not 404 page)
	for utils.Get_element_by_xpath(article_html, site_paths.Error_code_xpath, "text") == site_paths.Error_message {
		if amount_of_retries == config.MAX_amount_of_loading_retries {
			error_data := "1001"
			panic(error_data) //status code for max retries while loading article
		}
		if site_paths.Use_js_generated_pages {
			article_plain := utils.Get_js_genetated_page(link)
			article_html, err = htmlquery.Parse(strings.NewReader(article_plain))
			if err != nil {
				fmt.Println(err)
			}
		} else {
			article_html, err = htmlquery.LoadURL(link)
			if err != nil {
				fmt.Printf("An error occured while reading htmlfile on %s \n", link)
			}
		}
		amount_of_retries++
		fmt.Println("retrying to get: " + link)
		time.Sleep(config.Article_loading_retry_time * time.Millisecond)
	}
	var article = root_structs.Article{
		Title:       utils.Get_element_by_xpath(article_html, site_paths.Title_xpath, "title"),
		Content:     utils.Get_element_by_xpath(article_html, site_paths.Content_xpath, "content"),
		Image_url:   utils.Get_element_by_xpath(article_html, site_paths.Image_url_xpath, "image"),
		Pub_date:    utils.Get_element_by_xpath(article_html, site_paths.Pub_date_xpath, "pub_date"),
		Source_link: link,
		Site_alias:  site_paths.Site_alias,
	}
	article = formatters.Format_article(article, site_paths)
	utils.Write_article_to_db(article)

	<-Channel
	Wg.Done()

	return article
}
