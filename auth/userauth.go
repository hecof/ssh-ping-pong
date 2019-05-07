package auth

import (
	"crypto/subtle"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh"
)

// UserSecretsReader reads user's secrets
type UserSecretsReader interface {
	GetPublicKey(user string) (ssh.PublicKey, error)
	GetPasswordAndSalt(user string) (string, error)
}

// UserSecretsAuth provides authorization based on user secrets
type UserSecretsAuth struct {
	reader UserSecretsReader
}

// NewUserSecretsAuth craates a new UserAuth based on users secrets
func NewUserSecretsAuth(reader UserSecretsReader) UserSecretsAuth {
	return UserSecretsAuth{reader}
}

// CheckPublicKey checks the public key of the user is correct
func (u UserSecretsAuth) CheckPublicKey(user string, publicKey ssh.PublicKey) bool {
	userPubKey, err := u.reader.GetPublicKey(user)
	if err != nil {
		return false
	}

	userKeyBytes := userPubKey.Marshal()
	publicKeyBytes := publicKey.Marshal()

	return len(publicKeyBytes) == len(userKeyBytes) && subtle.ConstantTimeCompare(publicKeyBytes, userKeyBytes) == 1
}

// CheckPassword checks if the user's password is correct
func (u UserSecretsAuth) CheckPassword(user, password string) bool {
	hashedPsw, err := u.reader.GetPasswordAndSalt(user)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPsw), []byte(password))
	return err == nil
}
