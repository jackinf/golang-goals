package main

import (
	"context"
	"log"

	"firebase.google.com/go"
	"google.golang.org/api/option"
)

func CreateFirebaseApp(firebaseCredentialsJson string) *firebase.App {
	opt := option.WithCredentialsJSON([]byte(firebaseCredentialsJson)) //Alternatively, you can use `opt := option.WithCredentialsFile("secrets/firebaseServiceAccountKey.json")`

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return app
}
