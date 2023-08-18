package auth

import (
	"errors"

	"github.com/Tesohh/dotdepot/cli/db"
	"golang.org/x/crypto/bcrypt"
)

var ErrTryingToSignupWithNoPW = errors.New("since you're signing up, you need to specify a password in your login.yml")
var ErrWrongPW = errors.New("password invalid")
var ErrNoLoginFile = errors.New("no login file found: please create a login.yml in the dotdepot folder (~/.config/dotdepot)")
var ErrNoUsernameInLoginFile = errors.New("no username in login file: please specify at least a username in your login.yml to get read access")

// If you want to give write access, you need to VerifyWrite first
func VerifyWrite(store db.Storer[User], creds Credentials) error {
	user, err := store.Get(db.Query{"username": creds.Username})
	if err != nil { // this means the user doesn't exist: we need to create it.
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

// This needs to be run BEFORE Verify, to check if the login file exists and all that
func VerifyReadOnly(store db.Storer[User], creds Credentials) error {
	if creds.Username == "" {
		return ErrNoUsernameInLoginFile
	}
	if _, err := store.Get(db.Query{"username": creds.Username}); err != nil {
		return err
	}
	return nil
}
