package recipeinterface

import (
	"github.com/j-ew-s/recipes-api/internals/model"
)

// Repository represent the Recipe's repository contract
type Repository interface {
	Create(recipe *model.Recipe) (int64, error)
	Delete(id int64) error
	Get() (res []*model.Recipe)
	GetByID(id int64) (*model.Recipe, error)
	GetByTags(tags []string) (*model.Recipe, error)
	Update(recipe *model.Recipe) error
	Search(title string) (*model.Recipe, error)
}
