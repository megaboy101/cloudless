package user

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type FirebaseRepository struct {
	app        *firebase.App
	authClient *auth.Client
}

func (r *FirebaseRepository) Connect() (Repository, error) {
	// Initialize the Firebase SDK
	opt := option.WithCredentialsFile("config/cloudless-prototype-firebase-adminsdk-q4d0n-c1cd79b3a0.json")
	firebaseApp, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v\n", err)
	}

	r.app = firebaseApp

	// Initialize the authentication client
	client, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("Error initializing Firebase Auth client: %v\n", err)
	}

	r.authClient = client

	return r, nil
}
