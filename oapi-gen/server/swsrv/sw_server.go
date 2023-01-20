package swsrv

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/gorilla/mux"

	"server/api"
	"server/utils"
)

// r *mux.Router

func StartServer(port int) {

	utils.MaxOpenFiles()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Create an instance of our handler which satisfies the generated interface
	srvApi := api.NewAPI()

	// This is how you set up a basic Gorilla router
	r := mux.NewRouter()

	r.Use(utils.LoggingMiddleware)

	validatorOptions := &middleware.Options{}
	// Authentication option
	validatorOptions.Options.AuthenticationFunc = api.CustomAuthenticationFunc
	// Error handler
	validatorOptions.ErrorHandler = api.CustomErrorHandler

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	// r.Use(middleware.OapiRequestValidator(swagger))
	r.Use(middleware.OapiRequestValidatorWithOptions(swagger, validatorOptions)) // use Authentication option

	// We now register our srvApi above as the handler for the interface
	api.HandlerFromMux(srvApi, r)

	s := &http.Server{
		Handler: utils.CorsObj.Handler(r),
		Addr:    fmt.Sprintf(":%d", port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	fmt.Println("Listening on port", port)
	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
