package main

import (
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("yoi@123"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", "10.61.200.218:22", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	if err := session.RequestPty("linux", 80, 40, modes); err != nil {
		log.Fatal("request for pseudo terminal failed: ", err)
	}

	//set input and output
	session.Stdout = os.Stdout
	session.Stdin = os.Stdin
	session.Stderr = os.Stderr

	if err := session.Shell(); err != nil {
		log.Fatal("Failed to start shell: ", err)
	}

	err = session.Wait()
	if err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
}
