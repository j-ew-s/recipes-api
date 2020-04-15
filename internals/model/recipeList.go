package model

// RecipeList Represents a list of Recipes
type RecipeList struct {
	Recipes  []Recipe `json:"recipes,omitempty"`
	Quantity int      `json:"quantity,omitempty"`
	Err      error    `json:"err,omitempty"`
	Message  string   `json:"message,omitempty"`
	status   int
}

//SetStatus Sets a value to private status
func (r *RecipeList) SetStatus(newStatus int) {
	r.status = newStatus
}

// GetStatus gets value from private status
func (r *RecipeList) GetStatus() int {
	return r.status
}

//SetError sets the apropriated messages for info error
func (r *RecipeList) SetError(err error, msg string) {
	r.status = 500
	r.Err = err
	r.Message = msg
}

// SetSuccess sets the approciated info dor successfuly created object
func (r *RecipeList) SetSuccess(msg string, status int) {
	r.status = status
	r.Message = msg
	r.Quantity = len(r.Recipes)
}
