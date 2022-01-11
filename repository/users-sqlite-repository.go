package repository

import (
	"database/sql"
	"demo/entity"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteRepo struct{}

type UserItem struct {
	id        string
	uuid      string
	firstname string
	lastname  string
}

//NewUserSQLiteRepository
func NewUserSQLiteRepository() UserRepository {
	os.Remove("./database.db") // is for demonstration only

	file, err := os.Create("./database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()

	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createUserTableSQL := `CREATE TABLE user (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"uuid" TEXT,
		"firstname" TEXT,
		"lastname" TEXT		
	  );`

	log.Println("Create user table...")
	statement, err := db.Prepare(createUserTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("user table created")

	return &sqliteRepo{}
}

func (*sqliteRepo) Save(user *entity.User) (*entity.User, error) {
	log.Println("Save sqlite 1")
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	stmt, err := db.Prepare("INSERT INTO `user` (uuid, firstname, lastname) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id, user.Firstname, user.Lastname)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return user, nil
}

func (*sqliteRepo) FindAll() ([]entity.User, error) {
	log.Println("FindAll sqlite")

	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM `user`")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := []entity.User{}

	for rows.Next() { // Iterate and fetch the records from result cursor
		item := UserItem{}
		err2 := rows.Scan(&item.id, &item.uuid, &item.firstname, &item.lastname)
		if err2 != nil {
			panic(err2)
		}
		user := entity.User{
			Id:        item.uuid,
			Firstname: item.firstname,
			Lastname:  item.lastname,
		}
		users = append(users, user)
	}

	return users, nil
}
