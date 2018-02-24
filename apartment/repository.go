package apartment

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

	kitlog "github.com/go-kit/kit/log"
	"github.com/ricardo-ch/apparte-app/user"

	"github.com/jmoiron/sqlx"
)

// ErrNotFound is used when an user is not found
var ErrNotFound = errors.New("apartment not found")

// Repository represents an user repository interface
type Repository interface {
	GetApartment(ctx context.Context, id int) (*Apartment, error)
	GetUserApartments(ctx context.Context, id int) (*[]Apartment, error)
	// UpdateUser(ctx context.Context, u User, adminID string, hashedPassword string) error
	CreateApartment(ctx context.Context, a Apartment) (*int, error)
}

type apartmentRepository struct {
	db     *sqlx.DB
	logger kitlog.Logger
}

// NewApartmentRepository creates a new instance of an apartment repository
func NewApartmentRepository(db *sqlx.DB, logger kitlog.Logger) Repository {
	return apartmentRepository{
		db:     db,
		logger: logger,
	}
}

// GetApartment ...
func (r apartmentRepository) GetApartment(ctx context.Context, id int) (a *Apartment, err error) {
	var dao apartmentDAO

	err = r.db.Get(&dao, getApartmentQuery, id)

	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	a = mapApartmentDAO(&dao)

	return
}

// GetUserApartments ...
func (r apartmentRepository) GetUserApartments(ctx context.Context, id int) (a *[]Apartment, err error) {
	result := []apartmentDAO{}

	err = r.db.Select(&result, getUserApartmentsQuery, id)

	if err != nil {
		return nil, err
	}
	return mapApartmentDAOArray(&result), nil
}

// CreateApartment ...
func (r apartmentRepository) CreateApartment(ctx context.Context, a Apartment) (*int, error) {
	// Map apartment model to an create user model
	dao := mapCreateApartmentDAO(a)
	var apartmentID int
	res, err := r.db.Exec(insertApartmentQuery,
		dao.UserID,
		dao.Type,
		dao.Name,
		dao.Address,
		dao.Postalcode,
		dao.City,
		dao.Country,
		dao.Area,
		dao.Roomsnb,
		dao.Description,
	)
	if err != nil {
		println("Error apartment:", err.Error())
		return nil, err
	}
	ID, err := res.LastInsertId()
	if err != nil {
		println("Error last insert id apartment:", err.Error())
		return nil, err
	}

	apartmentID = int(ID)
	return &apartmentID, nil
}

func mapApartmentDAO(dao *apartmentDAO) (a *Apartment) {
	a = new(Apartment)
	a.ID = dao.ID
	a.User = &user.User{ID: int(dao.UserID.Int64)}
	a.Type = stringToPtr(strings.Trim(dao.Type.String, " "))
	a.Name = stringToPtr(strings.Trim(dao.Name.String, " "))
	a.Address = stringToPtr(strings.Trim(dao.Address.String, " "))
	a.Postalcode = stringToPtr(strings.Trim(dao.Postalcode.String, " "))
	a.City = stringToPtr(strings.Trim(dao.City.String, " "))
	a.Country = stringToPtr(strings.Trim(dao.Country.String, " "))
	a.Area = stringToPtr(strings.Trim(dao.Area.String, " "))
	a.Roomsnb = stringToPtr(strings.Trim(dao.Roomsnb.String, " "))
	a.Description = stringToPtr(strings.Trim(dao.Description.String, " "))

	return a
}

func mapApartmentDAOArray(DAOSlice *[]apartmentDAO) *[]Apartment {
	result := []Apartment{}
	for _, DAOItem := range *DAOSlice {
		result = append(result, *mapApartmentDAO(&DAOItem))
	}
	return &result
}

func mapCreateApartmentDAO(a Apartment) (dao *createApartmentDAO) {
	dao = new(createApartmentDAO)
	dao.UserID = intToNullInt64(a.User.ID)
	dao.Type = toNullString(strings.ToLower(*a.Type))
	dao.Name = toNullString(*a.Name)
	dao.Address = toNullString(*a.Address)
	dao.Postalcode = toNullString(*a.Postalcode)
	dao.City = toNullString(*a.City)
	dao.Country = toNullString(*a.Country)
	dao.Area = toNullString(*a.Area)
	dao.Roomsnb = toNullString(*a.Roomsnb)
	dao.Description = toNullString(*a.Description)
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
