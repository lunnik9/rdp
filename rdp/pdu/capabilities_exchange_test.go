package pdu

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lunnik9/rdp/rdp/mcs"
)

func TestServerDemandActivePDU_Deserialize(t *testing.T) {
	data, err := hex.DecodeString("030001d802f08068000103eb7081c9c9011100ea03ea0301000400b301524450001100000009000800ea03000001001800010003000002000000001d04000000000000010114000c0002000000400600000a0008000600000008000a000100190019001b00060001000e0008000100000002001c00100001000100010000040003000001000100001e010000001d00600004b91b8dca0f004f15589fae2d1a87e2d6000300010103122f777672bd6344afb3b73c9c6f788600040000000000a651439c3535ae42910ccdfce5760b5800040000000000d4cc44278a9d744e803c0ecbeea19c5400040000000000030058000000000000000000000000000000000040420f0001001400000001000000aa000101010101000000010000010000000101010101010101000101010100000000a106060040420f0040420f00010000000000000012000800010000000d00580075030000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000017000800ff00000018000b0002000000030c001a000800a79400001c000c0052000000000000001e0008000000000000000000")
	require.NoError(t, err)

	data = data[7:]

	wire := bytes.NewReader(data)

	var dom mcs.DomainPDU

	require.NoError(t, dom.Deserialize(wire))

	var resp ServerDemandActive

	require.NoError(t, resp.Deserialize(wire))

	_ = data
}

