package cmd

import (
	"fmt"
	"grpc-middleware/middleware"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(serverCmd)
}

var version = "v0.0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("grpc middleware --%v", version)
		fmt.Println()
	},
}

var serverCmd = &cobra.Command{
	Use:   "start",
	Short: "start server",
	Run: func(cmd *cobra.Command, args []string) {
		router := mux.NewRouter()
		router.Use(middleware.JwtAuthentication)

		port := os.Getenv("PORT")
		if port == "" {
			port = "8001"
		}

		fmt.Print("Server is start: http://localhost:", port)

		err := http.ListenAndServe(":"+port, router)

		if err != nil {
			fmt.Println(err)
		}

	},
}
