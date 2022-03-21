package root_structs

type Article_paths struct {
	Links_xpath            string `json:"link_xpath"`
	Title_xpath            string `json:"title_xpath"`
	Content_xpath          string `json:"content_xpath"`
	Pub_date_xpath         string `json:"pub_data_xpath"`
	Image_url_xpath        string `json:"image_url_xpath"`
	Site_link              string `json:"site_link"`
	Error_code_xpath       string `json:"error_code_xpath"`
	Error_message          string `json:"error_message"`
	Use_js_generated_pages bool   `json:"use_js_generated_pages"`
}
