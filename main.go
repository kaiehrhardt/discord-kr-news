package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/bwmarrin/discordgo"
)

const (
	baseUrl = "https://www.nexon.com/kartdrift/de/news/"
)

var (
	discordToken     = os.Getenv("KRN_DISCORD_TOKEN")
	discordChannel   = os.Getenv("KRN_DISCORD_CHANNEL")
	checkInterval, _ = time.ParseDuration(os.Getenv("KRN_CHECK_INTERVAL"))
)

type News struct {
	Posts []Post
}

type Post struct {
	Url string
}

func (n *News) AddPost(url string) {
	n.Posts = append(n.Posts, Post{Url: url})
}

func (n *News) Print() {
	if len(n.Posts) <= 1 {
		log.Fatalln("Empty Posts?!")
	} else {
		log.Println(n.Posts)
	}
}

func Compare(oldNews, newNews *News) (diffUrls []string) {
	for i, post := range oldNews.Posts {
		if post.Url != newNews.Posts[i].Url {
			diffUrls = append(diffUrls, newNews.Posts[i].Url)
			log.Println("Diff found. Old: " + post.Url + " New: " + newNews.Posts[i].Url)
		}
	}
	return diffUrls
}

func GetNews() *News {
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
			n.AddPost(baseUrl + category + "/" + url)
		})
	}
	return n
}

func main() {
	dg, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = dg.ChannelMessageSend(discordChannel, "KR-News für Channel aktiviert.")
	if err != nil {
		log.Println(err)
	}

	initNews := GetNews()
	initNews.Print()

	for {
		newNews := GetNews()
		newNews.Print()
		diff := Compare(initNews, newNews)

		if len(diff) > 0 {
			for _, news := range diff {
				message := fmt.Sprintf("Es gibt eine neue News: %s", news)
				_, err = dg.ChannelMessageSend(discordChannel, message)
				if err != nil {
					log.Println(err)
				}
			}
		}

		time.Sleep(checkInterval)
	}
}
