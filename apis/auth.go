package apis

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/auth"
	"github.com/jackinf/golang-goals/app"
	"github.com/jackinf/golang-goals/errors"
	"github.com/jackinf/golang-goals/models"

	"github.com/auth0-community/auth0"
	jose "gopkg.in/square/go-jose.v2"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Auth(signingKey string) routing.Handler {
	return func(c *routing.Context) error {
		var credential Credential
		if err := c.Read(&credential); err != nil {
			return errors.Unauthorized(err.Error())
		}

		identity := authenticate(credential)
		if identity == nil {
			return errors.Unauthorized("invalid credential")
		}

		token, err := auth.NewJWT(jwt.MapClaims{
			"id":   identity.GetID(),
			"name": identity.GetName(),
			"exp":  time.Now().Add(time.Hour * 72).Unix(),
		}, signingKey)
		if err != nil {
			return errors.Unauthorized(err.Error())
		}

		return c.Write(map[string]string{
			"token": token,
		})
	}
}

//func AuthMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		secret := []byte("{YOUR-AUTH0-API-SECRET}")
//		secretProvider := auth0.NewKeyProvider(secret)
//		audience := []string{"{YOUR-AUTH0-API-AUDIENCE}"}
//
//		configuration := auth0.NewConfiguration(secretProvider, audience, "https://{YOUR-AUTH0-DOMAIN}.auth0.com/", jose.HS256)
//		validator := auth0.NewValidator(configuration, nil)
//
//		token, err := validator.ValidateRequest(r)
//
//		if err != nil {
//			fmt.Println(err)
//			fmt.Println("Token is not valid:", token)
//			w.WriteHeader(http.StatusUnauthorized)
//			w.Write([]byte("Unauthorized"))
//		} else {
//			next.ServeHTTP(w, r)
//		}
//	})
//}

func AuthMiddleware1() routing.Handler {
	return func(c *routing.Context) error {
		secret := []byte("{YOUR-AUTH0-API-SECRET}")
		secretProvider := auth0.NewKeyProvider(secret)
		audience := []string{"{YOUR-AUTH0-API-AUDIENCE}"}

		configuration := auth0.NewConfiguration(secretProvider, audience, "https://{YOUR-AUTH0-DOMAIN}.auth0.com/", jose.HS256)
		validator := auth0.NewValidator(configuration, nil)

		token, err := validator.ValidateRequest(c.Request)

		if err != nil {
			fmt.Println(err)
			fmt.Println("Token is not valid:", token)
			return routing.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		} else {
			return nil
		}
	}
}

func authenticate(c Credential) models.Identity {
	if c.Username == "demo" && c.Password == "pass" {
		return &models.User{ID: "100", Name: "demo"}
	}
	return nil
}

func JWTHandler(c *routing.Context, j *jwt.Token) error {
	userID := j.Claims.(jwt.MapClaims)["id"].(string)
	app.GetRequestScope(c).SetUserID(userID)
	return nil
}
