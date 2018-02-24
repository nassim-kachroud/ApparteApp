package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
	"time"

	kitlog "github.com/go-kit/kit/log"
	kitopentracing "github.com/go-kit/kit/tracing/opentracing"
	stdopentracing "github.com/opentracing/opentracing-go"

	"github.com/ricardo-ch/apparte-app/apartment"
	"github.com/ricardo-ch/apparte-app/config"
	"github.com/ricardo-ch/apparte-app/database"
	"github.com/ricardo-ch/apparte-app/user"
)

const (
	appName                     = "apparte-api"
	defaultTracingOperationName = "client transaction start to end"
	httpReadTimeout             = 5 * time.Second
	httpWriteTimeout            = 10 * time.Second
)

var (
	err         error
	errc        chan error
	infoLogger  = kitlog.NewJSONLogger(os.Stdout)
	errorLogger = kitlog.NewJSONLogger(os.Stderr)
	tracer      stdopentracing.Tracer
)

func init() {
	// Initialize configuration
	config.Get()
	tracer = stdopentracing.NoopTracer{}
	infoLogger.Log("tracer", "noop tracer")
}

func main() {

	//Db Connection
	db, err := database.NewConnection(config.DbHost, config.DbPort, config.DbName, config.DbUser, config.DbPassword)
	if err != nil {
		fmt.Println("Cannot Access Database", err)
		os.Exit(-1)
	}
	defer db.Close()

	// Endpoints domain
	var userEndpoints user.Endpoints
	{
		var repo user.Repository
		{
			repo = user.NewUserRepository(db, errorLogger)

		}
		var service user.Service
		{
			service = user.NewService(repo)
		}

		userEndpoints = user.Endpoints{
			Create:  kitopentracing.TraceServer(tracer, defaultTracingOperationName)(user.MakeCreateUserEndpoint(service)),
			GetByID: kitopentracing.TraceServer(tracer, defaultTracingOperationName)(user.MakeGetUserByIDEndpoint(service)),
			// GetByUsernameAndPwd: kitopentracing.TraceServer(tracer, defaultTracingOperationName)(user.MakeGetUserByUsernameAndPwdEndpoint(service)),
			// Update:  kitopentracing.TraceServer(tracer, defaultTracingOperationName)(user.MakeUpdateEndpoint(service)),
		}
	}

	var loginEndpoints user.LoginEndpoints
	{
		var repo user.Repository
		{
			repo = user.NewUserRepository(db, errorLogger)

		}
		var service user.Service
		{
			service = user.NewService(repo)
		}
		loginEndpoints = user.LoginEndpoints{
			GetByUsernameAndPwd: kitopentracing.TraceServer(tracer, defaultTracingOperationName)(user.MakeGetUserByUsernameAndPwdEndpoint(service)),
		}
	}

	var apartmentEndpoints apartment.Endpoints
	{
		var repo apartment.Repository
		{
			repo = apartment.NewApartmentRepository(db, errorLogger)

		}
		var service apartment.Service
		{
			service = apartment.NewService(repo)
		}

		apartmentEndpoints = apartment.Endpoints{
			Create:       kitopentracing.TraceServer(tracer, defaultTracingOperationName)(apartment.MakeCreateApartmentEndpoint(service)),
			GetApartment: kitopentracing.TraceServer(tracer, defaultTracingOperationName)(apartment.MakeGetApartmentEndpoint(service)),
			// Update:  kitopentracing.TraceServer(tracer, defaultTracingOperationName)(user.MakeUpdateEndpoint(service)),
		}
	}

	var userApartmentsEndpoint apartment.UserApartmentsEndpoint
	{
		var repo apartment.Repository
		{
			repo = apartment.NewApartmentRepository(db, errorLogger)

		}
		var service apartment.Service
		{
			service = apartment.NewService(repo)
		}

		userApartmentsEndpoint = apartment.UserApartmentsEndpoint{
			GetUserApartments: kitopentracing.TraceServer(tracer, defaultTracingOperationName)(apartment.MakeGetUserApartmentsEndpoint(service)),
		}
	}
	// Interrupt handler.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// HTTP Transport
	go func() {
		httpAddr := ":" + config.Port
		var mux *http.ServeMux
		{
			mux = http.NewServeMux()
			mux.Handle("/apartments/", apartment.MakeApartmentsHTTPHandler(kitlog.With(errorLogger, "api", "apartment"), tracer, apartmentEndpoints))
			mux.Handle("/users/", user.MakeUsersHTTPHandler(kitlog.With(errorLogger, "api", "user"), tracer, userEndpoints))
			mux.Handle("/login/", user.MakeLoginHTTPHandler(kitlog.With(errorLogger, "api", "login"), tracer, loginEndpoints))
			mux.Handle("/user/", apartment.MakeUserApartmentsHTTPHandler(kitlog.With(errorLogger, "api", "user_apartment"), tracer, userApartmentsEndpoint))

			// mux.HandleFunc("/healthz", Healthz)
		}
		s := &http.Server{
			Addr:         httpAddr,
			ReadTimeout:  httpReadTimeout,
			WriteTimeout: httpWriteTimeout,
			Handler:      mux,
		}
		infoLogger.Log("service", appName, "transport", "http", "address", httpAddr, "msg", "listening")
		errc <- s.ListenAndServe()
	}()
	fmt.Println("exit", <-errc)

	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/healthz", Healthz)

	// log.Fatal(http.ListenAndServe(":8080", router))
}

//Healthz is an endpoint needed for kubernetes
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func getRootPath() string {
	_, dirname, _, _ := runtime.Caller(0)
	return filepath.Dir(dirname)
}
