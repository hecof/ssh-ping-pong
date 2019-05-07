package server

import (
	"errors"
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

// TerminalHandler user specific terminal stuff
type TerminalHandler func(conn *ssh.ServerConn, terminal *terminal.Terminal)

// SSHServer SSH server instance
type SSHServer struct {
	Key ssh.Signer
	Auth
	terminalHandler TerminalHandler
	config          *ssh.ServerConfig
}

// NewSSHServer creates a new SSHServer
func NewSSHServer(keyReader KeyReader, auth Auth) (*SSHServer, error) {
	key, err := keyReader.GetServerKey()

	if err != nil {
		return nil, err
	}

	return &SSHServer{
		Key:  key,
		Auth: auth,
	}, nil
}

// StartListening start to listen
func (server *SSHServer) StartListening(address string) error {
	socket, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer socket.Close()
	fmt.Printf("Listening on %v\n", socket.Addr())

	server.configSSH()

	for {
		conn, err := socket.Accept()
		if err != nil {
			fmt.Println("Failed to accept incoming connection", err) //@hecof: change this
			break
		}

		go server.handleConn(conn)
	}

	return nil
}

// HandleTerminal assign the handler to manage the terminal
func (server *SSHServer) HandleTerminal(terminalHandler TerminalHandler) {
	server.terminalHandler = terminalHandler
}

func (server *SSHServer) handleConn(conn net.Conn) {
	serverConn, sshChannel, err := server.startSSHSession(conn)
	if err != nil {
		return //@hecof: what to do here?
	}

	channel, requests, err := sshChannel.Accept()
	if err != nil {
		return
	}
	defer channel.Close()

	go server.listenSSHRequests(serverConn, requests)

	if server.terminalHandler != nil {
		term := terminal.NewTerminal(channel, "Connecting...")
		server.terminalHandler(serverConn, term)
	}
}

func (server *SSHServer) configSSH() {
	config := ssh.ServerConfig{
		PublicKeyCallback: func(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
			fmt.Println("Logging " + conn.User() + " with publicKey")

			isPubKeyOk := server.CheckPublicKey(conn.User(), key)
			if isPubKeyOk {
				return nil, nil
			}

			return nil, fmt.Errorf("Public key rejected for user %q", conn.User())
		},
		PasswordCallback: func(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
			fmt.Println("Logging " + conn.User() + " with password")

			isPasswordOk := server.CheckPassword(conn.User(), string(password))
			if isPasswordOk {
				return nil, nil
			}

			return nil, fmt.Errorf("Password rejected for user %q", conn.User())
		},
	}

	config.AddHostKey(server.Key)
	server.config = &config
}

func (server *SSHServer) startSSHSession(conn net.Conn) (*ssh.ServerConn, ssh.NewChannel, error) {
	sshConn, channels, requests, err := ssh.NewServerConn(conn, server.config)

	if err != nil {
		fmt.Println("Failed to handshake. ", err) //@hecof: review
		return nil, nil, err
	}

	fmt.Printf("SSH session established, user: %s\n", sshConn.User())
	go ssh.DiscardRequests(requests)

	for newChannel := range channels {
		if newChannel.ChannelType() != "session" {
			newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
			continue
		}

		return sshConn, newChannel, nil
	}

	return nil, nil, errors.New("no session channel")
}

func (server *SSHServer) listenSSHRequests(conn *ssh.ServerConn, requests <-chan *ssh.Request) {
	//@hecof: right now we are not doing anything
	for req := range requests {
		ok := false

		switch req.Type {
		case "shell":
			ok = true
		case "pty-req":
			ok = true
		case "window-change":
			ok = true
		}

		if req.WantReply {
			req.Reply(ok, nil)
		}
	}

	fmt.Printf("SSH session closed, user: %s\n", conn.User())
}
