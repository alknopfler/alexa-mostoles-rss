package main

import (
	"github.com/mmcdole/gofeed"
	"fmt"
)

func main (){
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://www.mostoles.es/mostoles/cm/ayto-mostoles/rss?locale=es_ES&rssContent=978")
	fmt.Println(feed.Title)
	fmt.Println(feed.Description)
	for _,val := range feed.Items {
		fmt.Println(val.Title)
		fmt.Println(val.PublishedParsed.Date())
	}
}

