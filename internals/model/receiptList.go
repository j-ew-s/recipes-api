package model

// ReceiptList Represents a list of Receipts
type ReceiptList struct {
	Receipts []Receipt `json:"receipts,omitempty"`
	Quantity int       `json:"quantity,omitempty"`
	Err      error     `json:"err,omitempty"`
	Message  string    `json:"message,omitempty"`
	status   int
}

//SetStatus Sets a value to private status
func (r *ReceiptList) SetStatus(newStatus int) {
	r.status = newStatus
}

// GetStatus gets value from private status
func (r *ReceiptList) GetStatus() int {
	return r.status
}

//SetError sets the apropriated messages for info error
func (r *ReceiptList) SetError(err error, msg string) {
	r.status = 500
	r.Err = err
	r.Message = msg
}

// SetSuccess sets the approciated info dor successfuly created object
func (r *ReceiptList) SetSuccess(msg string, status int) {
	r.status = status
	r.Message = msg
	r.Quantity = len(r.Receipts)
}
