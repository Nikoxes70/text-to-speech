package tts

import (
	"log"
	"os"
	"strings"

	htgotts "github.com/hegedustibor/htgo-tts"
)

const MaxFileSize = 20

type TTS struct{}

func New() *TTS {
	return &TTS{}
}

func (tts *TTS) CreateTranslate(text string) (string, int64, error) {
	speech := htgotts.Speech{Folder: "audio", Language: "en"}
	if len(text) > MaxFileSize {
		text = text[:MaxFileSize]
	}
	fileName := strings.Replace(text, " ", "_", 0)
	fileLocation, err := speech.CreateSpeechFile(text, "tts_"+fileName)
	if err != nil {
		log.Fatalf("Failed to create speech file with err: %s", err)
	}

	fi, err := os.Stat(fileLocation)
	if err != nil {
		return "", 0, err
	}

	return fileLocation, fi.Size(), err
}
