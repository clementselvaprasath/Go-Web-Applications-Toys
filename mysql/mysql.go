package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

// Init creates a singleton of DB connecting to local mysql instance.
func Init() {
	//
	// Configure the database connection (always check errors)
	d, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/mysql")

	if err != nil {
		log.Fatal(err)
	}

	// Initialize the first connection to the database, to see if everything works correctly.
	// Make sure to check the error.
	err = d.Ping()
	if err != nil {
		log.Fatal(err)
	}

	db = d
	fmt.Print("Succefully initialized MySQL DB\n")
}

// CreateTable simply creates a table "books"
func CreateTable() {
	query := `
	CREATE TABLE IF NOT EXISTS books (
		id INT AUTO_INCREMENT,
		name TEXT NOT NULL,
		price INT(11), 
		created_at DATETIME,
		PRIMARY KEY (id)
	);`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}

	fmt.Print("Succefully created table\n")
}

// InsertBook inserts a new book to the table.
func InsertBook(name string, price int) {
	createdAt := time.Now()

	result, err := db.Exec(`INSERT INTO books (name, price, created_at) VALUES (?, ?, ?)`, name, price, createdAt)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	fmt.Printf("Succefully insert a new book %v with price %v. Returned id: %v\n", name, price, id)
}

// ListBooks list all books we have
func ListBooks() {
	type book struct {
		id        int
		name      string
		price     int
		createdAt time.Time
	}
	rows, err := db.Query(`SELECT id, name, price, created_at FROM books`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var books []book
	for rows.Next() {
		var u book

		err := rows.Scan(&u.id, &u.name, &u.price, &u.createdAt)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", books)
}

// DeleteBook deletes a book given a name
func DeleteBook(name string) {
	_, err := db.Exec(`DELETE FROM books WHERE name = ?`, name)
	if err != nil {
		log.Fatal(err)
	}
}
