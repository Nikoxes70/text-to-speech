package managing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type transporter struct {
	service Servicer
}

type TranslateRequest struct {
	Text string
}

type TranslateResponse struct {
	Size int64 `json:"size"`
}

type Servicer interface {
	Translate(text string) (string, int64, error)
}

func NewTransport(s Servicer) *transporter {
	return &transporter{
		s,
	}
}

func (t transporter) HandleTranslate(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/tts" || r.Method != http.MethodPost {
		http.Error(w, "404 not found.", http.StatusMethodNotAllowed)
		return
	}

	var reqBody map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "server error - failed to parse request body", http.StatusInternalServerError)
		return
	}

	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "text cannot be empty", http.StatusBadRequest)
		return
	}

	l, s, err := t.service.Translate(text)
	if err != nil {
		http.Error(w, "server error - failed to create short url", http.StatusInternalServerError)
		return
	}

	responseBody := TranslateResponse{
		Size: s,
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(l))
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, l)

	b, err := json.Marshal(responseBody)
	if err != nil {
		http.Error(w, "server error - failed to create short url", 500)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Printf("failed to martshal response - %v\n", err)
	}

	return
}
