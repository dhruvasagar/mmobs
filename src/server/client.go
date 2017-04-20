package main

import (
	"fmt"
	"net"
)

type Client struct {
	Connection net.Conn

	Id int
}

type Action struct {
	Command   string
	Content   string
	Username  string
	IP        string
	Timestamp string
}

var clients []*Client

func (client *Client) Close(doSendMessage bool) {
	if doSendMessage {
		client.SendMessage("disconnect", "")
	}
	client.Connection.Close()
	clients = removeClient(client, clients)
}

func (client *Client) Register() {
	clients = append(clients, client)
	for i, val := range clients {
		if val == client {
			client.Id = i
			break
		}
	}
}

func removeClient(client *Client, arr []*Client) []*Client {
	res := arr
	index := -1
	for i, val := range arr {
		if val == client {
			index = i
			break
		}
	}

	if index >= 0 {
		res := make([]*Client, len(arr)-1)
		copy(res, arr[:index])
		copy(res[index:], arr[index+1:])
	}

	return res
}

func (client *Client) SendMessage(messageType string, message string) {
	message = fmt.Sprintf("/%v", messageType)
	fmt.Fprintln(client.Connection, message)
}
