package server

import (
	"golang.org/x/crypto/ssh"
)

// Auth provides user authorization
type Auth interface {
	CheckPublicKey(user string, publicKey ssh.PublicKey) bool
	CheckPassword(user, password string) bool
}
