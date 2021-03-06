// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/rs/cors"

	"github.com/montagao/monplan/restapi/operations"
)

//go:generate swagger generate server --target ../../monplan --name Mmplan --spec ../api.yml

func configureFlags(api *operations.MmplanAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MmplanAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.AddPlanHandler == nil {
		api.AddPlanHandler = operations.AddPlanHandlerFunc(func(params operations.AddPlanParams) middleware.Responder {
			return middleware.NotImplemented("operation .AddPlan has not yet been implemented")
		})
	}
	if api.DeletePlanHandler == nil {
		api.DeletePlanHandler = operations.DeletePlanHandlerFunc(func(params operations.DeletePlanParams) middleware.Responder {
			return middleware.NotImplemented("operation .DeletePlan has not yet been implemented")
		})
	}
	if api.GetPlanByIDHandler == nil {
		api.GetPlanByIDHandler = operations.GetPlanByIDHandlerFunc(func(params operations.GetPlanByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation .GetPlanByID has not yet been implemented")
		})
	}
	if api.GetPlansHandler == nil {
		api.GetPlansHandler = operations.GetPlansHandlerFunc(func(params operations.GetPlansParams) middleware.Responder {
			return middleware.NotImplemented("operation .GetPlans has not yet been implemented")
		})
	}
	if api.UpdatePlanHandler == nil {
		api.UpdatePlanHandler = operations.UpdatePlanHandlerFunc(func(params operations.UpdatePlanParams) middleware.Responder {
			return middleware.NotImplemented("operation .UpdatePlan has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	corsHandler := cors.New(cors.Options{
		Debug:          false,
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{},
		MaxAge:         1000,
	})
	return corsHandler.Handler(setupRedocMiddleware(handler))

}

func setupRedocMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/swagger.json" {
			swaggerSpec, err := loads.Analyzed(SwaggerJSON, "")
			if err != nil {
				panic("failed to get swagger.json")
			}
			rawSpec := swaggerSpec.Raw()
			rootHandler := middleware.Spec("/api/", rawSpec, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				rw.WriteHeader(http.StatusFound)
				return
			}))
			rootHandler.ServeHTTP(w, r)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
