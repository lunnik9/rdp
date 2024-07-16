package mcs

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lunnik9/rdp-html5/rdp/gcc"
)

// TestClientMCSConnectInitialPDU_Serialize from MS-RDPBCGR protocol examples 4.1.3.
// without TPKT and X224 headers
func TestClientMCSConnectInitialPDU_Serialize(t *testing.T) {
	userData := []byte{
		0x01, 0xc0, 0xd8, 0x00, 0x04, 0x00, 0x08, 0x00, 0x00, 0x05, 0x00, 0x04,
		0x01, 0xca, 0x03, 0xaa, 0x09, 0x04, 0x00, 0x00, 0xce, 0x0e, 0x00, 0x00, 0x45, 0x00, 0x4c, 0x00,
		0x54, 0x00, 0x4f, 0x00, 0x4e, 0x00, 0x53, 0x00, 0x2d, 0x00, 0x44, 0x00, 0x45, 0x00, 0x56, 0x00,
		0x32, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xca, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x18, 0x00, 0x07, 0x00, 0x01, 0x00, 0x36, 0x00, 0x39, 0x00, 0x37, 0x00, 0x31, 0x00, 0x32, 0x00,
		0x2d, 0x00, 0x37, 0x00, 0x38, 0x00, 0x33, 0x00, 0x2d, 0x00, 0x30, 0x00, 0x33, 0x00, 0x35, 0x00,
		0x37, 0x00, 0x39, 0x00, 0x37, 0x00, 0x34, 0x00, 0x2d, 0x00, 0x34, 0x00, 0x32, 0x00, 0x37, 0x00,
		0x31, 0x00, 0x34, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0xc0, 0x0c, 0x00,
		0x0d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xc0, 0x0c, 0x00, 0x1b, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x03, 0xc0, 0x2c, 0x00, 0x03, 0x00, 0x00, 0x00, 0x72, 0x64, 0x70, 0x64,
		0x72, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x80, 0x63, 0x6c, 0x69, 0x70, 0x72, 0x64, 0x72, 0x00,
		0x00, 0x00, 0xa0, 0xc0, 0x72, 0x64, 0x70, 0x73, 0x6e, 0x64, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc0,
	}
	initialPDU := NewClientMCSConnectInitial(userData)
	initialPDU.maximumParameters.maxUserIds = 64535

	req := ConnectPDU{
		Application:          connectInitial,
		ClientConnectInitial: initialPDU,
	}

	expected := []byte{
		0x7f, 0x65, 0x82, 0x01, 0x94, 0x04, 0x01, 0x01, 0x04,
		0x01, 0x01, 0x01, 0x01, 0xff, 0x30, 0x19, 0x02, 0x01, 0x22, 0x02, 0x01, 0x02, 0x02, 0x01, 0x00,
		0x02, 0x01, 0x01, 0x02, 0x01, 0x00, 0x02, 0x01, 0x01, 0x02, 0x02, 0xff, 0xff, 0x02, 0x01, 0x02,
		0x30, 0x19, 0x02, 0x01, 0x01, 0x02, 0x01, 0x01, 0x02, 0x01, 0x01, 0x02, 0x01, 0x01, 0x02, 0x01,
		0x00, 0x02, 0x01, 0x01, 0x02, 0x02, 0x04, 0x20, 0x02, 0x01, 0x02, 0x30, 0x1c, 0x02, 0x02, 0xff,
		0xff, 0x02, 0x02, 0xfc, 0x17, 0x02, 0x02, 0xff, 0xff, 0x02, 0x01, 0x01, 0x02, 0x01, 0x00, 0x02,
		0x01, 0x01, 0x02, 0x02, 0xff, 0xff, 0x02, 0x01, 0x02, 0x04, 0x82, 0x01, 0x33, 0x00, 0x05, 0x00,
		0x14, 0x7c, 0x00, 0x01, 0x81, 0x2a, 0x00, 0x08, 0x00, 0x10, 0x00, 0x01, 0xc0, 0x00, 0x44, 0x75,
		0x63, 0x61, 0x81, 0x1c, 0x01, 0xc0, 0xd8, 0x00, 0x04, 0x00, 0x08, 0x00, 0x00, 0x05, 0x00, 0x04,
		0x01, 0xca, 0x03, 0xaa, 0x09, 0x04, 0x00, 0x00, 0xce, 0x0e, 0x00, 0x00, 0x45, 0x00, 0x4c, 0x00,
		0x54, 0x00, 0x4f, 0x00, 0x4e, 0x00, 0x53, 0x00, 0x2d, 0x00, 0x44, 0x00, 0x45, 0x00, 0x56, 0x00,
		0x32, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xca, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x18, 0x00, 0x07, 0x00, 0x01, 0x00, 0x36, 0x00, 0x39, 0x00, 0x37, 0x00, 0x31, 0x00, 0x32, 0x00,
		0x2d, 0x00, 0x37, 0x00, 0x38, 0x00, 0x33, 0x00, 0x2d, 0x00, 0x30, 0x00, 0x33, 0x00, 0x35, 0x00,
		0x37, 0x00, 0x39, 0x00, 0x37, 0x00, 0x34, 0x00, 0x2d, 0x00, 0x34, 0x00, 0x32, 0x00, 0x37, 0x00,
		0x31, 0x00, 0x34, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0xc0, 0x0c, 0x00,
		0x0d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xc0, 0x0c, 0x00, 0x1b, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x03, 0xc0, 0x2c, 0x00, 0x03, 0x00, 0x00, 0x00, 0x72, 0x64, 0x70, 0x64,
		0x72, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x80, 0x63, 0x6c, 0x69, 0x70, 0x72, 0x64, 0x72, 0x00,
		0x00, 0x00, 0xa0, 0xc0, 0x72, 0x64, 0x70, 0x73, 0x6e, 0x64, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc0,
	}

	actual := req.Serialize()

	require.Equal(t, expected, actual)
}

