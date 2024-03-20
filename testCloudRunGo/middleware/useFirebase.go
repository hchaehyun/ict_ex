package middleware

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/messaging"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"strings"
)

// Set up Firebase Auth client globally
var authClient *auth.Client
var firestoreClient *firestore.Client
var app *firebase.App
var messageClient *messaging.Client

func init() {
	// Initialize Firebase Auth client
	opt := option.WithCredentialsFile(".key/vitalring-seoul-firebase-adminsdk-hmtib-79b2cd7ace.json")
	config := &firebase.Config{ProjectID: "vitalring-seoul"}
	fireApp, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase App: %v", err)
	}
	app = fireApp

	authClient, err = fireApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("Error initializing Firebase Auth client: %v", err)
	}

	firestoreClient, err = fireApp.Firestore(context.Background())
	if err != nil {
		log.Fatalf("Error initializing Firebase Firestore client: %v", err)
	}

	messageClient, err = fireApp.Messaging(context.Background())
	if err != nil {
		log.Fatalf("Error initializing Firebase Messaging client: %v", err)
	}
}

// Authorization is a middleware function that extracts the "Authorization" header (expected to be a Firebase idToken)
// and sets the authenticated user's ID in the request context.
func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		idToken := c.GetHeader("Authorization")
		// replace "Bearer " 가 있을 경우 제거
		idToken = strings.Replace(idToken, "Bearer ", "", 1)

		token, err := authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			c.Abort()
			log.Printf("Error verifying ID token: %v\n", err)
			return
		}

		// Set user ID in context
		c.Set("fireApp", app)
		c.Set("userId", token.UID)
		c.Set("authToken", token)
	}
}

func OnlyAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken, _ := c.MustGet("authToken").(*auth.Token)
		if !authToken.Claims["admin"].(bool) {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

func UseFirestore() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("firestoreClient", firestoreClient)
		c.Next()
	}
}

func UseMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("messageClient", messageClient)
		c.Next()
	}
}

func UseAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("authClient", authClient)
		c.Next()
	}
}

func UseStorageBucket() gin.HandlerFunc {
	return func(c *gin.Context) {
		storage, err := app.Storage(context.Background())
		if err != nil {
			log.Fatalf("Error initializing Firebase Storage client: %v", err)
		}
		bucket, err := storage.Bucket("vitalring-seoul.appspot.com")
		c.Set("bucket", bucket)
		c.Next()
	}
}
