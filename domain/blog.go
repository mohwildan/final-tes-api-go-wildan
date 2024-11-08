package domain



type BlogRequest struct {
	Title string `json:"title"`
	UrlImage string `json:"url_image"`
	Content string `json:"content"`
}