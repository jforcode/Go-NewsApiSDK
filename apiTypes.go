package newsApi

import "time"

type ApiError struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (apiErr ApiError) Error() string {
	return apiErr.Code + " - " + apiErr.Message
}

type ApiSource struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Category    string `json:"category"`
	Language    string `json:"language"`
	Country     string `json:"country"`
}

type ApiArticleSource struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ApiArticle struct {
	Source      ApiArticleSource `json:"source"`
	Author      string           `json:"author"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	URL         string           `json:"url"`
	URLToImage  string           `json:"urlToImage"`
	PublishedAt time.Time        `json:"publishedAt"`
}

type ApiSourcesResponse struct {
	Status  string       `json:"status"`
	Sources []*ApiSource `json:"sources"`
}

type ApiArticlesResponse struct {
	Status       string        `json:"status"`
	TotalResults int           `json:"totalResults"`
	Articles     []*ApiArticle `json:"articles"`
}
