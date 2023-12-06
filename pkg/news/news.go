package news

import (
	"encoding/json"
	"io"

	"github.com/vrtineu/news/pkg/api"
)

type NewsService struct {
	apiClient *api.Client
}

type NewsResponse struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []Article
}

type Article struct {
	Source      Source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	UrlToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

type Source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewNewsService(client *api.Client) *NewsService {
	return &NewsService{
		apiClient: client,
	}
}

func (s *NewsService) GetTopHeadlines() ([]Article, error) {
	res, err := s.apiClient.Get("/top-headlines?country=us")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var newsResponse NewsResponse
	err = json.Unmarshal(body, &newsResponse)
	if err != nil {
		return nil, err
	}

	return newsResponse.Articles, nil
}
