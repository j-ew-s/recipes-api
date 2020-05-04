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

	PrepareResponse(ctx, 200, response)

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

	PrepareResponse(ctx, createdRecipe.GetStatus(), createdRecipe)

}

// Delete removes a recipe fisically
// that match the ID parameter
// uses Recipe object
func Delete(ctx *fasthttp.RequestCtx) {

	id := fmt.Sprintf("%v", ctx.UserValue("id"))
	httpStatus := 200
	recipe := usecase.GetByID(id)
	response := true

	if recipe.Err != nil {
		ctx.Error(recipe.Err.Error(), recipe.GetStatus())
		response = false
		httpStatus = 404
	}

	if &recipe != nil {
		err := usecase.Delete(id)

		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
			response = false
			httpStatus = 500
		}

	}

	PrepareResponse(ctx, httpStatus, response)
}

// Get returns all the recipes
// No filtering is made
func Get(ctx *fasthttp.RequestCtx) {

	recipes := usecase.Get()
	if recipes.Err != nil {
		ctx.Error(recipes.Err.Error(), fasthttp.StatusInternalServerError)
	}

	PrepareResponse(ctx, recipes.GetStatus(), recipes)

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

	PrepareResponse(ctx, httpStatusCode, recipe)

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
		recipes.SetStatus(fasthttp.StatusInternalServerError)
	}

	PrepareResponse(ctx, recipes.GetStatus(), recipes)

}

// Put updates a recipe
// that match the ID parameter
// uses Recipe object
func Put(ctx *fasthttp.RequestCtx) {

	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)

	var httpStatus = 404

	var recipe model.Recipe

	id := fmt.Sprintf("%v", ctx.UserValue("id"))
	err := json.Unmarshal(ctx.PostBody(), &recipe)

	if err != nil {
		httpStatus = 500
	}

	httpStatus, err = usecase.Update(&recipe, id)

	PrepareResponse(ctx, httpStatus, recipe)

}

// Search performs search for recipes
// that contains the search term.
// Search is made upon all the filds
func Search(ctx *fasthttp.RequestCtx) {

	fmt.Println("Entrou no Search")

}

// PrepareResponse Prepare responses
func PrepareResponse(ctx *fasthttp.RequestCtx, code int, response interface{}) {

	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)

	ctx.Response.SetStatusCode(code)

	start := time.Now()

	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		elapsed := time.Since(start)
		fmt.Println(" ERROR : ", elapsed, err.Error(), response)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}

}
