package apartment

import (
	"context"
	"strconv"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints represent the apartment service endpoints
type Endpoints struct {
	GetApartment endpoint.Endpoint
	Create       endpoint.Endpoint
}

// UserApartmentsEndpoint represent the apartment service endpoints
type UserApartmentsEndpoint struct {
	GetUserApartments endpoint.Endpoint
}

// MakeGetApartmentEndpoint returns an endpoint used for getting an apartment
func MakeGetApartmentEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetApartmentRequest)
		ID, err := strconv.Atoi(req.ID)

		if err != nil {
			ID = -1
		}

		return s.GetApartment(ctx, ID)
	}
}

// MakeGetUserApartmentsEndpoint returns an endpoint used for getting a all apartments associated to the user
func MakeGetUserApartmentsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserApartmentsRequest)
		userID, err := strconv.Atoi(req.ID)

		if err != nil {
			userID = -1
		}

		return s.GetUserApartments(ctx, userID)
	}
}

// MakeCreateApartmentEndpoint returns an endpoint used for creating an apartment
func MakeCreateApartmentEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateApartmentRequest)

		return s.CreateApartment(ctx, req.Apartment)
	}
}

// GetApartmentRequest represents the request parameters used (ID) for getting one apartment
type GetApartmentRequest struct {
	ID string `json:"id"`
}

// GetUserApartmentsRequest represents the request parameters used (ID) for getting all apartments for a given user
type GetUserApartmentsRequest struct {
	ID string `json:"id"`
}

// CreateApartmentRequest represents the request parameters used for creating apartment
type CreateApartmentRequest struct {
	Apartment
}
