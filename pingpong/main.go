package main

import (
	"fmt"
	"io"
	"log"

	"github.com/hecof/ssh-ping-pong/auth"
	"github.com/hecof/ssh-ping-pong/console"
	"github.com/hecof/ssh-ping-pong/server"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	sshServer, err := server.NewSSHServer(server.FileKeyReader("../keys/id_rsa"), auth.NewUserSecretsAuth(auth.LocalUserReader("../users")))
	if err != nil {
		log.Fatal(err)
	}
	sshServer.HandleTerminal(playPingPong)
	sshServer.StartListening("0.0.0.0:222")
}

func playPingPong(conn *ssh.ServerConn, terminal *terminal.Terminal) {
	terminal.Write([]byte(console.GreenLn("Hello " + conn.User() + " lets play ping pong")))
	terminal.SetPrompt("ping$ ")
	for {
		line, err := terminal.ReadLine()
		if err == io.EOF {
			terminal.Write([]byte(console.RedLn("Bye bye " + conn.User())))
			break
		}
		fmt.Println(conn.User() + " says " + line)
		terminal.Write([]byte(console.Yellow("pong> " + console.CyanLn(line))))
	}
}
