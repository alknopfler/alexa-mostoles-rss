package function

import (
	"context"
	"github.com/ericdaugherty/alexa-skills-kit-golang"
	gf "github.com/mmcdole/gofeed"
	cfg "github.com/alknopfler/alexa-mostoles-rss/config"
	"log"
)

func NoticiasHoy(feed *gf.Feed,context context.Context, request *alexa.Request, session *alexa.Session, aContext *alexa.Context, response *alexa.Response) {
	log.Println("Get Noticias Hoy")

	result := cfg.SpeechBeginNoticias
	for _,val := range feed.Items{
		result += val.Title + " . "

	}
	response.SetStandardCard(cfg.CardTitle, result, cfg.ImageSmall, cfg.ImageLong)
	response.SetOutputText(result)
	response.ShouldSessionEnd = true
}