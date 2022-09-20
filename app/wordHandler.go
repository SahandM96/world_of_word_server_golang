package app

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sahandm96/word_game/domain"
	"github.com/sahandm96/word_game/service"
)

type WordHandler struct {
	service service.WordService
}

func (wh *WordHandler) getWords(w http.ResponseWriter, r *http.Request) {

	cList, err := wh.service.GetWords(json.NewDecoder(r.Body).Decode("number").Error())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(cList)
	if err != nil {
		return
	}
}

func (wh *WordHandler) getAllWord(w http.ResponseWriter, r *http.Request) {
	wList, err := wh.service.GetAllWords()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(wList)
	if err != nil {
		return
	}
}

func (wh WordHandler) getWordById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	word, err := wh.service.GetWordById(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(word)
	if err != nil {
		return
	}
}

func (wh WordHandler) getWordByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	words, err := wh.service.GetWordByName(params["name"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(words)
	if err != nil {
		return
	}
}
func (wh WordHandler) getWordByCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	words, err := wh.service.GetWordByCategory(params["category"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(words)
	if err != nil {
		return
	}
}
func (wh WordHandler) getWordByTag(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	words, err := wh.service.GetWordByName(params["tag"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(words)
	if err != nil {
		return
	}
}

func (wh WordHandler) create(w http.ResponseWriter, r *http.Request) {
	var wordM domain.Word
	wordM.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	wordM.LastUpdate = time.Now().Format("2006-01-02 15:04:05")

	err := json.NewDecoder(r.Body).Decode(&wordM)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	word, err := wh.service.CreateWord(wordM)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(word)
	if err != nil {
		return
	}

}

func (wh WordHandler) update(w http.ResponseWriter, r *http.Request) {
	var wordM domain.Word

	err := json.NewDecoder(r.Body).Decode(&wordM)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = wh.service.UpdateWord(wordM)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	word, err := wh.service.GetWordById(wordM.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(word)
	if err != nil {
		return
	}

}
