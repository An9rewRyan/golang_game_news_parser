package scrappers

import (
	"fmt"
	"parser/config"
	"parser/errors"
	"parser/formatters"
	"parser/root_structs"
	"parser/utils"
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

func Get_links(site_link string, site_paths root_structs.Article_paths) ([]string, error) {
	fmt.Println("Getting links for" + site_link)
	no_links_msg := []string{"no new links"}
	var links []string
	var page *html.Node
	var err error
	var new_links []string
	page, err = utils.Switch_between_common_and_js_page(site_link, site_paths)
	if err != nil {
		fmt.Println(err)
		return links, err
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

	for _, link := range links_no_duplicates {
		if utils.Check_if_article_in_db(link) == false {
			fmt.Println(utils.Check_if_article_in_db(link))
			new_links = append(new_links, link)
		}
	}
	fmt.Println(new_links)
	err = nil
	if len(links_no_duplicates) > 0 && len(new_links) == 0 { //it basically means that there are no new posts published
		return no_links_msg, err
	}
	return new_links, err
}

func Get_articles(site_link string, site_paths root_structs.Article_paths) []root_structs.Article {
	var articles []root_structs.Article
	amount_of_retries := 0
	links, _ := Get_links(site_link, site_paths)
	var err error
	// err = nil

	fmt.Println("links:", links)
	if len(links) != 0 {
		if links[0] == "no new links" {
			fmt.Println("No ne links found for ", site_paths.Site_alias)
			config.Wg_main.Done()
			return articles
		}
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("Ending get articles process for " + site_paths.Site_alias)
		// if site_paths.Site_alias != "knb" && site_paths.Site_alias != "pgd" {

		// }
	}()
	for {
		if amount_of_retries == config.MAX_amount_of_loading_retries {
			panic(errors.New_max_load_retry_error(fmt.Sprintf("Tryed to get %s for %d times, didn`t work!")))
		}
		amount_of_retries++
		time.Sleep(config.Links_loading_retry_time * time.Millisecond)
		fmt.Printf("tryed to load %s for %d amount of times, retrying...\n", site_link, amount_of_retries)
		links, err = Get_links(site_link, site_paths)
		fmt.Println(len(links), err, links)
		if links[0] == "no new links" {
			fmt.Println("No ne links found for ", site_paths.Site_alias)
			config.Wg_main.Done()
			return articles
		}
		if len(links) != 0 && err == nil && links[0] != "no new links" {
			fmt.Println("Here!")
			fmt.Println(links, err)
			break
		}
		fmt.Println(links, err)
	}
	Wg.Add(len(links))
	fmt.Println("Links len is ", len(links))
	fmt.Println("Starting links iteration for site ", site_paths.Site_alias)
	workerChan := make(chan *Worker, len(links))
	fmt.Println("Created worker chan")
	for _, link := range links {
		fmt.Println("Iterating links!")
		wk := &Worker{Link: link}
		go wk.Get_article(link, site_paths, workerChan)
		fmt.Println("Started handling ", link)
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
	config.Wg_main.Done()
	return articles
}

func (wk *Worker) Get_article(link string, site_paths root_structs.Article_paths, workerChan chan<- *Worker) (root_structs.Article, error) {
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
	fmt.Println("Getting html for article ", link)
	article_html, err = utils.Switch_between_common_and_js_page(link, site_paths)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Got html for article ", link)
	//checking if parser got proper response (not 404 page)
	for utils.Get_element_by_xpath(article_html, site_paths.Error_code_xpath, "text") == site_paths.Error_message {

		amount_of_retries++
		fmt.Println("retrying to get: " + link)

		if amount_of_retries == config.MAX_amount_of_loading_retries {
			panic(errors.New_max_load_retry_error(fmt.Sprintf("Tryed to get %s for %d times, didn`t work!")))
		}
		time.Sleep(config.Article_loading_retry_time * time.Millisecond)

		article_html, err = utils.Switch_between_common_and_js_page(link, site_paths)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	amount_of_retries = 0
	for len(article.Title) == 0 || len(article.Content) == 0 || len(article.Pub_date) == 0 {
		amount_of_retries++
		fmt.Println("retrying to get: "+link+" for ", amount_of_retries)
		if amount_of_retries == config.MAX_amount_of_loading_retries {
			panic(errors.New_max_load_retry_error(fmt.Sprintf("Tryed to get %s for %d times, didn`t work!"))) //status code for max retries while loading article
		}
		time.Sleep(config.Article_loading_retry_time * time.Millisecond)
		article_html, err = utils.Switch_between_common_and_js_page(link, site_paths)
		if err != nil {
			fmt.Println(err)
			continue
		}
		article = root_structs.Article{
			Title:       utils.Get_element_by_xpath(article_html, site_paths.Title_xpath, "title"),
			Content:     utils.Get_element_by_xpath(article_html, site_paths.Content_xpath, "content"),
			Image_url:   utils.Get_element_by_xpath(article_html, site_paths.Image_url_xpath, "image"),
			Pub_date:    utils.Get_element_by_xpath(article_html, site_paths.Pub_date_xpath, "pub_date"),
			Source_link: link,
			Site_alias:  site_paths.Site_alias,
		}
	}
	article = formatters.Format_article(article, site_paths)
	fmt.Println("Going to load: "+link+" to db", amount_of_retries)
	utils.Write_article_to_db(article)

	return article, err
}
