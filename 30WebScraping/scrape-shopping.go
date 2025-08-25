package main

import (
	"fmt"

	"math/rand"
	"time"

	"github.com/gocolly/colly"
)

type Item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"img_url"`
}

func scrape() {
	rand.Seed(time.Now().UnixNano())
	var userAgents = []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Safari/605.1.15",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
	}
	c := colly.NewCollector(
		colly.AllowedDomains("www.macys.com"),
		colly.UserAgent(userAgents[rand.Intn(len(userAgents))]),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*.macys.com",
		Parallelism: 2,
		Delay:       3 * time.Second,
		RandomDelay: 2 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.9")
		r.Headers.Set("Referer", "https://www.macys.com/")
		fmt.Println("Visiting:", r.URL.String())
	})

	var items []Item
	c.OnHTML("li.sortablegrid-product", func(h *colly.HTMLElement) {
		item := Item{
			Name:   h.ChildText("h3.product-name"),
			Price:  h.ChildText("span.price-reg"),
			ImgUrl: h.ChildAttr("img.picture-image", "src"),
		}
		items = append(items, item)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	c.OnScraped(func(r *colly.Response) {
		for _, item := range items {
			fmt.Printf("Name: %s\nPrice: %s\nImage URL: %s\n\n", item.Name, item.Price, item.ImgUrl)
		}
	})

	err := c.Visit("https://www.macys.com/shop/new-trending/new-at-macys/mens-new-arrivals/mens-shoes-new-arrivals?id=166743")
	if err != nil {
		fmt.Println("Failed to visit URL:", err)
	}
}
