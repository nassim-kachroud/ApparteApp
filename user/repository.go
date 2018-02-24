package user

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/mail"
	"strconv"
	"strings"
	"time"

	kitlog "github.com/go-kit/kit/log"

	"github.com/jmoiron/sqlx"
)

// ErrNotFound is used when an user is not found
var ErrNotFound = errors.New("user not found")

// ErrInvalidCredentials ...
var ErrInvalidCredentials = errors.New("credentials are invalid")

// Repository represents an user repository interface
type Repository interface {
	GetUser(ctx context.Context, id int) (*User, error)
	// UpdateUser(ctx context.Context, u User, adminID string, hashedPassword string) error
	CreateUser(ctx context.Context, u User) (*int, error)
	CheckUserCredentials(ctx context.Context, username, password string) (int, error)
}

type userRepository struct {
	db     *sqlx.DB
	logger kitlog.Logger
}

// NewUserRepository creates a new instance of a legacy user repository
func NewUserRepository(db *sqlx.DB, logger kitlog.Logger) Repository {
	return userRepository{
		db:     db,
		logger: logger,
	}
}

// GetUser ...
func (r userRepository) GetUser(ctx context.Context, id int) (u *User, err error) {
	var dao userDAO

	err = r.db.Get(&dao, getUserByIDQuery, id)

	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	u = mapUserDAO(&dao)

	return
}

// CreateUser ...
func (r userRepository) CreateUser(ctx context.Context, u User) (*int, error) {
	// Map user model to an create user model
	dao := mapCreateUserDAO(u)
	// TODO : replace SP exec by SQL queries
	var userID int
	res, err := r.db.Exec(insertUserQuery,
		dao.Gender,
		dao.Lastname,
		dao.Firstname,
		dao.Dateofbirth,
		dao.Email,
		dao.Username,
		dao.Address,
		dao.Mobile,
		dao.Password,
	)
	if err != nil {
		println("Error user:", err.Error())
		return nil, err
	}
	ID, err := res.LastInsertId()
	if err != nil {
		println("Error last insert id user:", err.Error())
		return nil, err
	}

	userID = int(ID)
	return &userID, nil
}

// CheckUserCredentials ...
func (r userRepository) CheckUserCredentials(ctx context.Context, username, password string) (int, error) {
	var query string
	if _, err := mail.ParseAddress(username); err != nil {
		// use nickname field for query
		query = checkCredentialsWithNickQuery
	} else {
		// use email field for query
		query = checkCredentialsWithEmailQuery
	}

	var userid *int
	err := r.db.Get(&userid, query, username, password)
	if err == sql.ErrNoRows {
		return -1, ErrInvalidCredentials
	}
	if err != nil {
		return -1, err
	}

	return *userid, nil
}

func mapUserDAO(dao *userDAO) (u *User) {
	u = new(User)
	u.ID = dao.ID
	u.Firstname = stringToPtr(strings.Trim(dao.Firstname.String, " "))
	u.Lastname = stringToPtr(strings.Trim(dao.Lastname.String, " "))
	u.Gender = stringToPtr(strings.Trim(dao.Gender.String, " "))
	u.Username = stringToPtr(strings.Trim(dao.Username.String, " "))
	layout := "2006-01-02"
	birthdate, err := time.Parse(layout, strings.Trim(dao.Dateofbirth.String, `"`))
	if err != nil {
		log.Fatal("Birth date parsing error: ", err)
	}
	u.Dateofbirth = timeToPtr(birthdate)
	regDate, err := time.Parse(layout, strings.Trim(dao.Registrationdate.String, `"`))
	if err != nil {
		log.Fatal("Registration date parsing error: ", err)
	}
	u.RegistrationDate = timeToPtr(regDate)
	u.Email = stringToPtr(strings.Trim(dao.Email.String, " "))
	u.Address = stringToPtr(strings.Trim(dao.Address.String, " "))
	u.Mobile = stringToPtr(strings.Trim(dao.Mobile.String, " "))

	return u
}

func mapCreateUserDAO(u User) (dao *createUserDAO) {
	dao = new(createUserDAO)
	dao.Gender = toNullString(strings.ToLower(*u.Gender))
	dao.Lastname = toNullString(*u.Lastname)
	dao.Firstname = toNullString(*u.Firstname)
	dao.Dateofbirth = toNullString((*u.Dateofbirth).Format("2006-01-02"))
	dao.Email = toNullString(*u.Email)
	dao.Username = toNullString(*u.Username)
	dao.Address = toNullString(*u.Address)
	dao.Mobile = toNullString(*u.Mobile)
	dao.Password = toNullString(*u.Password)
	return
}

func stringToPtr(s string) *string {
	return &s
}

func timeToPtr(t time.Time) *time.Time {
	return &t
}

func toNullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

func toNullInt64(s string) sql.NullInt64 {
	i, err := strconv.Atoi(s)
	return sql.NullInt64{Int64: int64(i), Valid: err == nil}
}

func intToNullInt64(i int) sql.NullInt64 {
	return sql.NullInt64{Int64: int64(i), Valid: true}
}
