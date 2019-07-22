package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"regexp"
	"strings"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ExampleScrape() {
	res, err := http.Get("https://www.3dmgame.com/news/201907/3766333.html")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)

	}
	//var result string
	//connect()
	doc.Find(".news_warp_center>p").Each(
		func(i int, s *goquery.Selection) {
			var name string
			var money string
			text := s.Text()
			re := regexp.MustCompile(`([0-9]|10)\.[^0-9].+[0-9][元,万,亿]`)
			match := re.FindString(text)
			if match != "" {
				match := strings.SplitAfterN(match, ".", 2)
				index := strings.Index(match[1], " ")
				if index != -1 {
					name = match[1][:index]
					money = match[1][index:]
				} else {
					index := strings.Index(match[1], "，")
					name = match[1][:index]
					money = match[1][index:]
				}
				mysql(name, money)
			}
		})
}

func main() {
	ExampleScrape()
}
