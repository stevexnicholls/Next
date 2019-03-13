package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/spf13/viper"
	log "github.com/stevexnicholls/next/logger"
	"github.com/stevexnicholls/next/restapi"
)

// User described a user that uses the API
type User struct {
	ID   int
	Role string
}

// Token parses and validates a token and return the logged in user
func Token(token string) (interface{}, error) {

	var user User

	t := viper.GetString("api_key")

	// If no api key was specified at startup then allow all
	if t == "" {
		user.ID = 1
		user.Role = "admin"
		if viper.Get("log_level") == "debug" {
			log.Infof("no api key specified")
		}
		return &user, nil
	}

	if token == "" {
		if viper.Get("log_level") == "debug" {
			log.Infof("no api key provided with call")
		}
		return nil, nil // unauthorized
	}

	// In a real authentication, here we should actually validate that the token is valid
	//err := json.Unmarshal([]byte(token), &user)

	if (token) == t {
		user.ID = 1
		user.Role = "admin"
		if viper.Get("log_level") == "debug" {
			log.Infof("admin authenticated")
		}
		return &user, nil
	}

	if viper.Get("log_level") == "debug" {
		log.Infof("api key provided (%v) did not match", token)
	}
	return nil, nil
}

// Request enforce policy on a given request
func Request(req *http.Request) error {
	var (
		route = middleware.MatchedRouteFrom(req)
		user  = FromContext(req.Context())
	)

	for _, auth := range route.Authenticators {
		scopes := auth.Scopes["token"]

		if len(scopes) == 0 {
			return nil // The token is valid for any user role
		}

		// Check if any of the scopes is the same as the user's role
		for _, scope := range scopes {
			if scope == user.Role {
				return nil
			}
		}
	}
	return fmt.Errorf("forbidden")
}

// FromContext extract the user from the context
func FromContext(ctx context.Context) *User {
	v := ctx.Value(restapi.AuthKey)
	if v == nil {
		return nil
	}
	return v.(*User)
}
