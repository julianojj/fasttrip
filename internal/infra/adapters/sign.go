package adapters

type Sign interface {
	Encode(sub string, expiresIn int64) (string, error)
	Verify(token string) bool
}
