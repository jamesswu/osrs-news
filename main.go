package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

func LoadEnv() error {
	if err := godotenv.Load; err != nil {
		fmt.Println("no env file found")
	}
	return nil
}

func main() {
	fmt.Println("hello, world")

	err := LoadEnv()
	if err != nil {
		panic(err)
	}

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting: ", r.URL.String())
	})

	c.OnHTML(".news-list-article", func(e *colly.HTMLElement) {
		fmt.Println("url: ", e.ChildAttr("a", "href"))
		fmt.Println("title: ", e.ChildText(".news-list-article__title-link"))
		fmt.Println("category: ", e.ChildText(".news-list-article__category"))
		fmt.Println("date: ", e.ChildText(".news-list-article__date"))
		fmt.Println("desc: ", e.ChildText(".news-list-article__summary"))
		fmt.Println("image: ", e.ChildAttr(".news-list-article__figure-img", "src"))
	})

	c.Visit("https://secure.runescape.com/m=news/archive?oldschool=1")
}
