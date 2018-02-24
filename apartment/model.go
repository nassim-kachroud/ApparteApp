package apartment

import (
	"database/sql"

	"github.com/ricardo-ch/apparte-app/user"
)

// Apartment model
type Apartment struct {
	ID          int        `json:"apartment_id"`
	User        *user.User `json:"user,omitempty"`
	Type        *string    `json:"type,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Address     *string    `json:"address,omitempty"`
	Postalcode  *string    `json:"postal_code"`
	City        *string    `json:"city,omitempty"`
	Country     *string    `json:"country,omitempty"`
	Area        *string    `json:"area,omitempty"`
	Roomsnb     *string    `json:"roomsnb,omitempty"`
	Description *string    `json:"description,omitempty"`
}

// ACL

type apartmentDAO struct {
	ID          int
	UserID      sql.NullInt64
	Type        sql.NullString
	Name        sql.NullString
	Address     sql.NullString
	Postalcode  sql.NullString
	City        sql.NullString
	Country     sql.NullString
	Area        sql.NullString
	Roomsnb     sql.NullString
	Description sql.NullString
}

type createApartmentDAO struct {
	UserID      sql.NullInt64
	Type        sql.NullString
	Name        sql.NullString
	Address     sql.NullString
	Postalcode  sql.NullString
	City        sql.NullString
	Country     sql.NullString
	Area        sql.NullString
	Roomsnb     sql.NullString
	Description sql.NullString
}

// Queries

const (
	getApartmentQuery = `
		SELECT A.ID as id,
			A.UserID as userid, 
			A.Type as type,
			A.Name as name,
			A.Address as address,
			A.Postalcode as postalcode,
			A.City as city,
			A.Country as country,
			A.Area as area,
			A.Roomsnb as roomsnb,
			A.Description as description
		FROM apparte_app.apartments A
		WHERE A.ID = ?
	`

	getUserApartmentsQuery = `
		SELECT A.ID as id,
			A.UserID as userid, 
			A.Type as type,
			A.Name as name,
			A.Address as address,
			A.Postalcode as postalcode,
			A.City as city,
			A.Country as country,
			A.Area as area,
			A.Roomsnb as roomsnb,
			A.Description as description
		FROM apparte_app.apartments A
		INNER JOIN apparte_app.users B ON B.ID = A.UserID
		WHERE B.ID = ?
	`

	insertApartmentQuery = "INSERT INTO apartments (ID, UserID, Type, Name, Address, Postalcode, City, Country, Area, Roomsnb, Description) VALUES (0, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
)
