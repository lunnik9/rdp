package mcs

import (
	"io"

	"github.com/lunnik9/rdp-html5/rdp/per"
)

type ServerSendDataIndication struct {
	Initiator uint16
	ChannelId uint16
}

func (d *ServerSendDataIndication) Deserialize(wire io.Reader) error {
	var err error

	d.Initiator, err = per.ReadInteger16(1001, wire)
	if err != nil {
		return err
	}

	d.ChannelId, err = per.ReadInteger16(0, wire)
	if err != nil {
		return err
	}

	_, err = per.ReadEnumerates(wire)
	if err != nil {
		return err
	}

	_, err = per.ReadLength(wire)
	if err != nil {
		return err
	}

	return nil
}

// Receive returns channelName, reader or error
func (p *Protocol) Receive() (uint16, io.Reader, error) {
	wire, err := p.x224Conn.Receive()
	if err != nil {
		return 0, nil, err
	}

	var resp DomainPDU
	if err = resp.Deserialize(wire); err != nil {
		return 0, nil, err
	}

	if resp.Application != SendDataIndication {
		return 0, nil, ErrUnknownDomainApplication
	}

	return resp.ServerSendDataIndication.ChannelId, wire, nil
}
