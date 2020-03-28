package usecase

import (
	"github.com/j-ew-s/receipts-api/internals/model"
	"github.com/j-ew-s/receipts-api/internals/repository"
)

//Create a receipt
//  Validates Required fields:
//   - Name and link.
//     - Returns message and http status 400
//   Validate duplicated Name and link.
//     - Returns Message and http status 400
//  Call Receipt Repository
//  Returns the model ReceiptCreate
func Create(receipt *model.Receipt) model.ReceiptCreate {

	isValidToCreate := receipt.IsValidToCreate()

	if isValidToCreate == false {
		receiptCreate := new(model.ReceiptCreate)
		receiptCreate.SetError(nil, "Name and Link are obligatory", 400)
		return *receiptCreate
	}

	haveDuplicated := validateDuplicated(receipt)

	if haveDuplicated.Receipts != nil {
		receiptCreate := new(model.ReceiptCreate)
		receiptCreate.SetError(nil, haveDuplicated.Message, haveDuplicated.GetStatus())
		return *receiptCreate
	}

	receiptCreated := repository.Create(receipt)

	return receiptCreated
}

// Delete a receipt
func Delete(id string) error {
	receiptsList := repository.Delete(id)

	return receiptsList
}

// Get a receipt
func Get() (res model.ReceiptList) {

	receiptsList := repository.Get()

	return receiptsList
}

// GetByID a receipt
func GetByID(id string) (res model.ReceiptList) {
	receipt := repository.GetByID(id)
	return receipt
}

// GetByTags a receipt
func GetByTags(tags []string) (res model.ReceiptList) {

	receiptsList := repository.GetByTags(tags)

	return receiptsList
}

// Update a receipt
func Update(receipt *model.Receipt, id string) (status int, err error) {

	var a = receipt.ID.Hex()

	if id != a {
		return 400, nil
	}

	receiptFound := GetByID(id)

	if err != nil {
		return 500, err
	}

	if &receiptFound == nil {
		return 404, nil
	}

	receiptDuplicated := validateDuplicated(receipt)

	if receiptDuplicated.Quantity > 0 {

		var duplicatedID = receiptDuplicated.Receipts[0].ID.Hex()
		if duplicatedID != a {
			return 409, nil
		}
	}

	err = repository.Update(receipt)

	if err != nil {
		return 500, nil
	}

	return 200, nil
}

// Search a receipt
func Search(title string) (*model.Receipt, error) {
	return nil, nil
}

//validateDuplicated checks for alredy existing objects with Name and Link duplicated.
func validateDuplicated(receipt *model.Receipt) model.ReceiptList {

	existingReceipts := repository.GetByNameOrLink(receipt)

	if len(existingReceipts.Receipts) > 0 {
		receiptList := new(model.ReceiptList)
		receiptList.Message = "There are already an Receipt with this Name or Link"
		receiptList.SetStatus(409)
		return *receiptList
	}

	return existingReceipts
}
