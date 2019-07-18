package main

import (
	"fmt"
	"github.com/jackdanger/collectlinks"
	"net/http"
)

func main() {
	res, _ := http.Get("https://www.dytt8.net/html/gndy/jddy/20190712/58842.html")
	links := collectlinks.All(res.Body)
	fmt.Println(links)
}
