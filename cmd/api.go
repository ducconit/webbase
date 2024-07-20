package cmd

import (
	"codebase/api"
	"codebase/utils"
	"github.com/spf13/cobra"
	"log"
	"syscall"
	"time"
)

var apiServe = &cobra.Command{
	Use:   "api:serve",
	Short: "Run an instance of the API server",
	Long:  "Run an instance of the API server.",
	RunE: func(cmd *cobra.Command, args []string) error {
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			return err
		}
		port, err := cmd.Flags().GetString("port")
		if err != nil {
			return err
		}

		return runServer(utils.StringOf(host).Join(":", port).Value())
	},
}

func init() {
	RootCmd.AddCommand(apiServe)

	apiServe.Flags().StringP("port", "P", "8000", "Port opened for the server")
	apiServe.Flags().StringP("host", "H", "127.0.0.1", "Host binding for the server")
}

func runServer(addr string) error {
	srv := api.New()
	utils.RegisterOSSignalHandler(func() {
		log.Println("Shutting down the server...")
		// Shutdown gracefully shutdowns the server with
		// a timeout of 5 seconds.
		if err := srv.ShutdownWithTimeout(time.Second * 5); err != nil {
			// TODO: transport error to logging services
			panic(err)
		}
	}, syscall.SIGTERM, syscall.SIGINT)

	if err := srv.Listen(addr); err != nil {
		panic(err)
	}

	return nil
}
