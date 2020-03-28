package model

// ReceiptCreate is a sctruct to describe an receipt object
type ReceiptCreate struct {
	ID      interface{} `json:"_id,omitempty"`
	Err     error       `json:"err,omitempty"`
	Message string      `json:"message,omitempty"`
	status  int
}

// CreateInstance creates new empty instance of ReceiptCreate
func (r *ReceiptCreate) CreateInstance() *ReceiptCreate {
	receipt := new(ReceiptCreate)
	return receipt
}

//SetStatus Sets a value to private status
func (r *ReceiptCreate) SetStatus(newStatus int) {
	r.status = newStatus
}

// GetStatus gets value from private status
func (r *ReceiptCreate) GetStatus() int {
	return r.status
}

//SetError sets the apropriated messages for info error
func (r *ReceiptCreate) SetError(err error, msg string, status int) {
	r.status = status
	r.Err = err
	r.Message = msg
}

// SetSuccess sets the approciated info dor successfuly created object
func (r *ReceiptCreate) SetSuccess(id interface{}, msg string, status int) {
	r.status = status
	r.ID = id
	r.Message = msg
}
