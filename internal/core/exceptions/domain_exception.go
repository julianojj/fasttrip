package exceptions

type DomainException struct {
	message string
}

var (
	ErrPeriodNotAllowed   = NewDomainException("Period not allowed")
	ErrCapacityExceeded   = NewDomainException("Capacity exceeded")
	ErrRoomAlreadyExists  = NewDomainException("Room already exists")
	ErrRequiredCategory   = NewDomainException("Category is required")
	ErrInvalidPrice       = NewDomainException("Invalid price")
	ErrInvalidCapacity    = NewDomainException("Invalid capacity")
	ErrRequiredAPIKeyName = NewDomainException("API key name is required")
	ErrInvalidEmail       = NewDomainException("Invalid email")
	ErrEmailAlreadyExists = NewDomainException("Email already exists")
	ErrInvalidPassword    = NewDomainException("Invalid password")
	ErrInvalidPeriod      = NewDomainException("Invalid period")
	ErrInsufficientPeriod = NewDomainException("Overnight should be greater than 1")
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
