package serve

import (
	"fmt"
	config "github.com/StiviiK/go-modules-http-proxy/cmd/serve/config"
	"net/http"
	"os"

	"github.com/StiviiK/go-modules-http-proxy/cmd"
	"github.com/StiviiK/go-modules-http-proxy/pkg/html"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"
)

var (
	configPath string
	address    string
	port       int
)

func init() {
	command := &cli.Command{
		Name: "serve",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Value:       "config.yaml",
				Usage:       "config file",
				Destination: &configPath,
			},
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
			&cli.StringFlag{
				Name:    "ssl-cert",
				Aliases: []string{"cert"},
				Usage:   "Path to SSL certificate",
			},
			&cli.StringFlag{
				Name:    "ssl-key",
				Aliases: []string{"key"},
				Usage:   "Path to SSL key",
			},
		},
		Action: func(c *cli.Context) error {
			// Load conf
			conf := &config.Config{}
			if err := conf.Load(configPath); err != nil {
				return err
			}

			// Create a new router
			httpRouter := mux.NewRouter()
			httpRouter.HandleFunc("/", html.Doge())
			httpRouter.PathPrefix("/{[a-zA-Z0-9=-/]+}").HandlerFunc(html.All(conf))

			// Create a new server
			server := &http.Server{
				Addr:    fmt.Sprintf("%s:%d", address, port),
				Handler: handlers.LoggingHandler(os.Stdout, httpRouter),
			}

			// Start the server
			if c.String("ssl-cert") != "" && c.String("ssl-key") != "" {
				return server.ListenAndServeTLS(c.String("ssl-cert"), c.String("ssl-key"))
			} else {
				return server.ListenAndServe()
			}
		},
	}

	cmd.Register(command)
}
