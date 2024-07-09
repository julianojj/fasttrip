package exceptions

type DomainException struct {
	message string
}

var (
	ErrPeriodNotAllowed  = NewDomainException("Period not allowed")
	ErrCapacityExceeded  = NewDomainException("Capacity exceeded")
	ErrRoomAlreadyExists = NewDomainException("Room already exists")
	ErrRequiredCategory  = NewDomainException("Category is required")
	ErrInvalidPrice      = NewDomainException("Invalid price")
	ErrInvalidCapacity   = NewDomainException("Invalid capacity")
)

var _ error = (*DomainException)(nil)

func NewDomainException(message string) *DomainException {
	return &DomainException{
		message,
	}
}

func (e *DomainException) Error() string {
	return e.message
}
