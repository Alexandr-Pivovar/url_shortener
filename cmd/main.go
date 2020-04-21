package main

import (
	"flag"
	"url_shortener/app"
	"url_shortener/interfaces"
)

func main() {
	domain := flag.String("d", "localhost:8000/", "-d <addr:port/>")
	gen := flag.String("g", "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_", "-g abcABC")
	lenOfGen := flag.Int("i", 10, "-i 10")
	addr := flag.String("a", ":8000", "-a :80")

	flag.Parse()

	store := interfaces.New()
	generator := app.NewUrlGenerator(*gen, *lenOfGen)
	urlService := app.NewUrlService(store, generator, *domain)

	controller := interfaces.NewController(urlService)
	controller.Run(*addr)
}
