package models

import (
	"database/sql"
	"errors"
	"time"
)

type KodesBoxInterface interface {
	Insert(title string, content string, expires int) (int, error)
	Get(id int) (*Kode, error)
	Latest() ([]*Kode, error)
}

type Kode struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expired time.Time
}

type KodesBox struct {
	DB *sql.DB
}

func (box *KodesBox) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO kodes (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	res, err := box.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (box *KodesBox) Get(id int) (*Kode, error) {

	stmt := `SELECT id, title, content, created, expires FROM
    kodes WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := box.DB.QueryRow(stmt, id)

	kode := &Kode{}

	err := row.Scan(&kode.ID, &kode.Title, &kode.Content, &kode.Created, &kode.Expired)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}

	return kode, nil
}

func (box *KodesBox) Latest() ([]*Kode, error) {

	stmt := `SELECT id, title, content, created, expires FROM kodes
    WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	rows, err := box.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	kodes := []*Kode{}

	for rows.Next() {
		kode := &Kode{}

		err = rows.Scan(&kode.ID, &kode.Title, &kode.Content, &kode.Created, &kode.Expired)
		if err != nil {
			return nil, err
		}

		kodes = append(kodes, kode)
	}

	if err = rows.Err(); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}

	return kodes, nil
}
