# Ping-Pong SSH server

This is a simple SSH server. When you connect to the server the terminal shows a "ping$" prompt. When you type something in the terminal the server responds with the same string and the preffix "pong" 

![SSH Client](docs/ssh_client.png)

## Security
The users directory contains the authorized users to connect to the server. The server uses a simple logic to look for authorized users. It looks for a file with the name of the user doing login. If the file ends with .pub, the server interprets the file as a SSH public key. On the contrary if the file ends with .psw the server interprets this file as the hashed password of the user. The users directory contains only one user: hecof, the password for hecof is restlessNeuron!

## Run
To run the server you must have Go installed in your computer. Go to the pingpong directory and type:
```bash
go run main.go
```
The server runs on port 222

## Connect
Once the server is running you can connect to it using you favorite shell.
```bash
ssh hecof@localhost -p 222
```
Use restlessNeuron! as the password

If you prefer to use the SSH Key 
```bash
ssh hecof@localhost -p 222 -i hecof_rsa
```