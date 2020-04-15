package receiptinterface

import (
	"github.com/j-ew-s/recipes-api/internals/model"
)

// Repository represent the Recipe's repository contract
type Repository interface {
	Create(receipt *model.Receipt) (int64, error)
	Delete(id int64) error
	Get() (res []*model.Receipt)
	GetByID(id int64) (*model.Receipt, error)
	GetByTags(tags []string) (*model.Receipt, error)
	Update(receipt *model.Receipt) error
	Search(title string) (*model.Receipt, error)
}
