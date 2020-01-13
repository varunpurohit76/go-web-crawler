package data_object

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type urlDBO struct {
	Id   sql.NullString
	Data sql.NullString
}

func createUrlDb(tx *sqlx.Tx, u *Url) (string, error) {
	if tx == nil {
		return "", errors.New("no tx")
	}
	id := sql.NullString{
		String: u.Id,
		Valid:  u.Id != "",
	}
	data := sql.NullString{
		String: u.Link,
		Valid:  u.Link != "",
	}
	_, err := tx.Exec("INSERT INTO url (id, data) VALUES(?, ?)", id, data)
	if err != nil {
		return "", err
	}
	return u.Id, nil
}

func getUrlDb(tx *sqlx.Tx, id string) (*Url, error) {
	if tx == nil {
		return nil, errors.New("no tx")
	}
	urlDBO := []urlDBO{}
	err := tx.Select(&urlDBO, "SELECT * FROM url WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	if len(urlDBO) != 1 {
		return nil, errors.New("gt 1 for id")
	}
	return &Url{
		Id: urlDBO[0].Id.String,
		Link: urlDBO[0].Data.String,
	}, nil
}
