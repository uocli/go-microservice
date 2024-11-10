package dberrors

type ConflictError struct{}

func (e *ConflictError) Error() string {
	return "attempting to create a record with an existing key"
}
