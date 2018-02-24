package user

import (
	"database/sql"
	"time"
)

// User model
type User struct {
	ID               int        `json:"user_id"`
	Gender           *string    `json:"gender,omitempty"`
	Lastname         *string    `json:"lastname,omitempty"`
	Firstname        *string    `json:"firstname,omitempty"`
	Dateofbirth      *time.Time `json:"date_of_birth,omitempty"`
	Email            *string    `json:"email,omitempty"`
	Username         *string    `json:"username,omitempty"`
	Mobile           *string    `json:"mobile,omitempty"`
	Address          *string    `json:"address,omitempty"`
	Password         *string    `json:"password,omitempty"`
	RegistrationDate *time.Time `json:"registration_date,omitempty"`
}

// ACL

type userDAO struct {
	ID               int
	Gender           sql.NullString
	Username         sql.NullString
	Email            sql.NullString
	Lastname         sql.NullString
	Firstname        sql.NullString
	Dateofbirth      sql.NullString
	Mobile           sql.NullString
	Address          sql.NullString
	Registrationdate sql.NullString
}

type createUserDAO struct {
	Gender      sql.NullString
	Lastname    sql.NullString
	Firstname   sql.NullString
	Dateofbirth sql.NullString
	Email       sql.NullString
	Username    sql.NullString
	Address     sql.NullString
	Mobile      sql.NullString
	Password    sql.NullString
}

// Queries

const (
	getUserByIDQuery = `
		SELECT U.ID as id,
			U.Gender as gender, 
			U.Lastname as lastname,
			U.Firstname as firstname,
			U.Dob as dateofbirth,
			U.Email as email,
			U.Username as username,
			U.Address as address,
			U.Mobile as mobile,
			U.Regdate as registrationdate
		FROM apparte_app.users U
		WHERE U.ID = ?
		LIMIT 1
	`

	insertUserQuery = "INSERT INTO users (ID, Gender, Lastname, Firstname, Dob, Email, Username, Address, Mobile, Password, Regdate) VALUES (0, ?, ?, ?, ?, ?, ?, ?, ?, ?, current_date());"
	// for performance reason dedicated query for login with nickname
	checkCredentialsWithNickQuery = `
		SELECT ID AS id
		FROM apparte_app.users
		WHERE Username = ? AND Password = ?
		LIMIT 1
		`

	// for performance reason dedicated query for login with email (do not remove cast to avoid implicit conversion)
	checkCredentialsWithEmailQuery = `
		SELECT ID AS id
		FROM [apparte_app].users
		WHERE Email = CAST(? AS VARCHAR(60)) AND Password = ?
		LIMIT 1
		`
)
