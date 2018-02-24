package apartment

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

// ErrInvalidParams thrown when the params of a request can not be parsed
var ErrInvalidParams = errors.New("invalid params")

// MakeApartmentsHTTPHandler returns all http handler for the apartment service
func MakeApartmentsHTTPHandler(logger log.Logger, tracer stdopentracing.Tracer, endpoints Endpoints) http.Handler {
	options := []kithttp.ServerOption{
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
		kithttp.ServerErrorEncoder(encodeError(logger)),
	}

	getApartmentHandler := kithttp.NewServer(
		endpoints.GetApartment,
		decodeGetApartmentRequest,
		encodeResponse,
		append(options, httptransport.ServerBefore(opentracing.FromHTTPRequest(tracer, "calling HTTP GET /{id}", logger)))...,
	)

	// updateUserHandler := kithttp.NewServer(
	// 	endpoints.Update,
	// 	decodeUpdateUserRequest,
	// 	encodeUpdateUserResponse,
	// 	append(options, httptransport.ServerBefore(opentracing.FromHTTPRequest(tracer, "calling HTTP PATCH /{id}", logger)))...,
	// )

	createApartmentHandler := kithttp.NewServer(
		endpoints.Create,
		decodeCreateApartmentRequest,
		encodeCreateApartmentResponse,
		append(options, httptransport.ServerBefore(opentracing.FromHTTPRequest(tracer, "calling HTTP POST /", logger)))...,
	)

	r := mux.NewRouter().PathPrefix("/apartments/").Subrouter().StrictSlash(true)

	r.Handle("/{id}", getApartmentHandler).Methods("GET")
	// r.Handle("/{id}", updateUserHandler).Methods("PATCH")
	r.Handle("/", createApartmentHandler).Methods("POST")

	return r
}

// MakeUserApartmentsHTTPHandler
func MakeUserApartmentsHTTPHandler(logger log.Logger, tracer stdopentracing.Tracer, endpoint UserApartmentsEndpoint) http.Handler {
	options := []kithttp.ServerOption{
		kithttp.ServerBefore(kithttp.PopulateRequestContext),
		kithttp.ServerErrorEncoder(encodeError(logger)),
	}

	getUserApartmentsHandler := kithttp.NewServer(
		endpoint.GetUserApartments,
		decodeGetUserApartmentsRequest,
		encodeResponse,
		append(options, httptransport.ServerBefore(opentracing.FromHTTPRequest(tracer, "calling HTTP GET /user/:id/apartments", logger)))...,
	)

	r := mux.NewRouter().PathPrefix("/user/").Subrouter().StrictSlash(true)
	r.Handle("/{id}/apartments", getUserApartmentsHandler).Methods("GET")

	return r
}

func decodeGetApartmentRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)

	return GetApartmentRequest{ID: vars["id"]}, nil
}

func decodeCreateApartmentRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	var req CreateApartmentRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrInvalidBody
	}

	return req, nil
}

func decodeGetUserApartmentsRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	_, err = strconv.Atoi(vars["id"])
	if err != nil {
		return nil, ErrInvalidParams
	}

	return GetUserApartmentsRequest{ID: vars["id"]}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeCreateApartmentResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	ID, ok := response.(*int)
	if !ok {
		return errors.New("An error occured while creating apartment")
	}
	w.Header().Set("Location", fmt.Sprintf("/apartments/%v", *ID))
	w.WriteHeader(http.StatusCreated)
	return nil
}

func getApartmentIDs(ids string) ([]int, error) {
	var apartmentIDs []int
	for _, id := range strings.Split(ids, ",") {
		i, err := strconv.Atoi(id)
		if err != nil {
			return nil, err
		}
		apartmentIDs = append(apartmentIDs, i)
	}
	return apartmentIDs, nil
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
