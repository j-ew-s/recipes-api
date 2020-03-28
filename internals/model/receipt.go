package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Receipt is a sctruct to describe an receipt object
type Receipt struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
	Link        string             `bson:"link,omitempty"`
	Rate        int                `bson:"rate,omitempty"`
	Tags        []string           `bson:"tags,omitempty"`
}

func createInstance() *Receipt {
	receipt := new(Receipt)
	return receipt
}

//IsValidToCreate is used to validate the obligatory fields Name and Link.
// When one of them is not valid (null or empty) will return false and otherwise true
func (r *Receipt) IsValidToCreate() bool {

	if len(r.Name) <= 0 {
		return false
	}

	if len(r.Link) <= 0 {
		return false
	}

	return true

}
