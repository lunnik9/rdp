package x224

import "github.com/lunnik9/rdp/rdp/tpkt"

type Protocol struct {
	tpktConn *tpkt.Protocol
}

func New(tpktConn *tpkt.Protocol) *Protocol {
	return &Protocol{
		tpktConn: tpktConn,
	}
}
