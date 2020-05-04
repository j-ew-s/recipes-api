package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Recipe is a sctruct to describe an recipe object
type Recipe struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Link        string             `json:"link,omitempty" bson:"link,omitempty"`
	Rate        int                `json:"rate,omitempty" bson:"rate,omitempty"`
	Tags        []string           `json:"tags,omitempty" bson:"tags,omitempty"`
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
