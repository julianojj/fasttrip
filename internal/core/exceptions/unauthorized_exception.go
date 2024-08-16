package exceptions

type UnauthorizedException struct {
	message string
}

var (
	ErrInvalidLogin = NewUnauthorizedException("invalid login")
	ErrDecodeToken  = NewUnauthorizedException("error to decode token")
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
