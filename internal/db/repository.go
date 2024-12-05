package db

import (
	"database/sql"
	"sample/internal/models"
)

func GetUsers(db *sql.DB) ([]models.User, error) {
	query := "SELECT [Id],[UserName],[Password],[CreatedDate],[IsAdmin] FROM [dbo].[Users]"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UserId, &user.UserName, &user.Password, &user.CreatedDate, &user.IsAdmin); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUser(username string, db *sql.DB) (*models.User, error) {
	query := `SELECT [Id],[UserName],[Password],[CreatedDate],[IsAdmin] FROM [dbo].[Users] WHERE UserName = @p1`
	user := new(models.User)
	err := db.QueryRow(query, username).Scan(&user.UserId, &user.UserName, &user.Password, &user.CreatedDate, &user.IsAdmin)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

func AddUser(user models.User, db *sql.DB) error {
	query := `INSERT INTO [dbo].[Users] ([UserName],[Password],[IsAdmin]) VALUES (@p1, @p2, @p3)`
	_, err := db.Exec(query, user.UserName, user.Password, user.IsAdmin)
	if err != nil {
		return err
	}

	return nil
}
