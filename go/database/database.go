package database

import (
	"context"
	"fmt"
	"log"
	"samfaina/dev/rootly_fetch/config"
	"samfaina/dev/rootly_fetch/models"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// conf App environment config
var conf *config.Config

func init() {

	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	conf = config.New()

}

// SaveIDSToDatabase handle save IDS operations
func SaveIDSToDatabase(plants []int) {
	// get database connection
	client, ctx, ctxCancel := GetDatabaseConnection()

	// clean up ids collection from database
	cleanUp(ctx, client)

	// save to mongodb atlas
	insertedID := saveIDS(ctx, client, plants)

	if conf.DebugMode {
		log.Printf("SAVED SUCCESSFULLY! %v\n", insertedID)
	}

	// disconnect from database
	disconnect(ctx, client)

	ctxCancel()

}

// GetPlantsIDList ...
func GetPlantsIDList() []int {
	// get database connection
	client, ctx, ctxCancel := GetDatabaseConnection()

	// result type
	var result struct {
		Ids []int `bson:"ids"`
	}
	// setup projection
	projection := bson.D{primitive.E{Key: "_id", Value: 0}}

	// fetch id list from database
	err := client.Database("rootly").Collection("plant_id").FindOne(
		ctx,
		bson.D{}, options.FindOne().SetProjection(projection)).Decode(&result)

	if err != nil {
		log.Fatal(err)
		ctxCancel()
		return nil
	}

	// disconnect from database
	disconnect(ctx, client)

	// release resources
	ctxCancel()

	return result.Ids
}

// PRIVATE METHODS

// GetDatabaseConnection returns mongo Client object and Context
func GetDatabaseConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	mongoURI := fmt.Sprintf("mongodb+srv://%s:%s@%s", conf.Mongo.Username, conf.Mongo.PWD, conf.Mongo.Host)
	if conf.DebugMode {
		log.Println("Connection string is:", mongoURI)
	}
	// fmt.Println("connection string is:", mongoURI)

	ctx, ctxCancel := context.WithCancel(context.Background())
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	if err != nil {
		log.Fatalf("[ERROR] GetDatabaseConnection has thrown an error: %v\n", err)
	}

	return client, ctx, ctxCancel
}

// disconnect and release all resources
func disconnect(ctx context.Context, client *mongo.Client) {

	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatalf("[ERROR] DisconnectFromDatabase has thrown an error: %v\n", err)
	}

	if conf.DebugMode {
		log.Println("Disconnected successfully!")
	}
}

// cleanUp in order to keep document up to date
func cleanUp(ctx context.Context, client *mongo.Client) {
	err := client.Database("rootly").Collection("plant_id").Drop(ctx)

	if err != nil {
		log.Fatalf("[ERROR] CleanUpPreviousCollection has thrown an error: %v\n", err)
	}

	if conf.DebugMode {
		log.Println("IDS Collection cleaned successfully")
	}

}

// saveIDS all fetched Ids
func saveIDS(ctx context.Context, client *mongo.Client, plants []int) interface{} {
	collection := client.Database("rootly").Collection("plant_id")

	res, err := collection.InsertOne(ctx, bson.M{"ids": plants})

	if err != nil {
		log.Fatalf("[ERROR] CleanUpPreviousCollection has thrown an error: %v\n", err)
	}

	return res.InsertedID
}

func requireCursorLength(cursor *mongo.Cursor) int {
	i := 0
	for cursor.Next(context.Background()) {
		i++
	}

	return i
}

// SavePlant to database
func SavePlant(ctx context.Context, client *mongo.Client, plant models.Plant) {
	collection := client.Database("rootly").Collection("plants")

	var operations []mongo.WriteModel

	operation := mongo.NewReplaceOneModel()

	operation.SetFilter(bson.D{primitive.E{Key: "_id", Value: plant.ID}})
	operation.SetUpsert(true)
	operation.SetReplacement(plant)

	operations = append(operations, operation)

	_, err := collection.BulkWrite(ctx, operations)
	if err != nil {
		log.Fatalf("[ERROR] BulkWrite has thrown an error: %v\n", err)
	}

}