// TestServerMCSConnectResponsePDU_Deserialize from MS-RDPBCGR protocol examples 4.1.4.
// without TPKT and X224 headers
func TestServerMCSConnectResponsePDU_Deserialize(t *testing.T) {
	var actual ConnectPDU

	expected := ConnectPDU{
		Application: connectResponse,
		ServerConnectResponse: &ServerConnectResponse{
			Result:          0,
			calledConnectId: 0,
			DomainParameters: domainParameters{
				maxChannelIds:   34,
				maxUserIds:      3,
				maxTokenIds:     0,
				numPriorities:   1,
				minThroughput:   0,
				maxHeight:       1,
				maxMCSPDUsize:   65528,
				protocolVersion: 2,
			},
			UserData: gcc.ConferenceCreateResponse{},
		},
	}

	input := bytes.NewBuffer([]byte{
		0x7f, 0x66, 0x82, 0x01, 0x45, 0x0a, 0x01, 0x00, 0x02,
		0x01, 0x00, 0x30, 0x1a, 0x02, 0x01, 0x22, 0x02, 0x01, 0x03, 0x02, 0x01, 0x00, 0x02, 0x01, 0x01,
		0x02, 0x01, 0x00, 0x02, 0x01, 0x01, 0x02, 0x03, 0x00, 0xff, 0xf8, 0x02, 0x01, 0x02, 0x04, 0x82,
		0x01, 0x1f, 0x00, 0x05, 0x00, 0x14, 0x7c, 0x00, 0x01, 0x2a, 0x14, 0x76, 0x0a, 0x01, 0x01, 0x00,
		0x01, 0xc0, 0x00, 0x4d, 0x63, 0x44, 0x6e, 0x81, 0x08, 0x01, 0x0c, 0x0c, 0x00, 0x04, 0x00, 0x08,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x0c, 0x10, 0x00, 0xeb, 0x03, 0x03, 0x00, 0xec, 0x03, 0xed,
		0x03, 0xee, 0x03, 0x00, 0x00, 0x02, 0x0c, 0xec, 0x00, 0x02, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00,
		0x00, 0x20, 0x00, 0x00, 0x00, 0xb8, 0x00, 0x00, 0x00, 0x10, 0x11, 0x77, 0x20, 0x30, 0x61, 0x0a,
		0x12, 0xe4, 0x34, 0xa1, 0x1e, 0xf2, 0xc3, 0x9f, 0x31, 0x7d, 0xa4, 0x5f, 0x01, 0x89, 0x34, 0x96,
		0xe0, 0xff, 0x11, 0x08, 0x69, 0x7f, 0x1a, 0xc3, 0xd2, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00,
		0x00, 0x01, 0x00, 0x00, 0x00, 0x06, 0x00, 0x5c, 0x00, 0x52, 0x53, 0x41, 0x31, 0x48, 0x00, 0x00,
		0x00, 0x00, 0x02, 0x00, 0x00, 0x3f, 0x00, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0xcb, 0x81, 0xfe,
		0xba, 0x6d, 0x61, 0xc3, 0x55, 0x05, 0xd5, 0x5f, 0x2e, 0x87, 0xf8, 0x71, 0x94, 0xd6, 0xf1, 0xa5,
		0xcb, 0xf1, 0x5f, 0x0c, 0x3d, 0xf8, 0x70, 0x02, 0x96, 0xc4, 0xfb, 0x9b, 0xc8, 0x3c, 0x2d, 0x55,
		0xae, 0xe8, 0xff, 0x32, 0x75, 0xea, 0x68, 0x79, 0xe5, 0xa2, 0x01, 0xfd, 0x31, 0xa0, 0xb1, 0x1f,
		0x55, 0xa6, 0x1f, 0xc1, 0xf6, 0xd1, 0x83, 0x88, 0x63, 0x26, 0x56, 0x12, 0xbc, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x48, 0x00, 0xe9, 0xe1, 0xd6, 0x28, 0x46, 0x8b, 0x4e,
		0xf5, 0x0a, 0xdf, 0xfd, 0xee, 0x21, 0x99, 0xac, 0xb4, 0xe1, 0x8f, 0x5f, 0x81, 0x57, 0x82, 0xef,
		0x9d, 0x96, 0x52, 0x63, 0x27, 0x18, 0x29, 0xdb, 0xb3, 0x4a, 0xfd, 0x9a, 0xda, 0x42, 0xad, 0xb5,
		0x69, 0x21, 0x89, 0x0e, 0x1d, 0xc0, 0x4c, 0x1a, 0xa8, 0xaa, 0x71, 0x3e, 0x0f, 0x54, 0xb9, 0x9a,
		0xe4, 0x99, 0x68, 0x3f, 0x6c, 0xd6, 0x76, 0x84, 0x61, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00,
	})

	require.NoError(t, actual.Deserialize(input))
	require.Equal(t, expected, actual)
}