func TestClientConfirmActivePDU_Serialize(t *testing.T) {
	const userID uint16 = 1007

	confirmActive := NewClientConfirmActive(66538, userID, 1280, 1024, false)
	confirmActive.SourceDescriptor = []byte("MSTSC\x00")
	confirmActive.CapabilitySets = []CapabilitySet{
		{
			CapabilitySetType: CapabilitySetTypeGeneral,
			GeneralCapabilitySet: &GeneralCapabilitySet{
				OSMajorType: 1,
				OSMinorType: 3,
				ExtraFlags:  0x041d,
			},
		},
		{
			CapabilitySetType: CapabilitySetTypeBitmap,
			BitmapCapabilitySet: &BitmapCapabilitySet{
				PreferredBitsPerPixel: 0x18,
				Receive1BitPerPixel:   1,
				Receive4BitsPerPixel:  1,
				Receive8BitsPerPixel:  1,
				DesktopWidth:          1280,
				DesktopHeight:         1024,
				DesktopResizeFlag:     1,
			},
		},
		{
			CapabilitySetType: CapabilitySetTypeOrder,
			OrderCapabilitySet: &OrderCapabilitySet{
				OrderFlags: 0x002a,
				OrderSupport: [32]byte{
					0x01, 0x01, 0x01, 0x01, 0x01, 0x00, 0x00, 0x01, 0x01, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
					0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x00, 0x01, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
				textFlags:        0x06a1,
				DesktopSaveSize:  0x38400,
				textANSICodePage: 0x04e4,
			},
		},
		{
			CapabilitySetType: CapabilitySetTypeBitmapCacheRev2,
			BitmapCacheCapabilitySetRev2: &BitmapCacheCapabilitySetRev2{
				CacheFlags:           0x0003,
				NumCellCaches:        3,
				BitmapCache0CellInfo: 0x00000078,
				BitmapCache1CellInfo: 0x00000078,
				BitmapCache2CellInfo: 0x800009fb,
			},
		},
		{
			CapabilitySetType: CapabilitySetTypeColorCache,
			ColorCacheCapabilitySet: &ColorCacheCapabilitySet{
				ColorTableCacheSize: 6,
			},
		},
		{
			CapabilitySetType:             CapabilitySetTypeActivation,
			WindowActivationCapabilitySet: &WindowActivationCapabilitySet{},
		},
		{
			CapabilitySetType:    CapabilitySetTypeControl,
			ControlCapabilitySet: &ControlCapabilitySet{},
		},
		{
			CapabilitySetType: CapabilitySetTypePointer,
			PointerCapabilitySet: &PointerCapabilitySet{
				ColorPointerFlag:      1,
				ColorPointerCacheSize: 20,
				PointerCacheSize:      21,
			},
		},
		{
			CapabilitySetType:  CapabilitySetTypeShare,
			ShareCapabilitySet: &ShareCapabilitySet{},
		},
		{
			CapabilitySetType: CapabilitySetTypeInput,
			InputCapabilitySet: &InputCapabilitySet{
				InputFlags:          0x0015,
				KeyboardLayout:      0x00000409,
				KeyboardType:        4,
				KeyboardFunctionKey: 12,
			},
		},
		{
			CapabilitySetType: CapabilitySetTypeSound,
			SoundCapabilitySet: &SoundCapabilitySet{
				SoundFlags: 0x0001,
			},
		},
		{
			CapabilitySetType: CapabilitySetTypeFont,
			FontCapabilitySet: &FontCapabilitySet{
				fontSupportFlags: 0x0001,
			},
		},
		{
			CapabilitySetType: CapabilitySetTypeGlyphCache,
			GlyphCacheCapabilitySet: &GlyphCacheCapabilitySet{
				GlyphCache: [10]CacheDefinition{
					{
						CacheEntries:         254,
						CacheMaximumCellSize: 4,
					},
					{
						CacheEntries:         254,
						CacheMaximumCellSize: 4,
					},
					{
						CacheEntries:         254,
						CacheMaximumCellSize: 8,
					},
					{
						CacheEntries:         254,
						CacheMaximumCellSize: 8,
					},
					{
						CacheEntries:         254,
						CacheMaximumCellSize: 16,
					},
					{
						CacheEntries:         254,
						CacheMaximumCellSize: 32,
					},
					{
						CacheEntries:         254,
						CacheMaximumCellSize: 64,
					},
					{
						CacheEntries:         254,
						CacheMaximumCellSize: 128,
					},
					{
						CacheEntries:         254,
						CacheMaximumCellSize: 256,
					},
					{
						CacheEntries:         64,
						CacheMaximumCellSize: 256,
					},
				},
				FragCache:         0x1000100,
				GlyphSupportLevel: 3,
			},
		},
		{
			CapabilitySetType: CapabilitySetTypeBrush,
			BrushCapabilitySet: &BrushCapabilitySet{
				BrushSupportLevel: 1,
			},
		},
		{
			CapabilitySetType: CapabilitySetTypeOffscreenBitmapCache,
			OffscreenBitmapCacheCapabilitySet: &OffscreenBitmapCacheCapabilitySet{
				OffscreenSupportLevel: 1,
				OffscreenCacheSize:    7680,
				OffscreenCacheEntries: 100,
			},
		},
		{
			CapabilitySetType: CapabilitySetTypeVirtualChannel,
			VirtualChannelCapabilitySet: &VirtualChannelCapabilitySet{
				Flags: 0x00000001,
			},
		},
		{
			CapabilitySetType: CapabilitySetTypeDrawNineGridCache,
			DrawNineGridCacheCapabilitySet: &DrawNineGridCacheCapabilitySet{
				drawNineGridSupportLevel: 2,
				drawNineGridCacheSize:    2560,
				drawNineGridCacheEntries: 256,
			},
		},
		{
			CapabilitySetType:        CapabilitySetTypeDrawGDIPlus,
			DrawGDIPlusCapabilitySet: &DrawGDIPlusCapabilitySet{},
		},
	}

	req := mcs.DomainPDU{
		Application: mcs.SendDataRequest,
		ClientSendDataRequest: &mcs.ClientSendDataRequest{
			Initiator: userID,
			ChannelId: 1003,
			Data:      confirmActive.Serialize(),
		},
	}

	expected := []byte{
		0x64, 0x00, 0x06, 0x03, 0xeb, 0x70, 0x81, 0xf0, // mcs header
		0xf0, 0x01, 0x13, 0x00, 0xef, 0x03, 0xea, 0x03, 0x01, 0x00, 0xea, 0x03, 0x06, 0x00, 0xda, 0x01,
		0x4d, 0x53, 0x54, 0x53, 0x43, 0x00, 0x12, 0x00, 0x00, 0x00, 0x01, 0x00, 0x18, 0x00, 0x01, 0x00,
		0x03, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x1d, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x02, 0x00, 0x1c, 0x00, 0x18, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x05,
		0x00, 0x04, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x03, 0x00,
		0x58, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x14, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00,
		0x2a, 0x00, 0x01, 0x01, 0x01, 0x01, 0x01, 0x00, 0x00, 0x01, 0x01, 0x01, 0x00, 0x01, 0x00, 0x00,
		0x00, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x00, 0x01, 0x01, 0x01, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xa1, 0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x84, 0x03, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xe4, 0x04, 0x00, 0x00, 0x13, 0x00, 0x28, 0x00, 0x03, 0x00, 0x00, 0x03, 0x78, 0x00,
		0x00, 0x00, 0x78, 0x00, 0x00, 0x00, 0xfb, 0x09, 0x00, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00,
		0x08, 0x00, 0x06, 0x00, 0x00, 0x00, 0x07, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x05, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x02, 0x00, 0x08, 0x00,
		0x0a, 0x00, 0x01, 0x00, 0x14, 0x00, 0x15, 0x00, 0x09, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x0d, 0x00, 0x58, 0x00, 0x15, 0x00, 0x00, 0x00, 0x09, 0x04, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0c, 0x00, 0x08, 0x00, 0x01, 0x00, 0x00, 0x00,
		0x0e, 0x00, 0x08, 0x00, 0x01, 0x00, 0x00, 0x00, 0x10, 0x00, 0x34, 0x00, 0xfe, 0x00, 0x04, 0x00,
		0xfe, 0x00, 0x04, 0x00, 0xfe, 0x00, 0x08, 0x00, 0xfe, 0x00, 0x08, 0x00, 0xfe, 0x00, 0x10, 0x00,
		0xfe, 0x00, 0x20, 0x00, 0xfe, 0x00, 0x40, 0x00, 0xfe, 0x00, 0x80, 0x00, 0xfe, 0x00, 0x00, 0x01,
		0x40, 0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x01, 0x03, 0x00, 0x00, 0x00, 0x0f, 0x00, 0x08, 0x00,
		0x01, 0x00, 0x00, 0x00, 0x11, 0x00, 0x0c, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x1e, 0x64, 0x00,
		0x14, 0x00, 0x0c, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x15, 0x00, 0x0c, 0x00,
		0x02, 0x00, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x01, 0x16, 0x00, 0x28, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
	actual := req.Serialize()

	require.Equal(t, expected, actual)
}
