package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
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
	var result string
	tmp := "\n"
	doc.Find(".news_warp_center").Each(
		func(i int, s *goquery.Selection) {
			text := s.Find("p").Text()
			result += text
			result += tmp
		})
	fmt.Println(result)
}

func main() {
	ExampleScrape()
}
