package main

type Message struct {
	Sender string `json:"sender,omitempty"` 
	Recipient string `json:"recipient,omitempty"`
	Content string `json:"content,omitempty"`
	ServerIP string `json:"server_ip,omitempty"`
	SenderIP string `json:"sender_ip,omitempty"`
}