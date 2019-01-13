package apis

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/go-ozzo/ozzo-routing"
	"log"
	"net/http"
	"strings"
)

func FirebaseAuth(app *firebase.App) routing.Handler {
	return func(ctx *routing.Context) error {
		context := context.Background()
		client, err := app.Auth(context)
		if err != nil {
			log.Fatalf("error getting Auth client: %v\n", err)
		}

		authorization := ctx.Request.Header.Get("Authorization")
		ss := strings.Split(authorization, " ")
		tokenId := ss[len(ss)-1]

		token, err := client.VerifyIDToken(context, tokenId)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Token is not valid:", token)
			return routing.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		log.Printf("Verified ID token: %v\n", token)
		return nil
	}
}
