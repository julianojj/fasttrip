package adapters

type Hash interface {
	EncryptPassword(password string) ([]byte, error)
	DecryptPassword(encryptedPassword []byte, password string) bool
}
