package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Recipe is a sctruct to describe an recipe object
type Recipe struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
	Link        string             `bson:"link,omitempty"`
	Rate        int                `bson:"rate,omitempty"`
	Tags        []string           `bson:"tags,omitempty"`
}

func createInstance() *Recipe {
	recipe := new(Recipe)
	return recipe
}

//IsValidToCreate is used to validate the obligatory fields Name and Link.
// When one of them is not valid (null or empty) will return false and otherwise true
func (r *Recipe) IsValidToCreate() bool {

	if len(r.Name) <= 0 {
		return false
	}

	if len(r.Link) <= 0 {
		return false
	}

	return true

}
