// This is an example of implementing the Pet Store from the OpenAPI documentation
// found at:
// https://github.com/OAI/OpenAPI-Specification/blob/master/examples/v3.0/petstore.yaml

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/gorilla/mux"
	"server/api"
)

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	flag.Parse()

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

	// Authentication option
	validatorOptions := &middleware.Options{}
	validatorOptions.Options.AuthenticationFunc = api.CustomAuthenticationFunc
	// end Authentication option

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	// r.Use(middleware.OapiRequestValidator(swagger))
	r.Use(middleware.OapiRequestValidatorWithOptions(swagger, validatorOptions)) // use Authentication option

	// We now register our srvApi above as the handler for the interface
	api.HandlerFromMux(srvApi, r)

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
	}

	fmt.Println("Listening on port", *port)
	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
