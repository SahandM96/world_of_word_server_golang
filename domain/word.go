package domain

type Word struct {
	Id         string `json:"id,omitempty"`
	Tag        string `json:"tag"`
	Category   string `json:"category"`
	LastUpdate string `json:"last_update"`
	CreateDate string `json:"create_date"`
	Name       string `json:"word"`
	Publish    bool   `json:"publish"`
	ImageCat	string `json:"img_cat"`
}

type WordRepository interface {
	GetWords(number string) ([]Word, error)
	FindAll() ([]Word, error)
	FindById(id string) (Word, error)
	FindByName(name string) (Word, error)
	FindByTag(tag string) ([]Word, error)
	FindByCategory(category string) ([]Word, error)
	Create(word Word) (Word, error)
	Update(word Word) (Word, error)
}
