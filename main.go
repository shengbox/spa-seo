package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
)

var target = flag.String("t", "https://www.cmvalue.com", "单页面应用域名")
var port = flag.Int("p", 8000, "本服务监听端口")

func main() {
	flag.Parse()
	cache := map[string]string{}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.String())
		if value, ok := cache[r.URL.String()]; ok {
			w.Write([]byte(value))
		} else {
			html, err := HttpHtmlContent(*target + r.URL.String())
			if err == nil {
				cache[r.URL.String()] = html
				w.Write([]byte(html))
			} else {
				w.Write([]byte(err.Error()))
			}
		}
	})
	log.Println("server start at port:", *port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), mux))
}

func HttpHtmlContent(url string) (htmlContent string, err error) {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true),
	}
	c, _ := chromedp.NewExecAllocator(context.Background(), options...)
	chromeCtx, cancel := chromedp.NewContext(c)
	timeoutCtx, cancel := context.WithTimeout(chromeCtx, 10*time.Second)
	defer cancel()

	err = chromedp.Run(timeoutCtx,
		chromedp.Navigate(url),
		chromedp.Sleep(time.Millisecond*500),
		chromedp.OuterHTML(`html`, &htmlContent),
	)
	return
}
