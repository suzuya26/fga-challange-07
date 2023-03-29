package query

const (
	CreateBook = `
		INSERT INTO 
			library.books
			(
				title,
				author,
				description
			)
			VALUES($1, $2, $3)
				RETURNING *;
	`
	GetBook = `
		SELECT * FROM library.books
			WHERE id=$1;
	`
	GetAllBook = `
		SELECT * FROM library.books
	`
	UpdateBook = `
		UPDATE library.books
			SET 
				title = $2, 
				author = $3, 
				description = $4
			WHERE
				id = $1;
	`
	DeleteBook = `
		DELETE FROM library.books 
			WHERE id = $1;

	`
)
