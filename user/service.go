package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"

	validator "gopkg.in/go-playground/validator.v9"
)

// ErrInconsistentID ...
var ErrInconsistentID = errors.New("inconsistent userid")

// ErrNicknameNotUnique ...
var ErrNicknameNotUnique = errors.New("nickname not unique")

// GetUser returns an user regarding the id passed in parameter
func (s service) GetUser(ctx context.Context, id int) (u *User, err error) {
	u, err = s.repository.GetUser(ctx, id)

	if err != nil {
		return nil, err
	}

	if u == nil {

		return nil, ErrNotFound
	}

	return u, err
}

// Service is the Order service interface
type Service interface {
	GetUser(ctx context.Context, id int) (*User, error)
	// UpdateUser(ctx context.Context, updateUserRequest UpdateUserRequest) error
	CreateUser(ctx context.Context, user User) (*int, error)
	GetUserByUsernameAndPwd(ctx context.Context, username, password string) (*User, error)
}

// NewService return a new instance of order service
func NewService(repository Repository) Service {
	return service{
		repository: repository,
	}
}

type service struct {
	repository Repository
}

// CreateUser creates an user
func (s service) CreateUser(ctx context.Context, u User) (*int, error) {
	// Create user
	userID, err := s.repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	return userID, nil
}

// GetUserByUsernameAndPwd ...
func (s service) GetUserByUsernameAndPwd(ctx context.Context, username, password string) (u *User, err error) {

	fmt.Println("GetUserBysernameAndPwd: ", username, password)
	userID, err := s.checkCredentials(ctx, username, password)
	if err != nil {
		return nil, err
	}
	u, err = s.GetUser(ctx, userID)

	if err != nil {
		return nil, err
	}

	return u, err
}

// checkCredentials verifies if credentials (username or e-mail + password) are correct.
// If yes it sends back user basics info, otherwise there will be empty object
func (s service) checkCredentials(ctx context.Context, username string, password string) (userID int, err error) {
	if username == "" || password == "" {
		return -1, ErrInvalidCredentials
	}

	// Check credentials using SHA1 hash
	userID, err = s.repository.CheckUserCredentials(ctx, username, password)

	return userID, err
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
		case ErrInvalidCredentials:
			w.WriteHeader(http.StatusUnauthorized)
		case ErrInconsistentID,
			ErrInvalidBody, ErrInvalidIPAddress, ErrNicknameNotUnique, ErrInvalidAdminID:
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
		fErrorName = strings.Replace(fErrorName, "updateuserrequest.user.", "user.", -1)
		fErrorName = strings.Replace(fErrorName, "updateuserrequest.", "user.", -1)
		if _, ok := m[fErrorName]; !ok {
			m[fErrorName] = fieldError{Code: fe.Tag(), Message: fe.Tag()}
		}
	}
	return multipleError(http.StatusUnprocessableEntity, "Validation form failed", &m)
}
