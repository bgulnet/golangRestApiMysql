package Book

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	id            int `gorm:"primaryKey"`
	ISBN          string
	Author_id     uint
	Publisher_id  uint
	Catalog_id    uint
	Title         string
	Language_id   uint
	Numberofpages uint
	Piece         uint
	Status        uint
}

var dsn string
var db *gorm.DB
var err error
var book Book

func ConnectionDB() {
	host := os.Getenv("MYSQL_DB_HOST")
	dbName := os.Getenv("MYSQL_DB")
	user := os.Getenv("MYSQL_DB_USER")
	pass := os.Getenv("MYSQL_DB_PASS")
	dsn = user + ":" + pass + "@tcp(" + host + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Book{})

}

// kitap ekleme
func CreateBookModel(book Book) {
	ConnectionDB()
	db.Create(&Book{ISBN: book.ISBN, Author_id: book.Author_id, Publisher_id: book.Publisher_id, Catalog_id: book.Catalog_id, Title: book.Title,
		Language_id: book.Language_id, Numberofpages: book.Numberofpages, Piece: book.Piece, Status: book.Status})
}
func UpdateBookModel(params string, book Book) {
	ConnectionDB()

	db.Model(&Book{}).Where("id = ?", params).Updates(&Book{ISBN: book.ISBN, Author_id: book.Author_id, Publisher_id: book.Publisher_id, Catalog_id: book.Catalog_id, Title: book.Title,
		Language_id: book.Language_id, Numberofpages: book.Numberofpages, Piece: book.Piece, Status: book.Status})

}
func DeleteBookModel(id int) {
	ConnectionDB()
	var book Book

	db.Where("id = ?", id).Delete(&book)
}
func SelectAll() []Book {
	ConnectionDB()
	var book []Book
	db.Raw("select * from books").Scan(&book)
	return book

}
