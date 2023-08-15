package auth

import (
	"errors"

	"github.com/Tesohh/dotdepot/cli/db"
	"golang.org/x/crypto/bcrypt"
)

var ErrTryingToSignupWithNoPW = errors.New("since you're signing up, you need to specify a password in your login.yml")
var ErrWrongPW = errors.New("password invalid")

// If you want to give write access, you need to Verify first
func Verify(store db.Storer[User], creds Credentials) error {
	user, _ := store.Get(db.Query{"username": creds.Username}) // TODO: err != nil when user doesn't exist
	if (*user == User{}) {                                     // this means the user doesn't exist: we need to create it.
		if creds.Password == "" {
			return ErrTryingToSignupWithNoPW
		}
		salted, err := Salt(creds.Password)
		if err != nil {
			return err
		}
		err = store.Put(User{Username: creds.Username, PasswordEncrypted: salted})
		if err != nil {
			return err
		}
		user, err = store.Get(db.Query{"username": creds.Username}) // "refresh" user
		if err != nil {
			return err
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordEncrypted), []byte(creds.Password)); err != nil {
		return ErrWrongPW
	}
	return nil
}
