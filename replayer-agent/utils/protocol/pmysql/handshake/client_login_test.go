package handshake

import (
	"bytes"
	"testing"

	"github.com/modern-go/parse"
	"github.com/stretchr/testify/require"
)

func TestDecodeClientLoginPacket(t *testing.T) {
	var testCase = []struct {
		rawBytes []byte
		expect   *ClientLogin
		err      error
	}{

		{
			rawBytes: []byte{0xbf, 0x00, 0x00, 0x01, 0x8d, 0xa6, 0xff, 0x01, 0x00, 0x00, 0x00, 0x01,
				0x21, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x72, 0x6f, 0x6f, 0x74, 0x00, 0x14, 0x2f, 0xab, 0x33, 0x55, 0x6a, 0xe7,
				0x75, 0x69, 0xea, 0x29, 0xc6, 0x8b, 0x59, 0x97, 0xaf, 0x1a, 0x3a, 0x04,
				0xaf, 0xf0, 0x6d, 0x79, 0x73, 0x71, 0x6c,
				0x5f, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73,
				0x77, 0x6f, 0x72, 0x64, 0x00, 0x69, 0x03, 0x5f, 0x6f, 0x73, 0x08, 0x6f,
				0x73, 0x78, 0x31, 0x30, 0x2e, 0x31, 0x31, 0x0c, 0x5f, 0x63, 0x6c, 0x69,
				0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x08, 0x6c, 0x69, 0x62,
				0x6d, 0x79, 0x73, 0x71, 0x6c, 0x04, 0x5f, 0x70, 0x69, 0x64, 0x05, 0x38,
				0x37, 0x35, 0x38, 0x38, 0x0f, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
				0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x06, 0x35, 0x2e, 0x37,
				0x2e, 0x31, 0x33, 0x09, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
				0x6d, 0x06, 0x78, 0x38, 0x36, 0x5f, 0x36, 0x34, 0x0c, 0x70, 0x72, 0x6f,
				0x67, 0x72, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x05, 0x6d, 0x79,
				0x73, 0x71, 0x6c},
			expect: &ClientLogin{
				Capabilities:         0xa68d,
				ExtendedCapabilities: 0x01ff,
				MaxPacketSize:        16777216,
				Charset:              0x21,
				UserName:             "root",
				PasswdLen:            20,
				Passwd: []byte{0x2f, 0xab, 0x33, 0x55, 0x6a, 0xe7,
					0x75, 0x69, 0xea, 0x29, 0xc6, 0x8b, 0x59, 0x97, 0xaf, 0x1a, 0x3a, 0x04,
					0xaf, 0xf0},
				Database:   "",
				AuthPlugin: "mysql_native_password",
				Attr: map[string]string{
					"_os":             "osx10.11",
					"_client_name":    "libmysql",
					"_pid":            "87588",
					"_client_version": "5.7.13",
					"_platform":       "x86_64",
					"program_name":    "mysql",
				},
			},
			err: nil,
		},
		{
			rawBytes: []byte{0xbf, 0x00, 0x00, 0x01, 0x8d, 0xa6, 0xff, 0x01, 0x00, 0x00, 0x00, 0x01,
				0x21, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00 /*insert a non-zero byte into consecutive zero bytes*/, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x72, 0x6f, 0x6f, 0x74, 0x00, 0x14, 0x2f, 0xab, 0x33, 0x55, 0x6a, 0xe7,
				0x75, 0x69, 0xea, 0x29, 0xc6, 0x8b, 0x59, 0x97, 0xaf, 0x1a, 0x3a, 0x04,
				0xaf, 0xf0, 0x6d, 0x79, 0x73, 0x71, 0x6c,
				0x5f, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73,
				0x77, 0x6f, 0x72, 0x64, 0x00, 0x69, 0x03, 0x5f, 0x6f, 0x73, 0x08, 0x6f,
				0x73, 0x78, 0x31, 0x30, 0x2e, 0x31, 0x31, 0x0c, 0x5f, 0x63, 0x6c, 0x69,
				0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x08, 0x6c, 0x69, 0x62,
				0x6d, 0x79, 0x73, 0x71, 0x6c, 0x04, 0x5f, 0x70, 0x69, 0x64, 0x05, 0x38,
				0x37, 0x35, 0x38, 0x38, 0x0f, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
				0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x06, 0x35, 0x2e, 0x37,
				0x2e, 0x31, 0x33, 0x09, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
				0x6d, 0x06, 0x78, 0x38, 0x36, 0x5f, 0x36, 0x34, 0x0c, 0x70, 0x72, 0x6f,
				0x67, 0x72, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x05, 0x6d, 0x79,
				0x73, 0x71, 0x6c},
			expect: nil,
			err:    errNotHandshakeResponsePacket,
		},
		{
			rawBytes: []byte{0xbf, 0x00, 0x00, 0x01, 0x8d, 0xa6, 0xff, 0x01, 0x00, 0x00, 0x00, 0x01,
				0x21, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x72, 0x6f, 0x6f, 0x74, 0x00, 0x14, 0x2f, 0xab, 0x33, 0x55, 0x6a, 0xe7,
				0x75, 0x69, 0xea, 0x29, 0xc6, 0x8b, 0x59, 0x97, 0xaf, 0x1a, 0x3a, 0x04,
				0xaf, 0xf0, 0x74, 0x65, 0x73, 0x74, 0x00, 0x6d, 0x79, 0x73, 0x71, 0x6c,
				0x5f, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73,
				0x77, 0x6f, 0x72, 0x64, 0x00, 0x69, 0x03, 0x5f, 0x6f, 0x73, 0x08, 0x6f,
				0x73, 0x78, 0x31, 0x30, 0x2e, 0x31, 0x31, 0x0c, 0x5f, 0x63, 0x6c, 0x69,
				0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x08, 0x6c, 0x69, 0x62,
				0x6d, 0x79, 0x73, 0x71, 0x6c, 0x04, 0x5f, 0x70, 0x69, 0x64, 0x05, 0x38,
				0x37, 0x35, 0x38, 0x38, 0x0f, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
				0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x06, 0x35, 0x2e, 0x37,
				0x2e, 0x31, 0x33, 0x09, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
				0x6d, 0x06, 0x78, 0x38, 0x36, 0x5f, 0x36, 0x34, 0x0c, 0x70, 0x72, 0x6f,
				0x67, 0x72, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x05, 0x6d, 0x79,
				0x73, 0x71, 0x6c},
			expect: &ClientLogin{
				Capabilities:         0xa68d,
				ExtendedCapabilities: 0x01ff,
				MaxPacketSize:        16777216,
				Charset:              0x21,
				UserName:             "root",
				PasswdLen:            20,
				Passwd: []byte{0x2f, 0xab, 0x33, 0x55, 0x6a, 0xe7,
					0x75, 0x69, 0xea, 0x29, 0xc6, 0x8b, 0x59, 0x97, 0xaf, 0x1a, 0x3a, 0x04,
					0xaf, 0xf0},
				Database:   "test",
				AuthPlugin: "mysql_native_password",
				Attr: map[string]string{
					"_os":             "osx10.11",
					"_client_name":    "libmysql",
					"_pid":            "87588",
					"_client_version": "5.7.13",
					"_platform":       "x86_64",
					"program_name":    "mysql",
				},
			},
			err: nil,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		src, err := parse.NewSource(bytes.NewBuffer(tc.rawBytes), 20)
		should.NoError(err)
		actual, err := DecodeClientLoginPacket(src)
		should.Equal(tc.err, err, "case #%d fail", idx)
		if tc.expect == nil {
			should.Nil(actual, "case #%d expect nil but actual isn't a nil", idx)
		} else {
			should.Equal(tc.expect.String(), actual.String(), "case #%d fail", idx)
		}
	}
}
