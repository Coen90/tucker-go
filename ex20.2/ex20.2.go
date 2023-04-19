package main

import "go_project/ex20.2/fedex"

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	var sender Sender = &fedex.FedexSender{}
	SendBook("어린왕자", sender)
	SendBook("악어때", sender)
} 