# tollbooth_buffalo

[Buffalo](https://gobuffalo.io/) middleware for rate limiting HTTP requests.

## Usage
```go
import (
    "github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth/v7/limiter"
    "github.com/gobuffalo/envy"
)

var ENV = envy.Get("GO_ENV", "development")

func App() *buffalo.App {
	appOnce.Do(func() {
		app = buffalo.New(buffalo.Options{
			Env: ENV
		})

        lmt := tollbooth.NewLimiter(500, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Minute})
		
        lmt.SetMethods([]string{"GET", "POST", "PUT", "DELETE"}).SetIPLookups([]string{"X-Forwarded-For", "X-Authorization", "RemoteAddr", "X-Real-IP"})
		
        lmt.SetMessage("Rate Limit Exceeded.").SetMessageContentType("application/json; charset=utf-8")

        app.Use(middlewares.RateLimitMiddleware(lmt))

		app.GET("/", HomeHandler)
    })

    return app
}
```