package bd

import "golang.org/x/crypto/bcrypt"

/* CryptPassword, Encrypt password for user row */
func CryptPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
