package mcs

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/lunnik9/rdp/rdp/asn1"
	"github.com/lunnik9/rdp/rdp/ber"
	"github.com/lunnik9/rdp/rdp/gcc"
)

type ConnectPDUApplication uint8

const (
	connectInitial ConnectPDUApplication = iota + 101
	connectResponse
	connectAdditional
	connectResult
)

type ConnectPDU struct {
	Application           ConnectPDUApplication
	ClientConnectInitial  *ClientConnectInitial
	ServerConnectResponse *ServerConnectResponse
}

func (pdu *ConnectPDU) Serialize() []byte {
	var data []byte

	switch pdu.Application {
	case connectInitial:
		data = pdu.ClientConnectInitial.Serialize()
	}

	buf := new(bytes.Buffer)

	ber.WriteApplicationTag(uint8(pdu.Application), len(data), buf)
	buf.Write(data)

	return buf.Bytes()
}

func (pdu *ConnectPDU) Deserialize(wire io.Reader) error {
	var (
		application uint8
		err         error
	)

	application, err = ber.ReadApplicationTag(wire)
	if err != nil {
		return err
	}
	pdu.Application = ConnectPDUApplication(application)

	_, err = ber.ReadLength(wire)
	if err != nil {
		return err
	}

	switch pdu.Application {
	case connectResponse:
		pdu.ServerConnectResponse = &ServerConnectResponse{}

		return pdu.ServerConnectResponse.Deserialize(wire)
	}

	return fmt.Errorf("%w: application=%v", ErrUnknownConnectApplication, pdu.Application)
}

type ClientConnectInitial struct {
	calledDomainSelector  []byte
	callingDomainSelector []byte
	upwardFlag            bool
	targetParameters      domainParameters
	minimumParameters     domainParameters
	maximumParameters     domainParameters
	userData              *gcc.ConferenceCreateRequest
}

func NewClientMCSConnectInitial(userData []byte) *ClientConnectInitial {
	pdu := ClientConnectInitial{
		calledDomainSelector:  []byte{0x01},
		callingDomainSelector: []byte{0x01},
		upwardFlag:            true,
		targetParameters: domainParameters{
			maxChannelIds:   34,
			maxUserIds:      2,
			maxTokenIds:     0,
			numPriorities:   1,
			minThroughput:   0,
			maxHeight:       1,
			maxMCSPDUsize:   65535,
			protocolVersion: 2,
		},
		minimumParameters: domainParameters{
			maxChannelIds:   1,
			maxUserIds:      1,
			maxTokenIds:     1,
			numPriorities:   1,
			minThroughput:   0,
			maxHeight:       1,
			maxMCSPDUsize:   1056,
			protocolVersion: 2,
		},
		maximumParameters: domainParameters{
			maxChannelIds:   65535,
			maxUserIds:      65535,
			maxTokenIds:     65535,
			numPriorities:   1,
			minThroughput:   0,
			maxHeight:       1,
			maxMCSPDUsize:   65535,
			protocolVersion: 2,
		},
		userData: gcc.NewConferenceCreateRequest(userData),
	}

	return &pdu
}

func (pdu *ClientConnectInitial) Serialize() []byte {
	buf := new(bytes.Buffer)

	ber.WriteOctetString(pdu.calledDomainSelector, buf)
	ber.WriteOctetString(pdu.callingDomainSelector, buf)
	ber.WriteBoolean(pdu.upwardFlag, buf)
	ber.WriteSequence(pdu.targetParameters.Serialize(), buf)
	ber.WriteSequence(pdu.minimumParameters.Serialize(), buf)
	ber.WriteSequence(pdu.maximumParameters.Serialize(), buf)
	ber.WriteOctetString(pdu.userData.Serialize(), buf)

	return buf.Bytes()
}

type ServerConnectResponse struct {
	Result           uint8
	calledConnectId  int
	DomainParameters domainParameters
	UserData         gcc.ConferenceCreateResponse
}

func (pdu *ServerConnectResponse) Deserialize(wire io.Reader) error {
	var err error

	pdu.Result, err = ber.ReadEnumerated(wire)
	if err != nil {
		return err
	}

	pdu.calledConnectId, err = ber.ReadInteger(wire)
	if err != nil {
		return err
	}

	universalTag, err := ber.ReadUniversalTag(asn1.TagSequence, true, wire)
	if err != nil {
		return err
	}

	if !universalTag {
		return errors.New("bad BER tags")
	}

	_, err = ber.ReadLength(wire)
	if err != nil {
		return err
	}

	err = pdu.DomainParameters.Deserialize(wire)
	if err != nil {
		return err
	}

	universalTag, err = ber.ReadUniversalTag(asn1.TagOctetString, false, wire)
	if err != nil {
		return err
	}

	if !universalTag {
		return errors.New("invalid expected BER tag")
	}

	_, err = ber.ReadLength(wire)
	if err != nil {
		return err
	}

	return pdu.UserData.Deserialize(wire)
}

func (p *Protocol) Connect(userData []byte) (io.Reader, error) {
	req := ConnectPDU{
		Application:          connectInitial,
		ClientConnectInitial: NewClientMCSConnectInitial(userData),
	}

	log.Println("MCS: Connect Initial")

	if err := p.x224Conn.Send(req.Serialize()); err != nil {
		return nil, fmt.Errorf("client MCS connect initial request: %w", err)
	}

	log.Println("MCS: Connect Response")

	wire, err := p.x224Conn.Receive()
	if err != nil {
		return nil, err
	}

	var resp ConnectPDU
	if err = resp.Deserialize(wire); err != nil {
		return nil, fmt.Errorf("server MCS connect response: %w", err)
	}

	if resp.ServerConnectResponse.Result != RTSuccessful {
		return nil, fmt.Errorf("unsuccessful MCS connect initial; result=%d", resp.ServerConnectResponse.Result)
	}

	return wire, nil
}
