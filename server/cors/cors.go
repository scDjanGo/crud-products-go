package cors

import (
	"time"

	"github.com/gin-contrib/cors"
)

func CORS_SETTINGS() cors.Config {
	return cors.Config{
		AllowOrigins: []string{
			"http://127.0.0.1:5500",
		},
		AllowMethods: []string{
			"GET", "POST", "PATCH", "DELETE",
		},
		AllowHeaders: []string{
			"Content-Type", "Authorization", "Origin",
		},
		MaxAge: 12 * time.Hour,
	}
}
