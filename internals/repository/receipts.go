package repository

import (
	"context"
	"fmt"
	"log"
	"os"

	"time"

	"github.com/j-ew-s/receipts-api/internals/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ReceiptRepository class
type ReceiptRepository struct {
	Client *mongo.Client
}

// NewReceiptRepository create instance.
//  Create an MongoConnect obect using ApplyURI and
//  a context with 10 seconds of execution.
//  Returns a new ReceiptRespository with the MongoConect Client
func NewReceiptRepository() (*ReceiptRepository, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	return &ReceiptRepository{client}, err
}

//Create an receipt
// Instanciate MongoDB from receiptDB and receipt table
// Prerepare the ID as index
// Inserts a new Receipt
// Returns model.ReceiptCreate object
//  When an error rise  ReceiptCreate.Err is set
//  When no errors ID is set with new _Id bson object
func Create(receipt *model.Receipt) model.ReceiptCreate {

	receiptCreate := new(model.ReceiptCreate)

	repository, err := NewReceiptRepository()
	if err != nil {
		receiptCreate.SetError(err, "Not able to create  NewReceiptRepository", 500)
		return *receiptCreate
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = repository.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		receiptCreate.SetError(err, "DB not responding", 500)
		return *receiptCreate
	}

	collection := repository.Client.Database("receiptDB").Collection("receipt")

	index := mongo.IndexModel{
		Keys: bson.M{
			"ID": 1, // index in ascending order
		}, Options: nil,
	}

	ind, err := collection.Indexes().CreateOne(ctx, index)

	if err != nil {
		fmt.Println("err ", ind)
		os.Exit(1) // exit in case of error
	}

	// Insert Datas
	response, err := collection.InsertOne(ctx, receipt)

	if err != nil {
		receiptCreate.SetError(err, "Error when Inserting on DB", 500)
		return *receiptCreate
	}

	receiptCreate.ID = response.InsertedID
	receiptCreate.Err = nil

	return *receiptCreate
}

//Delete an receipt by its ID
func Delete(id string) error {

	repository, err := NewReceiptRepository()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = repository.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	collection := repository.Client.Database("receiptDB").Collection("receipt")

	receiptID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println("Invalid ObjectID")
		return err
	}

	query := bson.M{"_id": bson.M{"$eq": receiptID}}

	_, err = collection.DeleteOne(ctx, query)

	if err != nil {
		return err
	}

	return nil
}

//Get an receipt
func Get() (res model.ReceiptList) {

	receiptList := model.ReceiptList{}
	repository, err := NewReceiptRepository()
	if err != nil {
		return receiptList
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = repository.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return receiptList
	}

	collection := repository.Client.Database("receiptDB").Collection("receipt")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(ctx, &receiptList.Receipts); err != nil {
		log.Fatal(err)
	}

	receiptList.SetSuccess("", 200)

	return receiptList
}

//GetByID an receipt
func GetByID(id string) (res model.ReceiptList) {

	receiptList := model.ReceiptList{}

	repository, err := NewReceiptRepository()
	if err != nil {
		receiptList.SetError(err, "Could not create NewREceiptRepository on Receipt GetById")
		return receiptList
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = repository.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		receiptList.SetError(err, "Could not ping repository on Receipt GetById")
		return receiptList
	}

	collection := repository.Client.Database("receiptDB").Collection("receipt")

	receiptID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		receiptList.SetError(err, "Invalid objectIDFromHex on Receipt GetById")
		return receiptList
	}

	response := collection.FindOne(ctx, bson.M{"_id": receiptID})

	re := model.Receipt{}
	response.Decode(&re)

	receiptList.Receipts = append(receiptList.Receipts, re)

	receiptList.SetSuccess(
		"Found",
		200)

	return receiptList
}

//GetByNameOrLink will get any object with same Name Or Link
func GetByNameOrLink(receipt *model.Receipt) (res model.ReceiptList) {

	responseReceiptList := model.ReceiptList{}

	repository, err := NewReceiptRepository()
	if err != nil {
		return responseReceiptList
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = repository.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return responseReceiptList
	}

	collection := repository.Client.Database("receiptDB").Collection("receipt")

	query := bson.M{
		"$or": []bson.M{
			bson.M{"name": receipt.Name},
			bson.M{"link": receipt.Link},
		},
	}

	response, err := collection.Find(ctx, query)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Resposne ")
	fmt.Println(response)
	if err = response.All(ctx, &responseReceiptList.Receipts); err != nil {
		log.Fatal(err)
	}

	fmt.Println("responseReceiptList.Receipts ")
	fmt.Println(responseReceiptList.Receipts)

	return responseReceiptList
}

//GetByTags an receipt
func GetByTags(tags []string) (res model.ReceiptList) {

	receiptList := model.ReceiptList{}

	repository, err := NewReceiptRepository()
	if err != nil {
		return receiptList
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = repository.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return receiptList
	}

	collection := repository.Client.Database("receiptDB").Collection("receipt")

	query := bson.M{"tags": bson.M{"$all": tags}}

	response, err := collection.Find(ctx, query)
	fmt.Println(response)
	if err != nil {
		log.Fatal(err)
	}

	if err = response.All(ctx, &receiptList.Receipts); err != nil {
		log.Fatal(err)
	}

	receiptList.SetSuccess("", 200)

	return receiptList
}

//Update an receipt
func Update(receipt *model.Receipt) error {

	repository, err := NewReceiptRepository()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = repository.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	collection := repository.Client.Database("receiptDB").Collection("receipt")

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": receipt.ID},
		bson.D{
			{"$set", bson.D{primitive.E{Key: "name", Value: receipt.Name}}},
			{"$set", bson.D{primitive.E{Key: "description", Value: receipt.Description}}},
			{"$set", bson.D{primitive.E{Key: "link", Value: receipt.Link}}},
			{"$set", bson.D{primitive.E{Key: "rate", Value: receipt.Rate}}},
			{"$set", bson.D{primitive.E{Key: "tags", Value: receipt.Tags}}},
		},
	)
	if err != nil {
		return err
	}

	fmt.Println(result)
	if result == nil {
		return nil
	}

	return nil
}

//Search an receipt
func Search(ctx context.Context, title string) (*model.Receipt, error) {
	return nil, nil
}
