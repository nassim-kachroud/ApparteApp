package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	httptransport "github.com/go-kit/kit/transport/http"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	stdopentracing "github.com/opentracing/opentracing-go"

	"github.com/pkg/errors"
)

const xForwardedForHeaderKey = "X-Forwarded-For"
const xAdminIDKey = "X-Admin-ID"

// ErrInvalidBody thrown when the body of a request can not be parsed
var ErrInvalidBody = errors.New("invalid body")

// ErrInvalidIPAddress thrown when customer IP is not valid or empty
var ErrInvalidIPAddress = errors.New("invalid IP address")

// ErrInvalidAdminID thrown when the body of a request can not be parsed
var ErrInvalidAdminID = errors.New("invalid admin id")

// MakeUsersHTTPHandler returns all http handler for the user service
func MakeUsersHTTPHandler(logger log.Logger, tracer stdopentracing.Tracer, endpoints Endpoints) http.Handler {
	options := []kithttp.ServerOption{
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
		kithttp.ServerErrorEncoder(encodeError(logger)),
	}

	getUserByIDHandler := kithttp.NewServer(
		endpoints.GetByID,
		decodeGetUserByIDRequest,
		encodeResponse,
		append(options, httptransport.ServerBefore(opentracing.FromHTTPRequest(tracer, "calling HTTP GET /{id}", logger)))...,
	)

	// updateUserHandler := kithttp.NewServer(
	// 	endpoints.Update,
	// 	decodeUpdateUserRequest,
	// 	encodeUpdateUserResponse,
	// 	append(options, httptransport.ServerBefore(opentracing.FromHTTPRequest(tracer, "calling HTTP PATCH /{id}", logger)))...,
	// )

	createUserHandler := kithttp.NewServer(
		endpoints.Create,
		decodeCreateUserRequest,
		encodeCreateUserResponse,
		append(options, httptransport.ServerBefore(opentracing.FromHTTPRequest(tracer, "calling HTTP POST /", logger)))...,
	)

	r := mux.NewRouter().PathPrefix("/users/").Subrouter().StrictSlash(true)

	r.Handle("/{id}", getUserByIDHandler).Methods("GET")
	// r.Handle("/{id}", updateUserHandler).Methods("PATCH")
	r.Handle("/", createUserHandler).Methods("POST")

	return r
}

// MakeUsersHTTPHandler returns all http handler for the user service
func MakeLoginHTTPHandler(logger log.Logger, tracer stdopentracing.Tracer, endpoints LoginEndpoints) http.Handler {
	options := []kithttp.ServerOption{
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
		kithttp.ServerErrorEncoder(encodeError(logger)),
	}

	getUserByUsernameAndPwdHandler := kithttp.NewServer(
		endpoints.GetByUsernameAndPwd,
		decodeGetUserByUsernameAndPwdRequest,
		encodeResponse,
		append(options, httptransport.ServerBefore(opentracing.FromHTTPRequest(tracer, "calling HTTP POST /login", logger)))...,
	)

	r := mux.NewRouter().PathPrefix("/login/").Subrouter().StrictSlash(true)
	r.Handle("/", getUserByUsernameAndPwdHandler).Methods("POST")

	return r
}

func decodeGetUserByIDRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)

	return GetUserByIDRequest{ID: vars["id"]}, nil
}

func decodeCreateUserRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrInvalidBody
	}

	return req, nil
}

func decodeGetUserByUsernameAndPwdRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req GetUserByUsernameAndPwdRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrInvalidBody
	}

	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeCreateUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	// TODO : refactor return
	ID, ok := response.(*int)
	if !ok {
		return errors.New("An error occured while creating user")
	}
	w.Header().Set("Location", fmt.Sprintf("/users/%v", *ID))
	w.WriteHeader(http.StatusCreated)
	return nil
}

// clientError is the object to return to the client
type clientError struct {
	HTTPCode int                    `json:"-"`
	Message  string                 `json:"message,omitempty"`
	Errors   *map[string]fieldError `json:"errors,omitempty"`
}

// fieldError is a detailed exception for a field
type fieldError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
