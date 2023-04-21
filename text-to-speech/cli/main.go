package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"shorten-url-go/text-to-speech/managing"
	tts "shorten-url-go/text-to-speech/text-to-speech"
	"strings"
)

func main() {
	for true {
		fmt.Println("Enter the text you would like to speech")

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		text2Speech := tts.New()
		s := managing.NewService(text2Speech)
		fmt.Println("Generating")
		location, size, err := s.Translate(text)
		if err != nil {
			log.Fatal(err)
		}

		path := strings.Split(location, "/")
		cmd := exec.Command("open", path[0])

		err = cmd.Run()

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Speech ready -> " + location + ", size: " + fmt.Sprint(size))
	}
}
