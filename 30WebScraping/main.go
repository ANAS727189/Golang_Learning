package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gocolly/colly"
)

type News struct {
	Label    string `json:"label"`
	Title    string `json:"title"`
	SubTitle string `json:"subtitle"`
	Author   string `json:"author"`
	ImgUrl   string `json:"img_url"`
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var userAgents = []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Safari/605.1.15",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
	}

	c := colly.NewCollector(
		colly.AllowedDomains("thehindu.com", "www.thehindu.com"),
		colly.UserAgent(userAgents[rand.Intn(len(userAgents))]),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*.thehindu.com",
		Parallelism: 2,
		Delay:       3 * time.Second,
		RandomDelay: 2 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.9")
		r.Headers.Set("Referer", "https://www.thehindu.com/")
		fmt.Println("Visiting:", r.URL.String())
	})

	var news []News

	// Main news
	c.OnHTML("div.element.bigger.main-element", func(h *colly.HTMLElement) {
		newsItem := News{
			Label:    h.ChildText("div.Label"),
			Title:    h.ChildText("h1.title a"),
			SubTitle: h.ChildText("div.sub-text a"),
			Author:   h.ChildText("div.author-name a"),
			ImgUrl:   h.ChildAttr("img", "src"),
		}
		news = append(news, newsItem)
	})

	// Latest News
	c.OnHTML("div[data-ga-label='Latest News'] li.time-list", func(h *colly.HTMLElement) {
		timeText := h.ChildText("span.timePublished")
		region := h.ChildText("span.time")
		full := h.ChildText("a")

		title := full
		if idx := len(region); idx < len(full) {
			title = full[idx:]
		}

		newsItem := News{
			Label:    region,
			Title:    title,
			SubTitle: timeText,
		}
		news = append(news, newsItem)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error:", err)
	})

	c.OnScraped(func(r *colly.Response) {
		file, err := os.Create("output.json")
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ") 

		if err := encoder.Encode(news); err != nil {
			fmt.Println("Error writing JSON:", err)
		} else {
			fmt.Println("✅ Data saved to output.json")
		}
	})

	err := c.Visit("https://www.thehindu.com/")
	if err != nil {
		fmt.Println("Failed to visit URL:", err)
	}
}
