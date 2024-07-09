package exceptions

type DomainException struct {
	message string
}

var (
	ErrPeriodNotAllowed = NewDomainException("Period not allowed")
	ErrCapacityExceeded = NewDomainException("Capacity exceeded")
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
