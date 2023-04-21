package main

import (
	"log"
	"net/http"
	"shorten-url-go/tts/managing"
	tts "shorten-url-go/tts/text-to-speech"
)

func main() {
	text2Speech := tts.New()
	s := managing.NewService(text2Speech)
	t := managing.NewTransport(s)

	http.HandleFunc("/tts", t.HandleTranslate)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
