package rdp

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"github.com/lunnik9/rdp/rdp/fastpath"
	"github.com/lunnik9/rdp/rdp/mcs"
	"github.com/lunnik9/rdp/rdp/pdu"
	"github.com/lunnik9/rdp/rdp/tpkt"
	"github.com/lunnik9/rdp/rdp/x224"
)

type RemoteApp struct {
	App        string
	WorkingDir string
	Args       string
}

type client struct {
	conn       net.Conn
	buffReader *bufio.Reader
	tpktLayer  *tpkt.Protocol
	x224Layer  *x224.Protocol
	mcsLayer   *mcs.Protocol
	fastPath   *fastpath.Protocol

	domain   string
	username string
	password string

	desktopWidth, desktopHeight uint16

	serverCapabilitySets []pdu.CapabilitySet
	remoteApp            *RemoteApp
	railState            RailState

	selectedProtocol       pdu.NegotiationProtocol
	serverNegotiationFlags pdu.NegotiationResponseFlag
	channels               []string
	channelIDMap           map[string]uint16
	skipChannelJoin        bool
	shareID                uint32
	userID                 uint16
}

const (
	tcpConnectionTimeout = 5 * time.Second
	readBufferSize       = 64 * 1024
)

func NewClient(
	hostname, username, password string,
	desktopWidth, desktopHeight int,
) (*client, error) {
	c := client{
		domain:   "",
		username: username,
		password: password,

		desktopWidth:  uint16(desktopWidth),
		desktopHeight: uint16(desktopHeight),

		selectedProtocol: pdu.NegotiationProtocolSSL,
	}

	var err error

	c.conn, err = net.DialTimeout("tcp", hostname, tcpConnectionTimeout)
	if err != nil {
		return nil, fmt.Errorf("tcp connect: %w", err)
	}

	c.buffReader = bufio.NewReaderSize(c.conn, readBufferSize)

	c.tpktLayer = tpkt.New(&c)
	c.x224Layer = x224.New(c.tpktLayer)
	c.mcsLayer = mcs.New(c.x224Layer)
	c.fastPath = fastpath.New(&c)

	return &c, nil
}
