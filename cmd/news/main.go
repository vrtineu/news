package main

import (
	"fmt"
	"log"

	"github.com/vrtineu/news/pkg/api"
	"github.com/vrtineu/news/pkg/config"
	"github.com/vrtineu/news/pkg/news"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	apiClient := api.NewClient(cfg.NewsAPIURL, cfg.NewsAPIKey)
	newsService := news.NewNewsService(apiClient)

	newsItems, err := newsService.GetTopHeadlines()
	if err != nil {
		log.Fatalf("Error fetching news: %v", err)
	}

	for _, item := range newsItems {
		fmt.Printf("Title: %s\nDescription: %s\nURL: %s\n\n", item.Title, item.Description, item.Url)
	}
}
