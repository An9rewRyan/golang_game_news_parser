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

type Worker struct {
	Link string `json:"link"`
	Err  error  `json:"err"`
}

// var Channel = make(chan int, config.MAX_amount_of_goroutines)

func Get_links(site_link string, site_paths root_structs.Article_paths) []string {
	fmt.Println("Getting links for" + site_link)
	no_links_msg := []string{"no new links"}
	var links []string
	var page *html.Node
	var err error
	var new_links []string
	if site_paths.Use_js_generated_pages {
		page_html := utils.Get_js_genetated_page(site_link)
		page, err = htmlquery.Parse(strings.NewReader(page_html))
		if err != nil {
			fmt.Println("error 1!")
			fmt.Println("error: ")
			fmt.Println(err)
			return links //let get_articles function to handle this
		}
	} else {
		page, err = htmlquery.LoadURL(site_link)
		if err != nil {
			fmt.Println("error 2!")
			fmt.Println("error: ")
			fmt.Println(err)
			return links //let get_articles function to handle this
		}
	}
	links_hex := htmlquery.Find(page, site_paths.Links_xpath)
	for _, link := range links_hex {
		links = append(links, utils.Add_domain_name(htmlquery.SelectAttr(link, "href"), site_paths))
	}
	fmt.Println(links)
	links_no_duplicates := utils.RemoveDuplicateStr(links)
	if len(links_no_duplicates) != 0 {
		links_no_duplicates = links_no_duplicates[0:3]
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
	fmt.Println("links:", links)
	if len(links) != 0 {
		if links[0] == "no new links" {
			config.Wg_main.Done()
			return articles
		}
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("(GA) Failed to load links for %s, tried for %d times. The link xpath is probably wrong.\n", site_link, amount_of_retries)
			fmt.Println(err)
		}
		fmt.Println("Ending get articles process for " + site_paths.Site_alias)
		config.Wg_main.Done()
	}()
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
	workerChan := make(chan *Worker, len(links))
	for _, link := range links {
		wk := &Worker{Link: link}
		go wk.Get_article(link, site_paths, workerChan)
	}
	passed := 0
	for passed != 2 {
		wk := <-workerChan
		fmt.Println("Wk is: ")
		fmt.Println(wk)

		if wk.Err == nil {
			passed++
		} else {
			wk.Err = nil
			fmt.Println("Got failed job, from " + wk.Link + ", retrying...")
			go wk.Get_article(wk.Link, site_paths, workerChan)
		}
	}
	fmt.Println("Articles for " + site_paths.Site_alias + " are loaded, but wait group is still not activated")
	Wg.Wait()

	fmt.Println("Sucessfully loaded articles for " + site_paths.Site_alias)
	return articles
}

func (wk *Worker) Get_article(link string, site_paths root_structs.Article_paths, workerChan chan<- *Worker) root_structs.Article {
	amount_of_retries := 0
	var article_html *html.Node
	var article root_structs.Article
	var err error
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("had troubles with article " + link)
			wk.Err = fmt.Errorf("Panic happened with %v", r)
			fmt.Println("Sending job back to loader ")
			fmt.Println("Wk: ")
			fmt.Println(wk)
		} else {
			fmt.Println("sucessfully loaded article " + link)
			Wg.Done()
		}
		workerChan <- wk
	}()

	article_html, err = utils.Switch_between_common_and_js_page(link, site_paths)
	if err != nil {
		fmt.Println(err)
	}

	//checking if parser got proper response (not 404 page)
	for utils.Get_element_by_xpath(article_html, site_paths.Error_code_xpath, "text") == site_paths.Error_message {
		if amount_of_retries == config.MAX_amount_of_loading_retries {
			error_data := "1001"
			fmt.Println(error_data)
			panic(error_data) //status code for max retries while loading article
		}
		article_html, err = utils.Switch_between_common_and_js_page(link, site_paths)
		if err != nil {
			fmt.Println(err)
		}
		amount_of_retries++
		fmt.Println("retrying to get: " + link)
		time.Sleep(config.Article_loading_retry_time * time.Millisecond)
	}
	amount_of_retries = 0
	for len(article.Title) == 0 || len(article.Content) == 0 || len(article.Pub_date) == 0 {
		fmt.Println(article.Title, article.Content, article.Pub_date)
		if amount_of_retries == config.MAX_amount_of_loading_retries {
			error_data := "1001"
			fmt.Println(error_data)
			panic(error_data) //status code for max retries while loading article
		}
		article_html, err = utils.Switch_between_common_and_js_page(link, site_paths)
		if err != nil {
			fmt.Println(err)
		}
		article = root_structs.Article{
			Title:       utils.Get_element_by_xpath(article_html, site_paths.Title_xpath, "title"),
			Content:     utils.Get_element_by_xpath(article_html, site_paths.Content_xpath, "content"),
			Image_url:   utils.Get_element_by_xpath(article_html, site_paths.Image_url_xpath, "image"),
			Pub_date:    utils.Get_element_by_xpath(article_html, site_paths.Pub_date_xpath, "pub_date"),
			Source_link: link,
			Site_alias:  site_paths.Site_alias,
		}
		amount_of_retries++
		fmt.Println("retrying to get: "+link+" for ", amount_of_retries)
		time.Sleep(config.Article_loading_retry_time * time.Millisecond)
	}
	article = formatters.Format_article(article, site_paths)
	fmt.Println("Going to load: "+link+" to db", amount_of_retries)
	utils.Write_article_to_db(article)

	return article
}
