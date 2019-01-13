package main

import (
	"fmt"
	"net/http"

	"firebase.google.com/go"
	"github.com/Sirupsen/logrus"
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/go-ozzo/ozzo-routing/cors"
	"github.com/jackinf/golang-goals/apis"
	"github.com/jackinf/golang-goals/app"
	"github.com/jackinf/golang-goals/daos"
	"github.com/jackinf/golang-goals/errors"
	"github.com/jackinf/golang-goals/services"
	_ "github.com/lib/pq"
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

	// firebase app
	firebaseApp := CreateFirebaseApp()

	// wire up API routing
	http.Handle("/", buildRouter(logger, db, firebaseApp))

	// start the server
	address := fmt.Sprintf(":%v", app.Config.ServerPort)
	logger.Infof("server %v is started at %v\n", app.Version, address)
	panic(http.ListenAndServe(address, nil))
}

func buildRouter(logger *logrus.Logger, db *dbx.DB, firebaseApp *firebase.App) *routing.Router {
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

	rg.Use(apis.FirebaseAuth(firebaseApp))

	goalDao := daos.NewGoalDAO()
	apis.ServeGoalResource(rg, services.NewGoalService(goalDao))

	// wire up more resource APIs here

	return router
}
