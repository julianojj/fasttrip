package adapters

import "golang.org/x/crypto/bcrypt"

type Bcrypt struct{}

func NewBcrypt() *Bcrypt {
	return &Bcrypt{}
}

var _ Hash = (*Bcrypt)(nil)

func (b *Bcrypt) EncryptPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (b *Bcrypt) DecryptPassword(encryptedPassword []byte, password string) bool {
	return bcrypt.CompareHashAndPassword(encryptedPassword, []byte(password)) == nil
}
