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

var target = flag.String("t", "", "单页面应用域名")
var port = flag.Int("p", 8000, "本服务监听端口")

func main() {
	flag.Parse()
	cache := map[string]string{}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if *target == "" {
			*target = "https://" + r.Host
		}
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
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36"),
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
