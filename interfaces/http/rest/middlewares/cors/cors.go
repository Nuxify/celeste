package cors

import (
	"github.com/go-chi/cors"

	corsConfig "celeste/configs/cors"
)

// Init initializes the CORS configuration
func Init() *cors.Cors {
	config := corsConfig.Config{}

	// instantiate cors rule
	cors := cors.New(cors.Options{
		AllowedOrigins:   config.AllowedOrigins(),
		AllowedMethods:   config.AllowedMethods(),
		AllowedHeaders:   config.AllowedHeaders(),
		ExposedHeaders:   config.ExposedHeaders(),
		AllowCredentials: config.AllowCredentials(),
		MaxAge:           config.MaxAge(),
	})

	return cors
}
