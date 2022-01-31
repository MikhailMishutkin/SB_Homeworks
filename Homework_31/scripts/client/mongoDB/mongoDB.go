package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectAndPing() {
	// Create a Client to a MongoDB server and use Ping to verify that the
	// server is running.

	clientOpts := options.Client().ApplyURI("mongodb+srv://Mishutkin:1Vfcmrf1@cluster0.3uyrh.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	// Call Ping to verify that the deployment is up and the Client was
	// configured successfully. As mentioned in the Ping documentation, this
	// reduces application resiliency as the server may be temporarily
	// unavailable when Ping is called.
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	collection := client.Database("httpservice").Collection("friendsService")

	res, err := collection.InsertOne(ctx context.Context, handlers.DataDB)
	id := res.InsertedID
}

func CreateDB() {

}
