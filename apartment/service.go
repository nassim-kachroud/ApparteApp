package apartment

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"

	validator "gopkg.in/go-playground/validator.v9"
)

// ErrInconsistentID ...
var ErrInconsistentID = errors.New("inconsistent apartmentID")

// Service is the Order service interface
type Service interface {
	GetApartment(ctx context.Context, id int) (*Apartment, error)
	GetUserApartments(ctx context.Context, id int) (*[]Apartment, error)
	// UpdateUser(ctx context.Context, updateUserRequest UpdateUserRequest) error
	CreateApartment(ctx context.Context, apartment Apartment) (*int, error)
}

// NewService return a new instance of apartment service
func NewService(repository Repository) Service {
	return service{
		repository: repository,
	}
}

type service struct {
	repository Repository
}

// GetApartment returns an apartment regarding the id passed in parameter
func (s service) GetApartment(ctx context.Context, id int) (a *Apartment, err error) {
	a, err = s.repository.GetApartment(ctx, id)

	if err != nil {
		return nil, err
	}
	if a == nil {
		return nil, ErrNotFound
	}
	return a, err
}

// CreateApartment creates an user
func (s service) CreateApartment(ctx context.Context, a Apartment) (*int, error) {
	// Create apartment
	apartmentID, err := s.repository.CreateApartment(ctx, a)
	if err != nil {
		return nil, err
	}

	return apartmentID, nil
}

// GetUserApartments ...
func (s service) GetUserApartments(ctx context.Context, id int) (a *[]Apartment, err error) {
	a, err = s.repository.GetUserApartments(ctx, id)

	if err != nil {
		return nil, err
	}
	if a == nil || len(*a) == 0 {
		return nil, ErrNotFound
	}
	return a, err
}

func encodeError(logger log.Logger) kithttp.ErrorEncoder {
	return func(ctx context.Context, err error, w http.ResponseWriter) {
		// Validation errors
		ve, ok := err.(validator.ValidationErrors)
		if ok {
			clientErr := validationError(ve)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(clientErr.HTTPCode)
			json.NewEncoder(w).Encode(clientErr)
			return
		}
		// Other errors
		switch err {
		default:
			logger.Log("err", err,
				"http.url", ctx.Value(kithttp.ContextKeyRequestURI), "http.path", ctx.Value(kithttp.ContextKeyRequestPath),
				"http.method", ctx.Value(kithttp.ContextKeyRequestMethod), "http.user_agent", ctx.Value(kithttp.ContextKeyRequestUserAgent),
				"http.proto", ctx.Value(kithttp.ContextKeyRequestProto))
			w.WriteHeader(http.StatusInternalServerError)
		case ErrInconsistentID,
			ErrInvalidBody, ErrInvalidParams:
			w.WriteHeader(http.StatusBadRequest)
		case ErrNotFound:
			w.WriteHeader(http.StatusNotFound)
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
	}
}

// simpleError returns a simple error without field error
func simpleError(httpCode int, message string) *clientError {
	return &clientError{HTTPCode: httpCode, Message: message}
}

// multipleError returns an error with multiple fields errors
func multipleError(httpCode int, message string, errors *map[string]fieldError) *clientError {
	cEx := simpleError(httpCode, message)
	cEx.Errors = errors
	return cEx
}

func validationError(ve validator.ValidationErrors) *clientError {
	m := make(map[string]fieldError)
	for _, fe := range ve {
		//Return first error on a field and skip others
		fErrorName := strings.ToLower(fe.Namespace())
		//Hack clean name
		fErrorName = strings.Replace(fErrorName, "updateapartmentrequest.apartment.", "apartment.", -1)
		fErrorName = strings.Replace(fErrorName, "updateapartmentrequest.", "apartment.", -1)
		if _, ok := m[fErrorName]; !ok {
			m[fErrorName] = fieldError{Code: fe.Tag(), Message: fe.Tag()}
		}
	}
	return multipleError(http.StatusUnprocessableEntity, "Validation form failed", &m)
}
