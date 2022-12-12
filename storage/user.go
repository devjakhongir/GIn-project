package storage

import (
	"database/sql"

	"app/models"
)

func Create(db *sql.DB, user models.User) (string, error) {

	var (
		id    string
		query string
	)

	query = `
		INSERT INTO 
			users (first_name, last_name)
		VALUES ( $1, $2 )
		RETURNING id
	`
	err := db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func GetById(db *sql.DB, id string) (models.User, error) {

	var (
		user  models.User
		query string
	)

	query = `
		SELECT
			id,
			first_name,
			last_name
		FROM
			users
		WHERE id = $1
	`
	err := db.QueryRow(
		query,
		id,
	).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
	)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetList(db *sql.DB) ([]models.User, error) {

	var (
		users []models.User
		query string
	)

	query = `
		SELECT
			id,
			first_name,
			last_name
		FROM
			users
	`
	rows, err := db.Query(query)

	if err != nil {
		return []models.User{}, err
	}

	for rows.Next() {
		var user models.User

		err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
		)

		if err != nil {
			return []models.User{}, err
		}

		users = append(users, user)
	}

	return users, nil
}


func Update(db *sql.DB, user models.User) (int64, error) {

	var (
		query string
	)

	query = `
		UPDATE 
			users 
		SET 
			first_name = $2, 
			last_name = $3
		WHERE 
			id = $1
	`

	result, err := db.Exec(
		query,
		user.Id,
		user.FirstName,
		user.LastName,
	)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func Delete(db *sql.DB, id string) error {

	_, err := db.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

func Update_patch(db *sql.DB, user models.User) (int64, error) {

	var (
		query1 string
		query string
		old_user models.User
	)

	query1 = `
		SELECT
			first_name,
			last_name
		FROM
			users
	`
	err := db.QueryRow(
		query1,
	).Scan(
		&old_user.FirstName,
		&old_user.LastName,
	)

	if err != nil {
		return 0, err
	}

	query = `
		UPDATE 
			users 
		SET 
			first_name = $2, 
			last_name = $3
		WHERE 
			id = $1
	`
	if user.FirstName ==""{
		user.FirstName=old_user.FirstName
	}

	if user.LastName ==""{
		user.LastName=old_user.LastName
	}

	result, err := db.Exec(
		query,
		user.Id,
		user.FirstName,
		user.LastName,
	)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
