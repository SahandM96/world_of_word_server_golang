package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sahandm96/word_game/domain"
	"github.com/sahandm96/word_game/service"
)

func Start() {
	router := mux.NewRouter()
	wHandler := WordHandler{
		service.NewWordService(domain.NewWordRepositoryDb()),
	}
	router.HandleFunc("/words", wHandler.create).Methods("Post")
	router.HandleFunc("/getWords", wHandler.getWords).Methods("Post")
	router.HandleFunc("/words", wHandler.getAllWord).Methods("Get")
	router.HandleFunc("/wordsByCategory", wHandler.getWordByCategory).Methods("Post")
	router.HandleFunc("/wordsByName", wHandler.getWordByName).Methods("Post")
	router.HandleFunc("/wordsById", wHandler.getWordById).Methods("Post")
	router.HandleFunc("/wordsByTag", wHandler.getWordByTag).Methods("Post")
	router.HandleFunc("/wordUpdate", wHandler.update).Methods("Post")

	log.Fatal(http.ListenAndServe("0.0.0.0:6969", router))
}
