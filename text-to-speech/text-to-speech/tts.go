package tts

import (
	"log"
	"os"

	htgotts "github.com/hegedustibor/htgo-tts"
)

type TTS struct{}

func New() *TTS {
	return &TTS{}
}

func (tts *TTS) CreateTranslate(text string) (string, int64, error) {
	speech := htgotts.Speech{Folder: "audio", Language: "en"}
	localtion, err := speech.CreateSpeechFile(text, "tts")
	if err != nil {
		log.Fatalf("Failed to create speech file with err: %s", err)
	}

	fi, err := os.Stat(localtion)
	if err != nil {
		return "", 0, err
	}

	return localtion, fi.Size(), err
}
