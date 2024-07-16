package x224

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ConnectionRequest(t *testing.T) {
	req := ConnectionRequest{
		CRCDT:        0xE0, // TPDU_CONNECTION_REQUEST
		DSTREF:       0,
		SRCREF:       0,
		ClassOption:  0,
		VariablePart: nil,
		UserData: []byte{
			0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x3a, 0x20, 0x6d, 0x73, 0x74, 0x73, 0x68, 0x61, 0x73, 0x68,
			0x3d, 0x65, 0x6c, 0x74, 0x6f, 0x6e, 0x73, 0x0d, 0x0a, 0x01, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00,
			0x00,
		},
	}

	actual := req.Serialize()
	expected := []byte{
		0x27, 0xe0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x3a, 0x20, 0x6d,
		0x73, 0x74, 0x73, 0x68, 0x61, 0x73, 0x68, 0x3d, 0x65, 0x6c, 0x74, 0x6f, 0x6e, 0x73, 0x0d, 0x0a,
		0x01, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	require.Equal(t, expected, actual)
}

func Test_ConnectionConfirm(t *testing.T) {
	var actual ConnectionConfirm

	expected := ConnectionConfirm{
		LI:          14,
		CCCDT:       0xd0,
		DSTREF:      0,
		SRCREF:      0x1234,
		ClassOption: 0,
	}

	input := bytes.NewBuffer([]byte{
		0x0e, 0xd0, 0x00, 0x00,
		0x12, 0x34, 0x00, 0x02,
		0x00, 0x08, 0x00, 0x00,
		0x00, 0x00, 0x00,
	})

	require.NoError(t, actual.Deserialize(input))
	require.Equal(t, expected, actual)
}