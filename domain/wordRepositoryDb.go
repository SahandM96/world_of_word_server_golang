package domain

import (
	"database/sql"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type WordRepositoryDb struct {
	db *sql.DB
}

func (d WordRepositoryDb) GetWords(number string) ([]Word, error) {
	getChunkOfDataSql := "SELECT * FROM words  LIMIT  ? ;"
	num, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		panic(err.Error())
	}
	rows, err := d.db.Query(getChunkOfDataSql, num)
	if err != nil {
		panic(err.Error())
	}
	words := make([]Word, 0)
	for rows.Next() {
		var w Word
		if err != nil {
			panic(err.Error())
		}
		words = append(words, w)
	}
	return words, nil
}
func (d WordRepositoryDb) FindAll() ([]Word, error) {
	findAllSql := "SELECT * From words"
	rows, err := d.db.Query(findAllSql)
	if err != nil {
		panic(err.Error())
	}
	words := make([]Word, 0)
	for rows.Next() {
		var w Word
		err := rows.Scan(&w.Id, &w.Name,
			&w.Tag,
			&w.Category,
			&w.Publish,
			&w.CreateDate,
			&w.LastUpdate,
			&w.ImageCat,
		)
		if err != nil {
			panic(err.Error())
		}
		words = append(words, w)
	}
	return words, nil
}
func (d WordRepositoryDb) Create(word Word) (Word, error) {
	createSql := "INSERT INTO words (name, tag,category,last_update,create_date,publish,img) VALUES (?,?,?,?,?,?,?)"
	wc, err := d.db.Exec(createSql, word.Name, word.Tag, word.Category, word.LastUpdate, word.CreateDate, word.Publish, word.ImageCat)
	if err != nil {
		panic(err.Error())
	}
	id, err := wc.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	word.Id = strconv.Itoa(int(id))
	return word, nil
}
func (d WordRepositoryDb) FindById(id string) (Word, error) {
	findByIdSql := "SELECT * FROM words WHERE id = ?"
	row := d.db.QueryRow(findByIdSql, id)
	var w Word
	err := row.Scan(&w.Id, &w.Name, &w.Tag, &w.Category, &w.Publish, &w.LastUpdate, &w.CreateDate, &w.ImageCat)
	if err != nil {
		panic(err.Error())
	}
	return w, nil
}
func (d WordRepositoryDb) FindByCategory(category string) ([]Word, error) {
	findByCategorySql := "SELECT * FROM words WHERE category = ?"
	rows, err := d.db.Query(findByCategorySql, category)
	if err != nil {
		panic(err.Error())
	}
	words := make([]Word, 0)
	for rows.Next() {
		var w Word
		if err != nil {
			panic(err.Error())
		}
		words = append(words, w)
	}
	return words, nil
}
func (d WordRepositoryDb) FindByTag(tag string) ([]Word, error) {
	findByTagSql := "SELECT * FROM words WHERE tag = ?"
	rows, err := d.db.Query(findByTagSql, tag)
	if err != nil {
		panic(err.Error())
	}
	words := make([]Word, 0)
	for rows.Next() {
		var w Word
		if err != nil {
			panic(err.Error())
		}
		words = append(words, w)
	}
	return words, nil
}
func (d WordRepositoryDb) FindByName(name string) (Word, error) {
	findByNameSql := "SELECT * FROM words WHERE name = ?"
	row := d.db.QueryRow(findByNameSql, name)
	var w Word
	err := row.Scan(&w.Id, &w.Name, &w.Tag, &w.Category, &w.LastUpdate, &w.CreateDate, &w.Publish, &w.ImageCat)
	if err != nil {
		panic(err.Error())
	}
	return w, nil
}
func (d WordRepositoryDb) Update(word Word) (Word, error) {
	updateSql := "UPDATE words SET name = ?, tag = ?, category = ?, publish = ?, last_update = ?, create_date = ? ,img= ? WHERE id = ? "
	updateDate := time.Now().Format("2006-01-02 15:04:05")

	_, err := d.db.Exec(updateSql, word.Name, word.Tag, word.Category, word.Publish, updateDate, word.CreateDate, word.ImageCat, word.Id)
	if err != nil {
		panic(err.Error())
	}

	return word, nil
}
func NewWordRepositoryDb() WordRepositoryDb {
	db, err := sql.Open("mysql", "--:--@tcp(localhost:3306)/word_game")
	if err != nil {
		panic(err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	return WordRepositoryDb{db}
}
