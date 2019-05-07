package server

import (
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)

// KeyReader reads the server private key
type KeyReader interface {
	GetServerKey() (ssh.Signer, error)
}

// FileKeyReader the name of the file that contains the server key
type FileKeyReader string

// GetServerKey reads a file that contains the server key
func (file FileKeyReader) GetServerKey() (ssh.Signer, error) {
	bytes, err := ioutil.ReadFile(string(file))

	if err != nil {
		return nil, err
	}

	key, err := ssh.ParsePrivateKey(bytes)
	if err != nil {
		return nil, err
	}

	return key, nil
}
