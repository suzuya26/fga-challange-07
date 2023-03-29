package repository

import (
	"fmt"
	"sesi_8/model"
	"sesi_8/repository/query"
)

// interface employee, index repo book
type BookRepo interface {
	CreateBook(in model.Book) (res model.Book, err error)
	GetAllBook() (res []model.Book, err error)
	GetBookById(bookID int) (res model.Book, err error)
	UpdateBook(bookID int, in model.Book) (res string, err error)
	DeleteBook(bookID int) (res string, err error)
}

func (r Repo) CreateBook(in model.Book) (res model.Book, err error) {
	err = r.db.QueryRow(
		query.CreateBook,
		in.Title,
		in.Author,
		in.Description,
	).Scan(
		&res.ID,
		&res.Title,
		&res.Author,
		&res.Description,
	)
	if err != nil {
		return res, err
	}
	return res, err
}

func (r Repo) GetAllBook() (res []model.Book, err error) {
	rows, err := r.db.Query(
		query.GetAllBook,
	)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		var book model.Book
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description)
		if err != nil {
			return res, err
		}
		res = append(res, book)
	}
	return res, err
}

func (r Repo) GetBookById(bookID int) (res model.Book, err error) {
	err = r.db.QueryRow(
		query.GetBook,
		bookID,
	).Scan(
		&res.ID,
		&res.Title,
		&res.Author,
		&res.Description,
	)
	if err != nil {
		return res, err
	}
	return res, err
}

func (r Repo) UpdateBook(bookID int, in model.Book) (res string, err error) {
	_, err = r.db.Exec(
		query.UpdateBook,
		bookID,
		in.Title,
		in.Author,
		in.Description,
	)

	if err != nil {
		return res, err
	}
	res = fmt.Sprintf("Book dengan Id %v sudah ter-update", bookID)
	return res, err
}

func (r Repo) DeleteBook(bookID int) (res string, err error) {
	_, err = r.db.Exec(
		query.DeleteBook,
		bookID,
	)

	if err != nil {
		return res, err
	}
	res = fmt.Sprintf("Book dengan Id %v sudah ter-hapus :<", bookID)
	return res, err
}
