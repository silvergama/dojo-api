package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/silvergama/dojo-api/entity"
)

var db *sql.DB

func Setup() error {
	sqlite, err := sql.Open("sqlite3", "data/dojo.db")
	if err != nil {
		return err
	}

	db = sqlite
	return nil
}

func GetUsers() ([]*entity.User, error) {
	var result []*entity.User
	queryString := "select * from user"

	rows, err := db.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u entity.User
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Phone)
		if err != nil {
			return nil, err
		}
		result = append(result, &u)
	}
	return result, nil
}

func AddUser(u entity.User) error {
	stmt, err := db.Prepare(`
			insert into user(id, first_name, last_name, email, phone) values(?, ?, ?, ?, ?)
		`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.ID, u.FirstName, u.LastName, u.Email, u.Phone)
	if err != nil {
		return err
	}
	return nil
}
