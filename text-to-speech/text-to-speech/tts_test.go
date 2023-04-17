package tts

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/hegedustibor/htgo-tts/handlers"
)

func TestTTS_Translate(t *testing.T) {
	tts := New()
	aa := "Niko is a god"
	l, s, err := tts.CreateTranslate(aa)
	println("SIZE " + (fmt.Sprint(len(aa))))
	if err != nil {
		t.Errorf("Failed to transalte text: %s, with err: %s", l, err)
	}
	print(s)
	mplayer := handlers.MPlayer{}
	err = mplayer.Play(l)
	if err != nil {
		t.Errorf("Failed to play file from location: %s, with err: %s", l, err)
	}

	e := os.Remove(l)
	if e != nil {
		log.Fatal(e)
	}
}
