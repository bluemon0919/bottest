package init

import (
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)


func init() {
	// Setup HTTP Server for receiving requests from LINE platform.
	http.HandleFunc("/callback", handler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	bot, err := linebot.New(
		os.Getenv("1a5df2f9ef2381b290c3954b1ff0dd3f"),
		os.Getenv("WRyh+ZsToMynMG5EoCJR8StGTCdLi6y78WeEIxfZo7G1MQiIOuihTaZbf47BDFsy34zd7gtHht2aXiCPuL4Ub07jlom36MMTDP9NIusH7dgtDPzlwX0xNwOH0SYmMmrxrOGMPYd6rRI65kljAy8+xQdB04t89/1O/w1cDnyilFU="),
		linebot.WithHTTPClient(urlfetch.Client(ctx)),
	)
	if err != nil {
		log.Errorf(ctx, "LINE bot new error.: %v", err)
		return
	}


	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					log.Errorf(ctx, "LINE bot reply message error.: %v", err)
				}
			}
		}
	}
}