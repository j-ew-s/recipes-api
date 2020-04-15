package repository

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/j-ew-s/recipes-api/configs"

	"time"

	"github.com/j-ew-s/recipes-api/internals/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// RecipeRepository class
type RecipeRepository struct {
	Client *mongo.Client
}

// NewRecipeRepository create instance.
//  Create an MongoConnect obect using ApplyURI and
//  a context with 10 seconds of execution.
//  Returns a new RecipeRespository with the MongoConect Client
func NewRecipeRepository() (*RecipeRepository, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(configs.MongoDBConfig.MongoServer))

	return &RecipeRepository{client}, err
}

//Create an recipe
// Instanciate MongoDB from recipeDB and recipe table
// Prerepare the ID as index
// Inserts a new Recipe
// Returns model.RecipeCreate object
//  When an error rise  RecipeCreate.Err is set
//  When no errors ID is set with new _Id bson object
func Create(recipe *model.Recipe) model.RecipeCreate {

	recipeCreate := new(model.RecipeCreate)

	repository, err := NewRecipeRepository()
	if err != nil {
		recipeCreate.SetError(err, "Not able to create  NewRecipeRepository", 500)
		return *recipeCreate
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = repository.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		recipeCreate.SetError(err, "DB not responding", 500)
		return *recipeCreate
	}

	collection := repository.Client.Database("recipeDB").Collection("recipe")

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
	response, err := collection.InsertOne(ctx, recipe)

	if err != nil {
		recipeCreate.SetError(err, "Error when Inserting on DB", 500)
		return *recipeCreate
	}

	recipeCreate.ID = response.InsertedID
	recipeCreate.Err = nil

	return *recipeCreate
}

//Delete an recipe by its ID
func Delete(id string) error {

	repository, err := NewRecipeRepository()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = repository.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	collection := repository.Client.Database("recipeDB").Collection("recipe")

	recipeID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println("Invalid ObjectID")
		return err
	}

	query := bson.M{"_id": bson.M{"$eq": recipeID}}

	_, err = collection.DeleteOne(ctx, query)

	if err != nil {
		return err
	}

	return nil
}

//Get an recipe
func Get() (res model.RecipeList) {

	recipeList := model.RecipeList{}
	repository, err := NewRecipeRepository()
	if err != nil {
		return recipeList
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = repository.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return recipeList
	}

	collection := repository.Client.Database("recipeDB").Collection("recipe")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(ctx, &recipeList.Recipes); err != nil {
		log.Fatal(err)
	}

	recipeList.SetSuccess("", 200)

	return recipeList
}

//GetByID an recipe
func GetByID(id string) (res model.RecipeList) {

	recipeList := model.RecipeList{}

	repository, err := NewRecipeRepository()
	if err != nil {
		recipeList.SetError(err, "Could not create NewREceiptRepository on Recipe GetById")
		return recipeList
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = repository.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		recipeList.SetError(err, "Could not ping repository on Recipe GetById")
		return recipeList
	}

	collection := repository.Client.Database("recipeDB").Collection("recipe")

	recipeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		recipeList.SetError(err, "Invalid objectIDFromHex on Recipe GetById")
		return recipeList
	}

	response := collection.FindOne(ctx, bson.M{"_id": recipeID})

	re := model.Recipe{}
	response.Decode(&re)

	recipeList.Recipes = append(recipeList.Recipes, re)

	recipeList.SetSuccess(
		"Found",
		200)

	return recipeList
}

//GetByNameOrLink will get any object with same Name Or Link
func GetByNameOrLink(recipe *model.Recipe) (res model.RecipeList) {

	responseRecipeList := model.RecipeList{}

	repository, err := NewRecipeRepository()
	if err != nil {
		return responseRecipeList
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = repository.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return responseRecipeList
	}

	collection := repository.Client.Database("recipeDB").Collection("recipe")

	query := bson.M{
		"$or": []bson.M{
			bson.M{"name": recipe.Name},
			bson.M{"link": recipe.Link},
		},
	}

	response, err := collection.Find(ctx, query)

	if err != nil {
		log.Fatal(err)
	}

	if err = response.All(ctx, &responseRecipeList.Recipes); err != nil {
		log.Fatal(err)
	}

	return responseRecipeList
}

//GetByTags an recipe
func GetByTags(tags []string) (res model.RecipeList) {

	recipeList := model.RecipeList{}

	repository, err := NewRecipeRepository()
	if err != nil {
		return recipeList
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = repository.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return recipeList
	}

	collection := repository.Client.Database("recipeDB").Collection("recipe")

	query := bson.M{"tags": bson.M{"$all": tags}}

	response, err := collection.Find(ctx, query)
	fmt.Println(response)
	if err != nil {
		log.Fatal(err)
	}

	if err = response.All(ctx, &recipeList.Recipes); err != nil {
		log.Fatal(err)
	}

	recipeList.SetSuccess("", 200)

	return recipeList
}

//Update an recipe
func Update(recipe *model.Recipe) error {

	repository, err := NewRecipeRepository()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = repository.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	collection := repository.Client.Database("recipeDB").Collection("recipe")

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": recipe.ID},
		bson.D{
			{"$set", bson.D{primitive.E{Key: "name", Value: recipe.Name}}},
			{"$set", bson.D{primitive.E{Key: "description", Value: recipe.Description}}},
			{"$set", bson.D{primitive.E{Key: "link", Value: recipe.Link}}},
			{"$set", bson.D{primitive.E{Key: "rate", Value: recipe.Rate}}},
			{"$set", bson.D{primitive.E{Key: "tags", Value: recipe.Tags}}},
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

//Search an recipe
func Search(ctx context.Context, title string) (*model.Recipe, error) {
	return nil, nil
}
