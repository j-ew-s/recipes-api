package recipescontroller

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/j-ew-s/recipes-api/configs"

	"github.com/j-ew-s/recipes-api/internals/usecase"

	"github.com/j-ew-s/recipes-api/internals/model"

	"github.com/valyala/fasthttp"
)

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
)

// Pong export
type Pong struct {
	Time            time.Time `json:"time,omitempty"`
	MongoConnection string    `json:"mongoConnection,omitempty"`
	MongoUser       string    `json:"mongoUser,omitempty"`
	ServerPort      string    `json:"serverPort,omitempty"`
	ServerAPI       string    `json:"serverAPI,omitempty"`
}

// Ping export
func Ping(ctx *fasthttp.RequestCtx) {

	response := new(Pong)

	response.Time = time.Now()
	response.MongoConnection = configs.MongoDBConfig.MongoServer
	response.MongoUser = configs.MongoDBConfig.User
	response.ServerPort = configs.ServerConfig.APIPort
	response.ServerAPI = configs.ServerConfig.APIServer

	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}

// Create insert a new Recipe.
//  Call Recipe UseCase
//  Prepare response and return model.RecipeCreat
func Create(ctx *fasthttp.RequestCtx) {

	var recipe model.Recipe

	err := json.Unmarshal(ctx.PostBody(), &recipe)
	if err != nil {
		panic(err)
	}

	createdRecipe := usecase.Create(&recipe)
	if createdRecipe.Err != nil {
		ctx.Error(createdRecipe.Err.Error(), fasthttp.StatusInternalServerError)
	}

	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
	ctx.Response.SetStatusCode(createdRecipe.GetStatus())

	start := time.Now()
	if err := json.NewEncoder(ctx).Encode(createdRecipe); err != nil {
		elapsed := time.Since(start)
		fmt.Println(" ERROR : ", elapsed, err.Error(), createdRecipe.Err)
		fmt.Println(" Message : ", createdRecipe.Message)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}

}

// Delete removes a recipe fisically
// that match the ID parameter
// uses Recipe object
func Delete(ctx *fasthttp.RequestCtx) {

	id := fmt.Sprintf("%v", ctx.UserValue("id"))

	recipe := usecase.GetByID(id)
	response := true

	if recipe.Err != nil {
		ctx.Error(recipe.Err.Error(), recipe.GetStatus())
		response = false
	}

	if &recipe != nil {
		err := usecase.Delete(id)

		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			response = false
		}

	}

	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)

	ctx.Response.SetStatusCode(200)

	start := time.Now()
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		elapsed := time.Since(start)
		fmt.Println(" ERROR : ", elapsed, err.Error(), err)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}

// Get returns all the recipes
// No filtering is made
func Get(ctx *fasthttp.RequestCtx) {

	recipes := usecase.Get()
	if recipes.Err != nil {
		ctx.Error(recipes.Err.Error(), fasthttp.StatusInternalServerError)
	}

	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
	ctx.Response.SetStatusCode(recipes.GetStatus())

	start := time.Now()
	if err := json.NewEncoder(ctx).Encode(recipes); err != nil {
		elapsed := time.Since(start)
		fmt.Println(" ERROR : ", elapsed, err.Error(), recipes.Err)
		fmt.Println(" Message : ", recipes.Message)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}

// GetByID will return any recipe with related ID
func GetByID(ctx *fasthttp.RequestCtx) {

	id := fmt.Sprintf("%v", ctx.UserValue("id"))

	recipe := usecase.GetByID(id)
	if recipe.Err != nil {
		ctx.Error(recipe.Err.Error(), recipe.GetStatus())
	}

	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)

	var httpStatusCode = 404

	if &recipe != nil {
		httpStatusCode = 200
	}

	ctx.Response.SetStatusCode(httpStatusCode)

	start := time.Now()
	if err := json.NewEncoder(ctx).Encode(recipe); err != nil {
		elapsed := time.Since(start)
		fmt.Println(" ERROR : ", elapsed, err.Error(), err)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}

}

// GetByTags performs search for a recipe Tags
// that contais the search parameter
// search is made upon only on Tag property
func GetByTags(ctx *fasthttp.RequestCtx) {

	tags := []string{}

	tagsQueryByte := ctx.QueryArgs().PeekMulti("tags")

	for i := 0; i < len(tagsQueryByte); i++ {
		tags = append(tags, string(tagsQueryByte[i]))
	}

	recipes := usecase.GetByTags(tags)
	if recipes.Err != nil {
		ctx.Error(recipes.Err.Error(), fasthttp.StatusInternalServerError)
	}

	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
	ctx.Response.SetStatusCode(recipes.GetStatus())

	start := time.Now()
	if err := json.NewEncoder(ctx).Encode(recipes); err != nil {
		elapsed := time.Since(start)
		fmt.Println(" ERROR : ", elapsed, err.Error(), recipes.Err)
		fmt.Println(" Message : ", recipes.Message)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}

}

// Put updates a recipe
// that match the ID parameter
// uses Recipe object
func Put(ctx *fasthttp.RequestCtx) {

	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)

	var httpStatusCode = 404

	var recipe model.Recipe

	id := fmt.Sprintf("%v", ctx.UserValue("id"))
	err := json.Unmarshal(ctx.PostBody(), &recipe)

	if err != nil {
		httpStatusCode = 500
	}

	httpStatusCode, err = usecase.Update(&recipe, id)

	ctx.Response.SetStatusCode(httpStatusCode)

	start := time.Now()
	if err := json.NewEncoder(ctx).Encode(true); err != nil {
		elapsed := time.Since(start)
		fmt.Println(" ERROR : ", elapsed, err.Error(), err)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}

}

// Search performs search for recipes
// that contains the search term.
// Search is made upon all the filds
func Search(ctx *fasthttp.RequestCtx) {

	fmt.Println("Entrou no Search")

}
