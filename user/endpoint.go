package user

import (
	"context"
	"strconv"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints represent the order service endpoints
type Endpoints struct {
	GetByID endpoint.Endpoint
	Create  endpoint.Endpoint
}

// LoginEndpoints represent the order service endpoints
type LoginEndpoints struct {
	GetByUsernameAndPwd endpoint.Endpoint
}

// MakeGetUserEndpoint returns an endpoint used for getting an user
func MakeGetUserByUsernameAndPwdEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserByUsernameAndPwdRequest)

		return s.GetUserByUsernameAndPwd(ctx, req.Username, req.Password)
	}
}

// MakeGetUserByIDEndpoint returns an endpoint used for getting one user
func MakeGetUserByIDEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserByIDRequest)
		ID, err := strconv.Atoi(req.ID)

		if err != nil {
			ID = -1
		}

		return s.GetUser(ctx, ID)
	}
}

// MakeCreateUserEndpoint returns an endpoint used for creating an user
func MakeCreateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)

		return s.CreateUser(ctx, req.User)
	}
}

// GetUserByUsernameAndPwdRequest represents the request parameters used (username and pwd) for getting one user
type GetUserByUsernameAndPwdRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// GetUserByIDRequest represents the request parameters used (ID) for getting one user
type GetUserByIDRequest struct {
	ID string `json:"id"`
}

// CreateUserRequest represents the request parameters used for creating user
type CreateUserRequest struct {
	User
}
