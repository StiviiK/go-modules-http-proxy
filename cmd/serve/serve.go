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
	host       string
	port       int
	httpRouter *mux.Router
)

func init() {
	httpRouter = mux.NewRouter()
	httpRouter.PathPrefix("/").HandlerFunc(html.GitHubHandler)

	command := &cli.Command{
		Name: "serve",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "address",
				Aliases:     []string{"host"},
				Value:       "0.0.0.0",
				Usage:       "Host to listen on",
				Destination: &host,
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
			fmt.Printf("Listening on %s:%d\n", host, port)
			return http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), httpRouter)
		},
	}

	cmd.Register(command)
}
