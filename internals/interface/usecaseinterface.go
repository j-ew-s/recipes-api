package receiptinterface

import (
	"github.com/j-ew-s/receipts-api/internals/model"
)

// UseCase represent the Receipt's UseCase contract
type UseCase interface {
	Create(receipt *model.Receipt) (int64, error)
	Delete(id int64) error
	Get() (res []*model.Receipt)
	GetByID(id int64) (*model.Receipt, error)
	GetByTags(tags []string) (*model.Receipt, error)
	Update(receipt *model.Receipt) error
	Search(title string) (*model.Receipt, error)
}
