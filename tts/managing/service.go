package managing

import (
	"fmt"
)

type TTS interface {
	CreateTranslate(text string) (string, int64, error)
}

type Service struct {
	tts TTS
}

func NewService(tts TTS) *Service {
	return &Service{
		tts: tts,
	}
}

func (s *Service) Translate(text string) (string, int64, error) {
	l, size, err := s.tts.CreateTranslate(text)
	if err != nil {
		return "", 0, fmt.Errorf("Service.tts.CreateTranslate failed - %v", err)
	}
	return l, size, nil
}
