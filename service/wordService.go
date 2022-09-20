package service

import "github.com/sahandm96/word_game/domain"

type WordService interface {
	GetWords(number string) ([]domain.Word, error)
	GetAllWords() ([]domain.Word, error)
	GetWordById(id string) (domain.Word, error)
	GetWordByName(name string) (domain.Word, error)
	GetWordByCategory(category string) ([]domain.Word, error)
	GetWordByTag(tag string) ([]domain.Word, error)
	UpdateWord(word domain.Word) (domain.Word, error)
	CreateWord(word domain.Word) (domain.Word, error)
}

type DefaultWordService struct {
	repo domain.WordRepository
}

func (s DefaultWordService) GetWords(number string) ([]domain.Word, error) {
	return s.repo.GetWords(number)
}

func (s DefaultWordService) GetAllWords() ([]domain.Word, error) {
	return s.repo.FindAll()
}

func (s DefaultWordService) GetWordById(id string) (domain.Word, error) {
	return s.repo.FindById(id)
}
func (s DefaultWordService) GetWordByName(name string) (domain.Word, error) {
	return s.repo.FindByName(name)
}
func (s DefaultWordService) GetWordByCategory(category string) ([]domain.Word, error) {
	return s.repo.FindByCategory(category)
}
func (s DefaultWordService) GetWordByTag(tag string) ([]domain.Word, error) {
	return s.repo.FindByTag(tag)
}
func (s DefaultWordService) UpdateWord(word domain.Word) (domain.Word, error) {
	return s.repo.Update(word)
}
func (s DefaultWordService) CreateWord(word domain.Word) (domain.Word, error) {
	return s.repo.Create(word)
}
func NewWordService(repo domain.WordRepository) WordService {
	return DefaultWordService{repo}
}
