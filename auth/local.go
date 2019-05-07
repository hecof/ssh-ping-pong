package auth

import (
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

// LocalUserReader reads user's secrets stored in a local directory
type LocalUserReader string

// GetPublicKey looks for a public key in a directory 
// The file must be made up of the user parameter plus the .pub suffix
func (reader LocalUserReader) GetPublicKey(user string) (ssh.PublicKey, error) {
	dir := string(reader)
	bytes, err := ioutil.ReadFile(dir + "/" + user + ".pub")
	if err != nil {
		return nil, err
	}

	pubKey, _, _, _, err := ssh.ParseAuthorizedKey(bytes)
	if err != nil {
		return nil, err
	}

	return pubKey, nil
}

// GetPasswordAndSalt looks for a the password of a user in a directory
// The file must be made up of the user parameter plus the .psw suffix 
func (reader LocalUserReader) GetPasswordAndSalt(user string) (string, error) {
	dir := string(reader)
	bytes, err := ioutil.ReadFile(dir + "/" + user + ".psw")
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}