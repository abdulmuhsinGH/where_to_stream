package firestore

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

/*
New sets up connection to firestore
*/
func New() (*firestore.Client, error) {
	// Sets your Google Cloud Platform project ID.
	projectID := "wheretostreamapp"

	// Get a Firestore client.
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil, err
	}
	// Close client when done.
	defer client.Close()

	return client, nil
}
