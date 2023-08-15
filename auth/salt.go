package auth

import "golang.org/x/crypto/bcrypt"

func Salt(password string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(res), err
}
