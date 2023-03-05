package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

const baseAnnouncementsUrl = "https://www.nexon.com/kartdrift/de/news/announcement"

type News struct {
	Posts []Post
}

type Post struct {
	Url string
}

func (n *News) AddPosts(url string) {
	n.Posts = append(n.Posts, Post{Url: url})
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
	n := &News{}
	res, err := http.Get(baseAnnouncementsUrl + "/list")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	htmlTokens := html.NewTokenizer(res.Body)
loop:
	for {
		tt := htmlTokens.Next()
		switch tt {
		case html.ErrorToken:
			break loop
		case html.StartTagToken:
			t := htmlTokens.Token()
			isAnchor := t.Data == "a"
			if isAnchor {
				for _, v := range t.Attr {
					if v.Key == "href" && strings.HasPrefix(v.Val, "view") {
						n.AddPosts(baseAnnouncementsUrl + v.Val)
					}
				}
			}
		}
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
