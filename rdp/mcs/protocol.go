package mcs

import "github.com/lunnik9/rdp-html5/rdp/x224"

type Protocol struct {
	x224Conn *x224.Protocol
}

func New(x224Conn *x224.Protocol) *Protocol {
	return &Protocol{
		x224Conn: x224Conn,
	}
}
