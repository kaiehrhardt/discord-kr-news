package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseUrl = "https://www.nexon.com/kartdrift/de/news/"
)

type News struct {
	Posts []Post
}

type Post struct {
	Url      string
	Category string
}

func (n *News) AddPost(url string, category string) {
	n.Posts = append(n.Posts, Post{Url: url, Category: category})
}

func (n *News) Print() error {
	if len(n.Posts) <= 1 {
		return errors.New("Empty Posts?!")
	}
	for _, p := range n.Posts {
		fmt.Println(p.Url)
	}
	return nil
}

func Init() *News {
	var categories = []string{"announcement", "update", "ingameevent", "communityevent"}
	n := &News{}
	for _, category := range categories {
		res, err := http.Get(baseUrl + category + "/list")
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		doc.Find(".board_list").Each(func(i int, s *goquery.Selection) {
			url, _ := s.Find("a").Attr("href")
			n.AddPost(baseUrl+category+"/"+url, category)
		})
	}
	return n
}

func main() {
	n := Init()
	err := n.Print()
	if err != nil {
		log.Fatal(err)
	}
}
