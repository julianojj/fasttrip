package exceptions

type UnauthorizedException struct {
	message string
}

var (
	ErrInvalidLogin = NewUnauthorizedException("invalid login")
)

var _ error = (*UnauthorizedException)(nil)

func NewUnauthorizedException(message string) *UnauthorizedException {
	return &UnauthorizedException{
		message,
	}
}

func (e *UnauthorizedException) Error() string {
	return e.message
}
