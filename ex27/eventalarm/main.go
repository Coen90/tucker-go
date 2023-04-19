package main

import "fmt"

type Event interface {
	Register(EventListener)
}

type EventListener interface {
	OnFire()
}

type Mail struct {
	listener EventListener
}

func (m *Mail) Register(listener EventListener) { // Event 구현체
	m.listener = listener
}

func (m *Mail) OnRecv() {
	m.listener.OnFire()
}

type Alarm struct {
}

type Push struct {

}

func (p *Push) OnFire() { // EventListner 구현체
	fmt.Println("푸시가 왔습니다.")
}

func (a *Alarm) OnFire() { // EventListner 구현체
	fmt.Println("메일이 왔습니다.")
} 

func main() {
	var mail = &Mail{}
	var listener EventListener

	listener = &Alarm{}
	
	mail.Register(listener)
	mail.OnRecv()

	listener = &Push{}
	mail.Register(listener)
	mail.OnRecv()
}