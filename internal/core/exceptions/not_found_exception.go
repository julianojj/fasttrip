package exceptions

type NotFoundException struct {
	message string
}

var (
	ErrRoomNotFound    = NewNotFoundException("Room not found")
	ErrBookingNotFound = NewNotFoundException("Booking not found")
)

var _ error = (*NotFoundException)(nil)

func NewNotFoundException(message string) *NotFoundException {
	return &NotFoundException{
		message,
	}
}

func (e *NotFoundException) Error() string {
	return e.message
}
