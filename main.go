package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
)

type GreetingInput struct {
	Name string `query:"name" doc:"Name to greet"`
}

type GreetingOutput struct {
	Body struct {
		Message string `json:"message"`
	}
}

type HealthOutput struct {
	Body struct {
		Status string `json:"status"`
	}
}

func newAPI() (huma.API, *http.ServeMux) {
	mux := http.NewServeMux()
	config := huma.DefaultConfig("Example API", "1.0.0")
	api := humago.New(mux, config)

	huma.Register(api, huma.Operation{
		OperationID: "get-greeting",
		Method:      http.MethodGet,
		Path:        "/greeting",
		Summary:     "Get a greeting",
	}, func(_ context.Context, input *GreetingInput) (*GreetingOutput, error) {
		out := &GreetingOutput{}
		out.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
		return out, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "get-health",
		Method:      http.MethodGet,
		Path:        "/health",
		Summary:     "Check API health",
	}, func(_ context.Context, _ *struct{}) (*HealthOutput, error) {
		out := &HealthOutput{}
		out.Body.Status = "ok"
		return out, nil
	})

	return api, mux
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "openapi" {
		api, _ := newAPI()
		spec, err := api.OpenAPI().YAML()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(string(spec))
		return
	}

	_, mux := newAPI()
	addr := ":8888"
	fmt.Printf("Listening on %s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
