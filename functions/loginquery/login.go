package loginquery

import (
	"context"
	"database/sql"
	"fmt"
	"pos-backend/models"
	"time"
)

func QueryReadAllUser(db *sql.DB, ctx context.Context) (error, []models.UserAccount) {
	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return err, []models.UserAccount{}
	}

	tsql := fmt.Sprintf("SELECT TOP 200 ID, 'Aktif' Status, InputDate, Nama, Keterangan, Email, Pass, UserNumber From UserAccount WHERE ISNULL(UserNumber,'') <> '' AND Status = 1 ORDER BY ID DESC;")

	// Execute query
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		return err, []models.UserAccount{}
	}

	defer rows.Close()

	var count int

	// Iterate through the result set.
	users := []models.UserAccount{}
	for rows.Next() {
		var user models.UserAccount
		var idget int
		var statusget string
		var inputdateget time.Time
		var namaget string
		var keteranganget string
		var emailget string
		var passget string
		var usernumberget string
		err := rows.Scan(&idget, &statusget, &inputdateget, &namaget, &keteranganget, &emailget, &passget, &usernumberget)
		if err != nil {
			return err, []models.UserAccount{}
		}
		user.ID = idget
		user.Status = statusget
		user.InputDate = inputdateget
		user.Nama = namaget
		user.Keterangan = keteranganget
		user.Email = emailget
		user.Pass = passget
		user.UserNumber = usernumberget
		users = append(users, user)
		count++
	}

	db.Close()

	return nil, users
}

func QueryUserLogin(username string, password string, db *sql.DB, ctx context.Context) (string, models.UserAccount) {
	var msg string = ""
	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		msg = err.Error()
		return msg, models.UserAccount{}
	}
	tsql := fmt.Sprintf("SELECT COUNT('') FROM UserAccount WHERE Status = 1 AND UserNumber = @UserNumber AND Pass = @Password")
	tsql1 := fmt.Sprintf("SELECT TOP 1 ID, 'Aktif' Status, InputDate, Nama, Keterangan, Email, Pass, UserNumber From UserAccount WHERE Status = 1 AND UserNumber = @UserNumber AND Pass = @Password")

	// Execute query
	rows := db.QueryRowContext(ctx,
		tsql,
		sql.Named("UserNumber", username),
		sql.Named("Password", password))
	var count int
	err = rows.Scan(&count)
	if err != nil {
		msg = err.Error()
		return msg, models.UserAccount{}
	}

	if count == 0 {
		msg = "Username Not Found"
		return msg, models.UserAccount{}
	}

	//Execute query
	rows1 := db.QueryRowContext(ctx,
		tsql1,
		sql.Named("UserNumber", username),
		sql.Named("Password", password))
	user := models.UserAccount{}
	var idget int
	var statusget string
	var inputdateget time.Time
	var namaget string
	var keteranganget string
	var emailget string
	var passget string
	var usernumberget string

	err = rows1.Scan(&idget, &statusget, &inputdateget, &namaget, &keteranganget, &emailget, &passget, &usernumberget)
	if err != nil {
		msg = err.Error()
		return msg, models.UserAccount{}
	}

	user.ID = idget
	user.Status = statusget
	user.InputDate = inputdateget
	user.Nama = namaget
	user.Keterangan = keteranganget
	user.Email = emailget
	user.Pass = passget
	user.UserNumber = usernumberget

	db.Close()

	return msg, user
}
