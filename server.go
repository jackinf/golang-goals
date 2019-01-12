package main

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/go-ozzo/ozzo-routing"
	//"github.com/go-ozzo/ozzo-routing/auth"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/go-ozzo/ozzo-routing/cors"
	"github.com/jackinf/golang-goals/apis"
	"github.com/jackinf/golang-goals/app"
	"github.com/jackinf/golang-goals/daos"
	"github.com/jackinf/golang-goals/errors"
	"github.com/jackinf/golang-goals/services"
	_ "github.com/lib/pq"

	"github.com/auth0-community/auth0"
	jose "gopkg.in/square/go-jose.v2"
)

func main() {
	// load application configurations
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	// load error messages
	if err := errors.LoadMessages(app.Config.ErrorFile); err != nil {
		panic(fmt.Errorf("Failed to read the error message file: %s", err))
	}

	// create the logger
	logger := logrus.New()

	// connect to the database
	db, err := dbx.MustOpen("postgres", app.Config.DSN)
	if err != nil {
		panic(err)
	}
	db.LogFunc = logger.Infof

	// wire up API routing
	http.Handle("/", buildRouter(logger, db))

	// start the server
	address := fmt.Sprintf(":%v", app.Config.ServerPort)
	logger.Infof("server %v is started at %v\n", app.Version, address)
	panic(http.ListenAndServe(address, nil))
}

func buildRouter(logger *logrus.Logger, db *dbx.DB) *routing.Router {
	router := routing.New()

	router.To("GET,HEAD", "/ping", func(c *routing.Context) error {
		c.Abort() // skip all other middlewares/handlers
		return c.Write("OK " + app.Version)
	})

	router.Use(
		app.Init(logger),
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.Options{
			AllowOrigins: "*",
			AllowHeaders: "*",
			AllowMethods: "*",
		}),
		app.Transactional(db),
	)

	rg := router.Group("/v1")

	rg.Post("/auth", apis.Auth(app.Config.JWTSigningKey))
	rg.Use(authMiddleware1())
	//rg.Use(auth.JWT(app.Config.JWTVerificationKey, auth.JWTOptions{
	//	SigningMethod: app.Config.JWTSigningMethod,
	//	TokenHandler:  apis.JWTHandler,
	//}))

	goalDao := daos.NewGoalDAO()
	apis.ServeGoalResource(rg, services.NewGoalService(goalDao))

	// wire up more resource APIs here

	return router
}

func authMiddleware1() routing.Handler {
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
