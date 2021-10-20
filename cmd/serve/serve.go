package serve

import (
	"fmt"
	"net/http"

	"github.com/StiviiK/go-modules-http-proxy/cmd"
	"github.com/StiviiK/go-modules-http-proxy/pkg/html"
	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"
)

var (
	address string
	port    int
)

func init() {
	command := &cli.Command{
		Name: "serve",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "address",
				Value:       "0.0.0.0",
				Usage:       "Address to listen on",
				Destination: &address,
			},
			&cli.IntFlag{
				Name:        "port",
				Aliases:     []string{"p"},
				Value:       8080,
				Usage:       "Port to listen on",
				Destination: &port,
			},
		},
		Action: func(c *cli.Context) error {
			httpRouter := mux.NewRouter()
			httpRouter.HandleFunc("/", html.Doge())
			httpRouter.PathPrefix("/{[a-zA-Z0-9=-/]+}").HandlerFunc(html.Git("https://github.com/uber-go"))

			return http.ListenAndServe(fmt.Sprintf("%s:%d", address, port), httpRouter)
		},
	}

	cmd.Register(command)
}