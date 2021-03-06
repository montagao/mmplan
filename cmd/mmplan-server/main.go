// Code generated by go-swagger; DO NOT EDIT.

package main

import (
	"log"
	"os"
	"time"

	loads "github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	flags "github.com/jessevdk/go-flags"
	"github.com/montagao/monplan/internal/store"
	"github.com/montagao/monplan/models"
	"github.com/montagao/monplan/restapi"
	"github.com/montagao/monplan/restapi/operations"
)

const (
	listeningPort = 8080
)

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}
	planStore, err := store.New()
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewMmplanAPI(swaggerSpec)
	// ReDoc is default
	// api.UseSwaggerUI()

	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "MM Plan API"
	parser.LongDescription = "For use with the Monta Monta Plan application"

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	api.AddPlanHandler = operations.AddPlanHandlerFunc(
		func(params operations.AddPlanParams) middleware.Responder {
			newPlan := &models.Plan{
				ID:         params.Body.ID,
				IsComplete: params.Body.IsComplete,
				List1:      params.Body.List1,
				List2:      params.Body.List2,
				Name1:      params.Body.Name1,
				Name2:      params.Body.Name2,
				Timestamp:  time.Now().String(),
				PlanName:   params.Body.PlanName,
			}
			err := planStore.Put(newPlan)
			if err != nil {
				log.Printf("%v", err)
			}
			return operations.NewAddPlanCreated().WithPayload(newPlan)
		})

	api.GetPlansHandler = operations.GetPlansHandlerFunc(
		func(params operations.GetPlansParams) middleware.Responder {
			plans, err := planStore.GetAll(int(*params.Limit))
			if err != nil {
				log.Printf("%v", err)
			}
			return operations.NewGetPlansOK().WithPayload(plans)
		})

	api.GetPlanByIDHandler = operations.GetPlanByIDHandlerFunc(
		func(params operations.GetPlanByIDParams) middleware.Responder {
			plan, err := planStore.GetByID(params.ID)
			if err != nil {
				log.Printf("%v", err)
			}

			if plan != nil {
				return operations.NewGetPlanByIDOK().WithPayload(plan)
			} else {
				return operations.NewGetPlanByIDDefault(404)
			}
		})

	api.DeletePlanHandler = operations.DeletePlanHandlerFunc(
		func(params operations.DeletePlanParams) middleware.Responder {
			err := planStore.Delete(params.ID)
			if err != nil {
				log.Printf("%v", err)
			}

			if err != nil {
				return operations.NewGetPlanByIDOK().WithPayload(nil)
			} else {
				return operations.NewGetPlanByIDDefault(404)
			}
		})

	server.ConfigureAPI()
	server.Port = listeningPort

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
