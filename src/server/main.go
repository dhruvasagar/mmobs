package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
)

const (
	HOST = "localhost"
	PORT = "3333"
	TYPE = "tcp"
)

var connections []net.Conn

func main() {
	l, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println("Listening on " + HOST + ":" + PORT)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		client := Client{Connection: conn}
		client.Register()

		channel := make(chan string)
		go waitForInput(channel, &client)
		go handleInput(channel, &client)

		client.SendMessage("ready", PORT)
	}
}

func waitForInput(channel chan string, client *Client) {
	defer close(channel)

	reader := bufio.NewReader(client.Connection)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			client.Close(true)
			return
		}
		channel <- string(line)
	}
}

func handleInput(channel <-chan string, client *Client) {
	for {
		message := <-channel
		if message != "" {
			message = strings.TrimSpace(message)
			action, body := getAction(message)

			if action != "" {
				switch action {
				case "message":
					client.SendMessage("message", body)
				case "user":
					client.SendMessage("user", body)
				default:
					client.SendMessage("unrecognized", action)
				}
			}
		}
	}
}

func getAction(message string) (string, string) {
	actionRegex, _ := regexp.Compile(`^\/([^\s]*)\s*(.*)$`)
	res := actionRegex.FindAllStringSubmatch(message, -1)
	if len(res) == 1 {
		return res[0][1], res[0][2]
	}
	return "", ""
}
