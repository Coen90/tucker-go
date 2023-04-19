package main

import "fmt"

// multi interface test

type Closer interface {
	Close()
}

type Opener interface {
	Open()
}

type Multi interface {
	Close()
	Open()
}

type Processor interface {
	Proceed()
}

type PartTimer struct {

}

func (p *PartTimer) Close() {}
func (p *PartTimer) Open() {}
func (p *PartTimer) Proceed() {}

func main() {
	partTimer := &PartTimer{}
	var closer Multi
	closer = partTimer
	fmt.Println(closer)
}