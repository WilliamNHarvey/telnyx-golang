package pkg

import (
	"net/http"

	"github.com/WilliamNHarvey/telnyx-golang/config"
	"github.com/WilliamNHarvey/telnyx-golang/internal"
	"github.com/WilliamNHarvey/telnyx-golang/internal/messaging"
	"github.com/WilliamNHarvey/telnyx-golang/internal/numbers"
	"github.com/WilliamNHarvey/telnyx-golang/internal/verify"
)

/*
*	the main telnyx struct type which contains all internal types with their methods (numbers, messaging, etc)
 */
type telnyx struct {
	Config     internal.Config
	httpClient http.Client
	Numbers    numbers.Numbers
	Messaging  *messaging.Messaging
	Verify     verify.Verify
}
type Telnyx = telnyx

/*
*	sets up config
 */
func setupConfig(t *telnyx, cfg map[string]string) {
	t.Config = internal.Config{
		Api: config.ApiKeys{
			V1:   cfg["v1"],
			V2:   cfg["v2"],
			User: cfg["user"],
		},
	}
}

/*
*	sets up the http client
 */
func setupHttpClient(t *telnyx) {
	t.httpClient = internal.InitHttpClient()
}

/*
*	initializes and returns a configured telnyx structure ready to be used
 */
func Init(cfg map[string]string) *telnyx {
	// make struct
	s := telnyx{}

	setupConfig(&s, cfg)
	setupHttpClient(&s)

	// init numbers
	s.Numbers = numbers.InitNumbers(&s.Config, &s.httpClient)

	// init messaging
	s.Messaging = messaging.InitMessaging(&s.Config, &s.httpClient)

	// init verify
	s.Verify = verify.InitVerify(&s.Config, &s.httpClient)

	return &s
}
