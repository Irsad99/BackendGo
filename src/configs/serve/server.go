package serve

import (
	"log"
	"net/http"
	"os"

	"BackendGo/src/routers"

	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "server",
	Short: "start api server",
	RunE:  serve,
}

func serve(cmd *cobra.Command, args []string) error {
	mainRoute, err := routers.New()

	// headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	// methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	// origins := handlers.AllowedOrigins([]string{"http://localhost:8080/"})

	c := cors.New(cors.Options{
		AllowedHeaders:   []string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedOrigins:   []string{"https://www.google.com", "http://localhost:3000/"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	handler := c.Handler(mainRoute)
	// mainRoute.Use(mainRoute.W)

	if err == nil {
		var addrs string = ""

		if pr := os.Getenv("PORT"); pr != "" {
			addrs += pr
		}

		log.Println("App running on server " + addrs)

		if err := http.ListenAndServe(addrs, handler); err != nil {
			return err
		}

		return nil

	} else {
		return err
	}
}