package main

import (
	"url_shortener/app"
	"url_shortener/interfaces"
)

func main() {
	store := interfaces.New()
	generator := app.NewUrlGenerator("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_", 10)
	urlService := app.NewUrlService(store, generator, "localhost:8000/")

	controller := interfaces.NewController(urlService)
	controller.Run(":8000")
}
