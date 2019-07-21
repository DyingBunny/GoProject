package main

import (
	"fmt"
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
	doc.Find(".news_warp_center>p").Each(
		func(i int, s *goquery.Selection) {
			text := s.Text()
			re := regexp.MustCompile(`([0-9]|10)\.[^0-9].+[0-9][元,万,亿]`)
			match := re.FindString(text)
			if match != "" {
				match := strings.SplitAfterN(match, ".", 2)
				re := regexp.MustCompile(``)
				result := re.FindString(match[1][])
				fmt.Println(result)
			}
		})
}

func main() {
	ExampleScrape()
}
