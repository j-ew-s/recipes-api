package usecase

import (
	"github.com/j-ew-s/recipes-api/internals/model"
	"github.com/j-ew-s/recipes-api/internals/repository"
)

//Create a recipe
//  Validates Required fields:
//   - Name and link.
//     - Returns message and http status 400
//   Validate duplicated Name and link.
//     - Returns Message and http status 400
//  Call Recipe Repository
//  Returns the model RecipeCreate
func Create(recipe *model.Recipe) model.RecipeCreate {

	isValidToCreate := recipe.IsValidToCreate()

	if isValidToCreate == false {
		recipeCreate := new(model.RecipeCreate)
		recipeCreate.SetError(nil, "Name and Link are obligatory", 400)
		return *recipeCreate
	}

	haveDuplicated := isDuplicated(recipe)

	if haveDuplicated == true {
		recipeCreate := new(model.RecipeCreate)
		recipeCreate.SetError(nil, "There are already one item with same Name and Link.", 409)
		return *recipeCreate
	}

	recipeCreated := repository.Create(recipe)

	return recipeCreated
}

// Delete a recipe
func Delete(id string) error {
	recipesList := repository.Delete(id)

	return recipesList
}

// Get a recipe
func Get() (res model.RecipeList) {

	recipesList := repository.Get()

	return recipesList
}

// GetByID a recipe
func GetByID(id string) (res model.RecipeList) {
	recipe := repository.GetByID(id)
	return recipe
}

// GetByTags a recipe
func GetByTags(tags []string) (res model.RecipeList) {

	recipesList := repository.GetByTags(tags)

	return recipesList
}

// Update a recipe
func Update(recipe *model.Recipe, id string) (status int, err error) {

	var a = recipe.ID.Hex()

	if id != a {
		return 400, nil
	}

	recipeFound := GetByID(id)

	if err != nil {
		return 500, err
	}

	if &recipeFound == nil {
		return 404, nil
	}

	recipeDuplicated := repository.GetByNameOrLink(recipe)

	recipesFoundLen := len(recipeDuplicated.Recipes)

	if recipesFoundLen > 0 {

		var duplicatedID = recipeDuplicated.Recipes[0].ID.Hex()
		if duplicatedID != a {
			return 409, nil
		}
	}

	err = repository.Update(recipe)

	if err != nil {
		return 500, nil
	}

	return 200, nil
}

// Search a recipe
func Search(title string) (*model.Recipe, error) {
	return nil, nil
}

//isDuplicated checks for alredy existing objects with Name and Link duplicated.
func isDuplicated(recipe *model.Recipe) bool {

	existingRecipes := repository.GetByNameOrLink(recipe)

	var recipesFound = len(existingRecipes.Recipes)

	return recipesFound > 0
}
