package root_structs

type Article struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Pub_date    string `json:"pub_data"`
	Image_url   string `json:"image_url"`
	Source_link string `json:"source_link"`
	Site_alias  string `json:"site_alias"`
}
