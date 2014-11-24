package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func bash_buildpack_bash() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xac, 0x55,
		0x4d, 0x73, 0xdb, 0x36, 0x10, 0x3d, 0x93, 0xbf, 0x62, 0x87, 0xd1, 0x38,
		0x76, 0x67, 0x68, 0x55, 0x4d, 0x4f, 0x89, 0x95, 0xa9, 0x6d, 0x29, 0xad,
		0xa6, 0xfe, 0xaa, 0x2c, 0x9d, 0x5c, 0x8f, 0x06, 0x22, 0x96, 0x12, 0xc6,
		0x10, 0xc0, 0x02, 0xa0, 0x13, 0x8f, 0xeb, 0xff, 0xde, 0x05, 0x49, 0xf1,
		0xc3, 0x92, 0x9b, 0x4b, 0x4e, 0x12, 0xb0, 0xd8, 0xdd, 0xb7, 0xbb, 0x6f,
		0x1f, 0x43, 0x8e, 0x89, 0x64, 0x06, 0xc1, 0xa2, 0xc4, 0xc4, 0x21, 0x5f,
		0x64, 0xcc, 0xad, 0x9b, 0x93, 0x62, 0x1b, 0x0c, 0xc3, 0x65, 0x2e, 0x24,
		0xcf, 0x58, 0xf2, 0x10, 0x17, 0xff, 0x0e, 0x8f, 0xe0, 0x39, 0x0c, 0xb6,
		0x9e, 0x1c, 0x6d, 0x32, 0x8c, 0xce, 0xbc, 0x01, 0x98, 0x02, 0x96, 0x65,
		0x52, 0x24, 0xcc, 0x09, 0xad, 0x20, 0xb7, 0x42, 0xad, 0x40, 0x28, 0xeb,
		0x98, 0x94, 0xc8, 0xa1, 0x8e, 0x63, 0xa3, 0x30, 0x40, 0x65, 0x73, 0x83,
		0xb1, 0xcf, 0x67, 0xc3, 0xa0, 0x49, 0x61, 0xd1, 0xe5, 0x19, 0x7c, 0x86,
		0x3e, 0xc7, 0xc7, 0xbe, 0xca, 0xa5, 0xec, 0x1a, 0x3d, 0x30, 0xf8, 0x97,
		0x82, 0x72, 0x54, 0xae, 0x6d, 0x4a, 0xf4, 0x26, 0x13, 0x12, 0x5b, 0xb6,
		0xcc, 0xe8, 0x24, 0xa5, 0xab, 0xd8, 0x3d, 0x65, 0x68, 0x1b, 0xc3, 0x4b,
		0xbb, 0xa2, 0x0a, 0xdd, 0xbe, 0x9a, 0x26, 0xa5, 0xa9, 0x81, 0x0d, 0xa9,
		0xd1, 0x1b, 0xf8, 0x5d, 0x38, 0x98, 0x4f, 0x2f, 0xa8, 0x58, 0x0e, 0x3a,
		0xf3, 0x85, 0x32, 0x09, 0x94, 0x7d, 0x23, 0x9c, 0x13, 0x76, 0x1d, 0x35,
		0x61, 0x72, 0x23, 0x87, 0x51, 0x6f, 0x10, 0x55, 0x56, 0xfa, 0xff, 0x4b,
		0x04, 0xbe, 0xa5, 0xf4, 0xef, 0xf9, 0xc3, 0xc7, 0xb8, 0x77, 0xb8, 0x64,
		0x16, 0xfd, 0x05, 0xf4, 0x06, 0x47, 0x2f, 0x3b, 0x4d, 0x91, 0x3a, 0xa1,
		0xd0, 0x8e, 0x99, 0x15, 0xba, 0x62, 0x30, 0xe4, 0x57, 0x83, 0x29, 0x2e,
		0xfa, 0x3d, 0xef, 0x4d, 0x8e, 0x22, 0x85, 0xbb, 0x3b, 0x88, 0x15, 0x44,
		0xbd, 0x32, 0x59, 0x04, 0xf7, 0xf7, 0x9f, 0xc0, 0xad, 0x51, 0x85, 0x41,
		0xb0, 0x22, 0xc8, 0x89, 0xd4, 0x0a, 0xc9, 0x4c, 0xa0, 0x22, 0xfa, 0x69,
		0x45, 0x25, 0xf7, 0x20, 0xe1, 0xbb, 0x77, 0x85, 0xd7, 0x1a, 0x93, 0x07,
		0x9d, 0x3b, 0x88, 0xe3, 0x7f, 0x72, 0x81, 0xae, 0x89, 0x5f, 0x3a, 0xc5,
		0xdd, 0x49, 0xa1, 0xb4, 0xd8, 0xc9, 0x17, 0xc7, 0x1c, 0x33, 0x02, 0x3e,
		0x78, 0x2b, 0x73, 0x2a, 0xc2, 0xc0, 0x6c, 0x20, 0x36, 0x69, 0xd7, 0xd4,
		0x3f, 0x5e, 0xf9, 0x24, 0x9d, 0x59, 0x49, 0x61, 0xdd, 0xbe, 0x41, 0x5d,
		0xd0, 0xfd, 0x5b, 0x34, 0x93, 0x16, 0x62, 0x9f, 0xbc, 0xdb, 0xb7, 0x57,
		0x81, 0x0b, 0xce, 0x95, 0x91, 0xdf, 0xc1, 0x59, 0x3d, 0x6e, 0xfc, 0x96,
		0x11, 0xdb, 0x0a, 0x2e, 0xd3, 0x34, 0xe8, 0xa4, 0x8d, 0x83, 0xd3, 0x9b,
		0x9b, 0xc5, 0x68, 0x32, 0xa5, 0x51, 0x10, 0xd5, 0xb7, 0x55, 0x54, 0xb6,
		0x3f, 0xae, 0x2f, 0xc7, 0x7b, 0x0d, 0xd3, 0xf1, 0x5f, 0xf3, 0xf1, 0xed,
		0x6c, 0x31, 0x19, 0x0d, 0xa3, 0x22, 0x6d, 0xdc, 0x9b, 0x9e, 0x5e, 0x8d,
		0xae, 0x2f, 0x9b, 0x27, 0xb7, 0xb3, 0xd3, 0xf3, 0x3f, 0x87, 0x51, 0x82,
		0x9c, 0x99, 0x78, 0xf0, 0x2b, 0x19, 0x92, 0x8c, 0xda, 0x02, 0x4d, 0xb8,
		0xfe, 0x71, 0xb4, 0x2d, 0x64, 0x1b, 0xde, 0xe3, 0x1d, 0x19, 0x9d, 0x65,
		0x54, 0x76, 0x66, 0xc4, 0x23, 0xb1, 0x7d, 0x85, 0x84, 0x35, 0xb7, 0x68,
		0x36, 0x9a, 0xc6, 0x13, 0xaf, 0xb5, 0x67, 0x97, 0x07, 0x06, 0x4a, 0x2f,
		0x35, 0x7f, 0xa2, 0xb8, 0x6b, 0xfd, 0x55, 0x41, 0x3c, 0xad, 0x2e, 0x3e,
		0x2a, 0xbd, 0x32, 0x9a, 0x76, 0xee, 0x6f, 0x9a, 0x5c, 0x0b, 0x7c, 0x75,
		0x6e, 0xe5, 0xab, 0x6e, 0x12, 0x46, 0xac, 0xa8, 0x10, 0x78, 0x00, 0x73,
		0x8b, 0x69, 0x2e, 0x49, 0x34, 0x88, 0xff, 0x6a, 0x65, 0xa1, 0x0f, 0x29,
		0x32, 0x47, 0x3c, 0x6e, 0x9a, 0x76, 0x4e, 0x0b, 0xb3, 0x38, 0xbf, 0xbe,
		0xba, 0x1a, 0x9f, 0xcf, 0x16, 0xb3, 0xc9, 0xe5, 0xf8, 0x7a, 0x3e, 0x1b,
		0x46, 0x1f, 0x7e, 0x2e, 0x03, 0x14, 0x1d, 0xb7, 0x0e, 0x33, 0x58, 0x52,
		0xdb, 0xbf, 0x32, 0xc3, 0xad, 0x5f, 0x19, 0xca, 0x20, 0x96, 0x42, 0x0a,
		0xf7, 0x54, 0xf3, 0x3b, 0xed, 0xb4, 0x03, 0xd5, 0x63, 0x87, 0xe6, 0x56,
		0xe7, 0x26, 0xc1, 0x9d, 0x27, 0x05, 0xc9, 0x5e, 0xcd, 0xdb, 0xcb, 0x48,
		0x39, 0xf0, 0xd6, 0xe6, 0x9c, 0xcd, 0x27, 0x17, 0xa3, 0x1b, 0x1a, 0xc3,
		0x82, 0xe0, 0x76, 0x22, 0x3b, 0xe1, 0x48, 0x5a, 0xa2, 0x2f, 0xe8, 0x92,
		0xb5, 0xd7, 0xb4, 0x24, 0xb7, 0x8e, 0xa4, 0xa0, 0x8e, 0xe8, 0xb7, 0xc1,
		0x03, 0x68, 0xab, 0xe8, 0xee, 0xb2, 0x96, 0x5e, 0xfe, 0x6d, 0xcd, 0xf8,
		0x8e, 0x87, 0xef, 0x46, 0x30, 0xf9, 0x72, 0x3b, 0x7c, 0xff, 0xee, 0x3d,
		0x18, 0x64, 0xdc, 0x2b, 0x48, 0x25, 0x1e, 0x70, 0x72, 0x72, 0xb2, 0x03,
		0x91, 0x9e, 0xef, 0x28, 0x59, 0xb3, 0x68, 0x5b, 0x21, 0xa8, 0xc0, 0x1e,
		0xb4, 0x37, 0xb5, 0x0d, 0xb6, 0x52, 0xa4, 0xc3, 0x5c, 0xd5, 0x14, 0xe2,
		0xd0, 0x45, 0xd6, 0x5f, 0x0a, 0x45, 0xce, 0xce, 0x8b, 0x6f, 0x8b, 0x11,
		0x47, 0x51, 0xbd, 0xf2, 0xa5, 0x56, 0x35, 0xab, 0x37, 0x3c, 0x7c, 0x5d,
		0xfd, 0x4f, 0x47, 0xf4, 0x2c, 0xd5, 0xa6, 0xa5, 0xa7, 0xc2, 0x77, 0xfd,
		0xb9, 0x71, 0xba, 0xfb, 0xed, 0xfe, 0x25, 0xfa, 0x04, 0x5c, 0xd3, 0xcb,
		0xef, 0xe1, 0xab, 0xbd, 0xde, 0xc4, 0x56, 0xd0, 0x35, 0x08, 0x0e, 0x0e,
		0xe0, 0xcd, 0xb9, 0xb4, 0xde, 0x2c, 0xa9, 0xe1, 0x0f, 0x74, 0xe0, 0x24,
		0x5b, 0xa5, 0x2c, 0xb5, 0x98, 0xd1, 0x1d, 0xd3, 0x1e, 0x66, 0xf4, 0x3a,
		0x68, 0xfd, 0x57, 0x10, 0x4a, 0x48, 0xc8, 0x9b, 0x1e, 0x55, 0x6f, 0xe7,
		0x8a, 0x2d, 0xe9, 0xd7, 0xe9, 0x0a, 0x17, 0xb0, 0x2e, 0x91, 0xf0, 0x1b,
		0x8d, 0x7b, 0xb0, 0xcb, 0xda, 0xea, 0x0b, 0xd7, 0xa6, 0x2d, 0x75, 0xc5,
		0x0b, 0xdc, 0x29, 0xf4, 0x88, 0xe8, 0xdb, 0xb2, 0x5b, 0xe8, 0x3a, 0x2d,
		0x8b, 0xf6, 0xcc, 0xb4, 0x8a, 0xf9, 0x4a, 0x58, 0xa0, 0xb3, 0xe4, 0x74,
		0xda, 0x46, 0x6f, 0x6a, 0xf9, 0x21, 0x91, 0xcb, 0x3e, 0x7f, 0x37, 0x94,
		0xa1, 0x1b, 0xfa, 0x50, 0xfe, 0x3f, 0xc8, 0xcf, 0x1d, 0x6b, 0xff, 0x78,
		0xeb, 0x14, 0xbe, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x2a, 0xf5, 0x07,
		0xcf, 0xe5, 0x08, 0x00, 0x00,
	},
		"bash/buildpack.bash",
	)
}

func bash_cmd_bash() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x9c, 0x54,
		0x41, 0x73, 0x9b, 0x3c, 0x10, 0x3d, 0xc3, 0xaf, 0xd8, 0x4f, 0x51, 0x32,
		0xc9, 0x81, 0xe1, 0xc3, 0xa7, 0x0e, 0x1e, 0x77, 0x9c, 0x69, 0x7b, 0x6b,
		0x7b, 0xc9, 0xd1, 0x38, 0x33, 0x2a, 0x88, 0xc0, 0x58, 0x16, 0x1e, 0x84,
		0xdd, 0x74, 0x30, 0xff, 0xbd, 0xbb, 0x42, 0x60, 0x88, 0x7d, 0x6a, 0x2e,
		0x46, 0xbb, 0xab, 0xdd, 0xf7, 0x9e, 0xde, 0xc6, 0xcf, 0x64, 0xaa, 0x44,
		0x2d, 0x21, 0x78, 0x86, 0x2f, 0x3f, 0xbe, 0xbe, 0xf8, 0x7e, 0xba, 0xcf,
		0x02, 0x55, 0x9a, 0xe6, 0xf1, 0x09, 0x5a, 0xdf, 0x1b, 0xd2, 0x99, 0x34,
		0xe9, 0x8a, 0x7d, 0xc7, 0xb8, 0x01, 0x71, 0x12, 0xa5, 0x12, 0xbf, 0x94,
		0x84, 0xb4, 0xda, 0xef, 0x85, 0xce, 0x0c, 0xbb, 0x14, 0x6a, 0xb3, 0x62,
		0x3c, 0xc2, 0xc0, 0xd0, 0x27, 0xd8, 0xc9, 0x3f, 0x06, 0x18, 0xd7, 0x86,
		0xc1, 0x19, 0x8c, 0xcc, 0x80, 0x99, 0x10, 0x4f, 0x71, 0x18, 0x32, 0xbf,
		0xbb, 0xcc, 0xb3, 0x75, 0xf3, 0xa1, 0x63, 0xaf, 0xbc, 0xaa, 0x61, 0x07,
		0xa5, 0xc6, 0x36, 0xed, 0x7f, 0x04, 0x73, 0xb3, 0xde, 0x76, 0x6c, 0x09,
		0x59, 0xe5, 0x7b, 0x9e, 0x4c, 0x8b, 0x0a, 0x13, 0x3b, 0x02, 0x51, 0x69,
		0x89, 0x43, 0xde, 0x6a, 0x79, 0x00, 0xf6, 0x4a, 0x43, 0xec, 0xcc, 0xaa,
		0x6e, 0x66, 0x93, 0xb4, 0x9b, 0xf3, 0x2f, 0x6d, 0x83, 0x13, 0xc4, 0x1f,
		0x7a, 0xca, 0xf7, 0x03, 0x9e, 0x6e, 0xe9, 0xf5, 0xcd, 0x66, 0x50, 0x31,
		0xc8, 0x8f, 0x3a, 0x6d, 0xca, 0x4a, 0x83, 0xa0, 0x93, 0xd3, 0x6d, 0x22,
		0x5b, 0xae, 0x2d, 0x55, 0x4c, 0xe3, 0x6f, 0xbb, 0x88, 0x03, 0x1e, 0x75,
		0x98, 0x56, 0x55, 0x2a, 0x94, 0xd5, 0xc1, 0xa9, 0xa0, 0x09, 0x2e, 0x7f,
		0x9c, 0x50, 0x79, 0x9a, 0xe3, 0xcd, 0x35, 0x9b, 0x2a, 0x10, 0x30, 0x78,
		0xf8, 0x0c, 0x61, 0x26, 0x4f, 0xa1, 0x3e, 0x2a, 0x05, 0x0f, 0x0f, 0xbd,
		0xaa, 0xda, 0xd1, 0xf2, 0x3d, 0xcb, 0x9b, 0x9e, 0x27, 0xe6, 0xad, 0x30,
		0xe1, 0x1d, 0x7e, 0x05, 0x61, 0xc7, 0xb6, 0x2b, 0xdb, 0x6b, 0xce, 0x71,
		0x54, 0xee, 0xc3, 0x0b, 0x39, 0xba, 0x7c, 0x81, 0x5d, 0xe5, 0x09, 0x11,
		0x63, 0x90, 0x0a, 0x01, 0x71, 0xcd, 0x14, 0x49, 0x18, 0xa7, 0xdf, 0x04,
		0xeb, 0xac, 0x41, 0xb4, 0x01, 0x1e, 0x41, 0xc2, 0x12, 0xbe, 0x4e, 0x50,
		0x78, 0xdf, 0xeb, 0x9c, 0x73, 0xfa, 0x79, 0xd0, 0x3f, 0xbf, 0x83, 0x18,
		0x59, 0x50, 0xd1, 0x88, 0x69, 0x00, 0x73, 0x11, 0x09, 0x93, 0x4b, 0x30,
		0x45, 0x99, 0x37, 0x30, 0x84, 0xb1, 0x70, 0x16, 0x3f, 0x9f, 0xa1, 0xa9,
		0x8f, 0x72, 0x48, 0x9b, 0x46, 0x34, 0x47, 0xb3, 0xfa, 0xdf, 0xf7, 0xca,
		0x1c, 0x06, 0x55, 0x47, 0xb7, 0x5a, 0x19, 0x5f, 0x39, 0xc6, 0x13, 0x3e,
		0xd3, 0x71, 0x09, 0x4d, 0x21, 0x35, 0x92, 0xe0, 0xed, 0x44, 0x3f, 0xac,
		0x63, 0xdb, 0x0e, 0x6f, 0xaf, 0x49, 0x07, 0x65, 0x70, 0x0a, 0xb5, 0xdd,
		0x6c, 0x30, 0x44, 0x39, 0xd8, 0x6e, 0xc7, 0x8b, 0xee, 0xbd, 0x7e, 0x56,
		0x60, 0x8e, 0x69, 0x31, 0x38, 0x22, 0x06, 0x5b, 0x48, 0x79, 0x87, 0x6c,
		0x41, 0x4f, 0xab, 0x86, 0x2e, 0x04, 0xeb, 0xba, 0x09, 0x7f, 0xcc, 0x75,
		0x40, 0xba, 0xf6, 0x15, 0x4f, 0x74, 0x3f, 0x2f, 0x9d, 0x27, 0x46, 0x6b,
		0x3c, 0x5f, 0xad, 0x6d, 0x6c, 0x0b, 0xd1, 0x55, 0x38, 0x73, 0xee, 0x2b,
		0xd7, 0xc8, 0x59, 0xcb, 0x3b, 0xd4, 0xa5, 0x6e, 0x72, 0x60, 0x00, 0xf7,
		0xc1, 0xe2, 0x93, 0x81, 0x7b, 0x93, 0xa0, 0xcd, 0x1c, 0xa9, 0xd9, 0xf8,
		0x6b, 0x39, 0x7a, 0x38, 0xbd, 0xd7, 0x46, 0x40, 0xef, 0x65, 0x03, 0xbc,
		0x67, 0xe8, 0x13, 0x54, 0xf7, 0xa2, 0x85, 0x54, 0x87, 0x5b, 0x7b, 0xf4,
		0x52, 0x54, 0xbf, 0x0d, 0x50, 0x16, 0x61, 0x22, 0xe0, 0xbd, 0xb0, 0xbb,
		0x44, 0xd0, 0x6f, 0x2d, 0x93, 0xa8, 0xdf, 0xc8, 0x0c, 0xf4, 0x0a, 0x83,
		0x70, 0x14, 0x9a, 0x48, 0x07, 0xf8, 0x37, 0x30, 0x27, 0x92, 0xd6, 0x13,
		0xf4, 0x81, 0x8b, 0x7e, 0x07, 0x4a, 0xa0, 0x02, 0x78, 0xa3, 0x2f, 0x9b,
		0xd8, 0xab, 0xa5, 0x36, 0xe1, 0x3d, 0x31, 0x0b, 0xe9, 0x1f, 0x45, 0x1f,
		0xd4, 0x18, 0x82, 0xf9, 0x19, 0xc2, 0x00, 0x9d, 0x3c, 0xb9, 0x6e, 0xf7,
		0xfb, 0x86, 0x38, 0x0e, 0x89, 0x0e, 0x88, 0x97, 0x5b, 0xdf, 0x68, 0x34,
		0x8f, 0x5b, 0x11, 0xbb, 0xfc, 0x24, 0xd2, 0x64, 0x33, 0x06, 0xb9, 0xac,
		0x2a, 0xfe, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x2e, 0xf9, 0x29,
		0xc7, 0x05, 0x00, 0x00,
	},
		"bash/cmd.bash",
	)
}

func bash_fn_bash() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x8c, 0x51,
		0x41, 0x6e, 0xc2, 0x30, 0x10, 0x3c, 0xc7, 0xaf, 0x18, 0xad, 0x22, 0x01,
		0xaa, 0xa2, 0x08, 0xae, 0x34, 0x3d, 0x56, 0xea, 0x1b, 0x28, 0x07, 0xcb,
		0xac, 0x89, 0xd5, 0xd4, 0x89, 0x6c, 0x03, 0xaa, 0x28, 0x7f, 0xef, 0x1a,
		0xd2, 0x92, 0x43, 0x55, 0x55, 0xb9, 0x64, 0x77, 0x66, 0x67, 0x76, 0xd6,
		0xca, 0xfa, 0x4a, 0x87, 0x7d, 0x9c, 0x2f, 0x70, 0x56, 0xc5, 0x8e, 0x4d,
		0xa7, 0x03, 0x63, 0xc7, 0xd1, 0x34, 0xf4, 0xe2, 0xe3, 0xc0, 0x26, 0x41,
		0xc3, 0x1e, 0xbc, 0x49, 0xae, 0xf7, 0xb3, 0x08, 0x21, 0x1f, 0xde, 0xd9,
		0xa7, 0x48, 0xaa, 0xe8, 0x7a, 0xa3, 0xbb, 0xdc, 0xe9, 0x9c, 0xe7, 0xa6,
		0x9c, 0xa7, 0x8f, 0x81, 0x51, 0x2e, 0xf1, 0x89, 0x7d, 0xe0, 0x01, 0xdf,
		0x6a, 0x63, 0x59, 0x1d, 0x41, 0x53, 0x03, 0x12, 0xa0, 0x65, 0xbd, 0x43,
		0xb5, 0x5c, 0xa8, 0x82, 0x4d, 0xdb, 0xa3, 0x62, 0x50, 0x79, 0x1e, 0x05,
		0xeb, 0x1a, 0x35, 0xbd, 0x7a, 0xba, 0x64, 0xa2, 0x3e, 0xbd, 0xa1, 0x7a,
		0x6e, 0x30, 0xab, 0x9b, 0xfa, 0x3c, 0x04, 0xe7, 0x13, 0xe8, 0x91, 0xca,
		0x25, 0x3d, 0xd1, 0x65, 0x26, 0x78, 0x0a, 0xc8, 0x5c, 0xc8, 0xa7, 0x2e,
		0x2a, 0xa7, 0xca, 0x16, 0xff, 0x4e, 0x95, 0xa1, 0xe0, 0x86, 0x5c, 0x51,
		0x1e, 0xc8, 0x44, 0xf9, 0xe1, 0xa3, 0xe4, 0xa3, 0x5f, 0x82, 0x45, 0x33,
		0x59, 0x9e, 0xd6, 0xb8, 0x6e, 0x5f, 0xe6, 0xfe, 0xe8, 0xee, 0xbc, 0xed,
		0xff, 0x70, 0x8f, 0x13, 0x7b, 0xba, 0x73, 0xac, 0x6f, 0x72, 0x26, 0xc4,
		0xb6, 0x3f, 0xc5, 0xfe, 0x10, 0x0c, 0x4b, 0xbd, 0xa2, 0xf1, 0x3a, 0x54,
		0x5a, 0x8f, 0x72, 0x3e, 0xbe, 0x18, 0xa4, 0x5a, 0xfc, 0x40, 0xb8, 0x01,
		0xd7, 0xcd, 0x26, 0x80, 0x2a, 0x9c, 0xc5, 0x66, 0x23, 0xa3, 0x77, 0x49,
		0xc2, 0x76, 0xbb, 0x46, 0x6a, 0xd9, 0xab, 0xa2, 0xb8, 0x25, 0x13, 0x5d,
		0x39, 0xa1, 0x76, 0x1d, 0x2a, 0x8f, 0x87, 0x95, 0xf4, 0x6f, 0xc3, 0xd6,
		0x49, 0x9c, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xaf, 0x12, 0xe0,
		0x23, 0x02, 0x00, 0x00,
	},
		"bash/fn.bash",
	)
}

func bash_herokuish_bash() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x74, 0x54,
		0xef, 0x4f, 0xdb, 0x30, 0x10, 0xfd, 0x1c, 0xff, 0x15, 0x87, 0x55, 0xd1,
		0x0d, 0x29, 0x43, 0x9b, 0xf8, 0x04, 0xca, 0x44, 0xe9, 0xba, 0x81, 0xb6,
		0x41, 0xc5, 0x2f, 0x69, 0x02, 0x84, 0x4c, 0x7c, 0x4d, 0x2d, 0x5c, 0xdb,
		0xb2, 0x1d, 0x28, 0x42, 0xfc, 0xef, 0xb3, 0xd3, 0xe6, 0x47, 0x5b, 0xc8,
		0xa7, 0xe4, 0xde, 0xbb, 0xf3, 0xbb, 0xe7, 0xbb, 0x10, 0x31, 0x81, 0x9b,
		0x1b, 0xa0, 0xbd, 0xa3, 0xc1, 0xc5, 0xf1, 0xfd, 0xf5, 0xe8, 0xfc, 0xe2,
		0xe4, 0xf4, 0xe7, 0x19, 0x85, 0x54, 0x7a, 0xa0, 0x7b, 0x14, 0xee, 0xee,
		0x0e, 0xc0, 0x4f, 0x51, 0x91, 0x04, 0xf3, 0xa9, 0x06, 0xba, 0xb5, 0x05,
		0xff, 0x74, 0x69, 0xc1, 0xbd, 0x38, 0x8f, 0x33, 0x38, 0x62, 0x6e, 0x0a,
		0xc2, 0x81, 0x2e, 0x3d, 0xe8, 0x09, 0x70, 0xe6, 0x71, 0x1f, 0x3a, 0xb5,
		0xce, 0x4e, 0x69, 0x27, 0x73, 0x2c, 0x91, 0x39, 0x84, 0xd2, 0x14, 0x96,
		0x71, 0x04, 0xaf, 0x17, 0xf9, 0x7b, 0xa0, 0x2d, 0x14, 0x16, 0x43, 0xb2,
		0xfd, 0x12, 0xf9, 0x73, 0xe1, 0xe1, 0x1b, 0x99, 0x08, 0x42, 0x42, 0x90,
		0x6b, 0x25, 0x5f, 0x80, 0x19, 0x73, 0x6f, 0x98, 0x9f, 0x66, 0xb4, 0xf7,
		0x3a, 0x18, 0x8f, 0xef, 0xc7, 0x83, 0xcb, 0xe3, 0xfd, 0x74, 0x37, 0x84,
		0xdf, 0x68, 0xcb, 0x42, 0xf5, 0xd4, 0xb0, 0x46, 0xa7, 0xd7, 0x35, 0xcb,
		0xcf, 0xcc, 0x6e, 0x80, 0xba, 0xcc, 0x87, 0x52, 0x48, 0xde, 0x70, 0x8f,
		0xae, 0x4e, 0xfe, 0xfc, 0xe8, 0xb2, 0x2b, 0xb8, 0xcb, 0xcf, 0x59, 0x3e,
		0xc5, 0x86, 0x3f, 0x1c, 0x0c, 0x8f, 0x47, 0x5d, 0x7e, 0x05, 0x6f, 0xd4,
		0x37, 0x2c, 0x7f, 0x5c, 0x3d, 0x63, 0x3c, 0x18, 0xfe, 0xde, 0x38, 0x27,
		0xd2, 0x5c, 0x48, 0x26, 0xa8, 0x5c, 0x69, 0x31, 0x8d, 0x19, 0xee, 0xd3,
		0x67, 0x78, 0x25, 0xc9, 0xec, 0x91, 0x0b, 0x0b, 0xa9, 0x81, 0x5b, 0x92,
		0x24, 0xb4, 0x57, 0xbb, 0x40, 0x97, 0xdf, 0x75, 0xbf, 0xf5, 0x77, 0xdb,
		0x55, 0x1d, 0x69, 0x75, 0xaf, 0x70, 0x1a, 0x65, 0x94, 0xbc, 0x11, 0xd2,
		0x39, 0x90, 0x63, 0x2e, 0x99, 0x45, 0xe0, 0xe8, 0xf2, 0x8c, 0x5e, 0x4c,
		0xf5, 0xb3, 0x83, 0x08, 0x83, 0x43, 0xef, 0x85, 0x2a, 0x5c, 0x73, 0x9f,
		0xf5, 0x2d, 0x64, 0xad, 0xaa, 0x1a, 0xaa, 0xad, 0xcf, 0x5a, 0x81, 0x35,
		0xd4, 0x3a, 0x9d, 0x75, 0xd5, 0xd6, 0x70, 0x6b, 0x6c, 0xd6, 0x95, 0xbe,
		0x92, 0xdd, 0x78, 0x98, 0xbd, 0xd7, 0xcb, 0x13, 0x5a, 0x27, 0xb4, 0x7a,
		0xaf, 0x1b, 0xaf, 0xb9, 0xde, 0x07, 0x17, 0x7a, 0x82, 0x25, 0x0b, 0x98,
		0xe2, 0xe0, 0x4a, 0x63, 0xb4, 0xf5, 0xc8, 0x9b, 0xa8, 0x50, 0x13, 0x5d,
		0x15, 0xf3, 0xc2, 0x4b, 0x5c, 0x94, 0xaa, 0x04, 0xf4, 0xfa, 0xb7, 0x78,
		0xf3, 0xf5, 0x57, 0x1a, 0x9f, 0xef, 0x7d, 0xe8, 0xed, 0x44, 0x92, 0x50,
		0x1c, 0x95, 0x5f, 0xb0, 0x9e, 0xa7, 0x42, 0x22, 0xc4, 0x29, 0x00, 0x29,
		0x14, 0x1e, 0x00, 0xd7, 0xc1, 0xf4, 0x7a, 0xc7, 0x62, 0x88, 0x42, 0x96,
		0x41, 0x9a, 0xee, 0x74, 0x36, 0x2b, 0x59, 0x29, 0xde, 0xaf, 0x68, 0x21,
		0x8a, 0xd2, 0xe1, 0x3a, 0x08, 0xd5, 0xd3, 0xaf, 0x6b, 0x05, 0x38, 0x6c,
		0x49, 0x12, 0x46, 0x0e, 0xa3, 0x92, 0x52, 0x19, 0x2b, 0x9e, 0x82, 0x82,
		0x02, 0xf9, 0x42, 0x4f, 0xb8, 0xb5, 0x52, 0xf0, 0x42, 0x70, 0x50, 0xfa,
		0x41, 0xf3, 0x17, 0xe8, 0x1d, 0x46, 0xe2, 0x8c, 0x09, 0xd5, 0x10, 0x20,
		0x45, 0x0d, 0x46, 0x18, 0x9c, 0x30, 0x21, 0x0f, 0x16, 0x4a, 0x2f, 0xcf,
		0x07, 0xc3, 0x51, 0xdc, 0x7e, 0xd8, 0xde, 0x86, 0x8a, 0x33, 0x27, 0x24,
		0xc9, 0x67, 0x3c, 0xc5, 0x79, 0x34, 0xab, 0x9a, 0x09, 0xb7, 0x12, 0x59,
		0x9a, 0xb7, 0x42, 0x4b, 0x95, 0x6b, 0x37, 0x01, 0xe8, 0x55, 0x58, 0xfe,
		0xe8, 0xb8, 0x50, 0xce, 0x33, 0x29, 0x5b, 0x28, 0x4e, 0x55, 0xa7, 0x54,
		0x13, 0x4f, 0xab, 0xb7, 0x0f, 0xb0, 0x65, 0x95, 0x0f, 0x50, 0x29, 0x9c,
		0xdf, 0xd0, 0xe2, 0x64, 0x59, 0x00, 0xfd, 0xcb, 0x14, 0x2b, 0x30, 0xfe,
		0x52, 0xa4, 0xc8, 0x99, 0x8f, 0x37, 0x1e, 0x81, 0x35, 0x11, 0x31, 0x94,
		0x8a, 0x59, 0x7c, 0xdf, 0x8c, 0x17, 0xa8, 0xd0, 0x86, 0xff, 0xd5, 0x26,
		0xb2, 0x78, 0xdf, 0x38, 0xd9, 0x58, 0x9d, 0x4f, 0xe2, 0x6c, 0xd0, 0x31,
		0xb3, 0x4b, 0x1b, 0x70, 0x8e, 0x79, 0xe9, 0x11, 0xc6, 0x4b, 0x6c, 0x4d,
		0x40, 0x9d, 0x92, 0x86, 0x36, 0xd7, 0x34, 0x34, 0x50, 0x2c, 0xf1, 0x3e,
		0x62, 0xe2, 0x31, 0x24, 0x59, 0x80, 0x41, 0x00, 0xa5, 0xe1, 0x5a, 0x0f,
		0xc3, 0x54, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x9c, 0x28, 0x97,
		0xf2, 0x05, 0x00, 0x00,
	},
		"bash/herokuish.bash",
	)
}

func bash_procfile_bash() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x74, 0x91,
		0xcf, 0x6a, 0xc3, 0x30, 0x0c, 0xc6, 0xcf, 0xf1, 0x53, 0x08, 0x13, 0x4a,
		0x0b, 0xcb, 0xf2, 0x00, 0x63, 0x3b, 0xf5, 0x3c, 0x76, 0xef, 0x4a, 0xf1,
		0x1c, 0xb9, 0x31, 0x73, 0xe3, 0x60, 0x2b, 0xa3, 0xa1, 0xcb, 0xbb, 0xcf,
		0xf9, 0xd7, 0x61, 0x9a, 0x9c, 0x22, 0xf3, 0xe9, 0xfb, 0xf4, 0x93, 0xc2,
		0x6a, 0x67, 0xa5, 0xd2, 0x06, 0x33, 0x4f, 0xc2, 0xd1, 0x76, 0x07, 0x37,
		0x96, 0x90, 0x6b, 0x90, 0x75, 0xec, 0x5f, 0xc3, 0x2b, 0xca, 0x15, 0xa9,
		0x16, 0xce, 0xe3, 0x8a, 0x46, 0x6d, 0x8d, 0x7e, 0xd2, 0x34, 0x19, 0x04,
		0xbe, 0xd7, 0x5e, 0xda, 0x1f, 0x74, 0xba, 0x3a, 0x43, 0xdf, 0x86, 0xde,
		0xc3, 0xd0, 0xc5, 0x59, 0xa2, 0x15, 0x1c, 0x0e, 0x90, 0x29, 0xe0, 0xe9,
		0x57, 0xa3, 0x4d, 0x71, 0xaa, 0x05, 0x95, 0xf9, 0xc7, 0x14, 0xc6, 0xe1,
		0x78, 0x7c, 0x01, 0x2a, 0xb1, 0x62, 0x49, 0x62, 0xac, 0x14, 0x66, 0x34,
		0x86, 0xd7, 0xf0, 0x7d, 0xe5, 0xe9, 0x56, 0x0a, 0x82, 0x25, 0x2f, 0xfc,
		0x42, 0x2b, 0x2e, 0x26, 0xfb, 0xc6, 0xd6, 0x87, 0xfa, 0x2a, 0xdc, 0xd9,
		0x03, 0xca, 0xd2, 0xee, 0xc2, 0xd8, 0xa4, 0x2f, 0x80, 0xdf, 0x7b, 0x0b,
		0x94, 0x46, 0x38, 0x9c, 0xb8, 0x20, 0x7b, 0x83, 0xf4, 0x36, 0x94, 0x79,
		0x0e, 0xf9, 0x13, 0x74, 0xbd, 0xc5, 0x21, 0x35, 0x2e, 0x80, 0x28, 0x7d,
		0xc7, 0xf6, 0x31, 0xf6, 0xb3, 0x43, 0x83, 0xc2, 0x2f, 0x61, 0x17, 0xa8,
		0x44, 0x63, 0xe8, 0x34, 0xe3, 0x47, 0xef, 0xa5, 0x35, 0xe6, 0xac, 0x68,
		0x8d, 0xd9, 0x35, 0x5d, 0x71, 0x74, 0x3f, 0x2e, 0x17, 0xd0, 0x78, 0x1a,
		0x4d, 0xe8, 0x89, 0x60, 0xb3, 0x81, 0xcf, 0xa0, 0x4e, 0xbb, 0xef, 0x47,
		0x3d, 0xfe, 0x23, 0xa0, 0xac, 0x83, 0xd4, 0x87, 0xd1, 0x92, 0xb0, 0x38,
		0x55, 0xe2, 0x82, 0xe3, 0x31, 0xa2, 0xb4, 0x95, 0xa3, 0x8c, 0xb9, 0xef,
		0xf6, 0x21, 0xb2, 0xa9, 0x0a, 0xce, 0xba, 0xbf, 0x00, 0x00, 0x00, 0xff,
		0xff, 0x63, 0xf5, 0xc2, 0xec, 0x76, 0x02, 0x00, 0x00,
	},
		"bash/procfile.bash",
	)
}

func bash_slug_bash() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xe2, 0x2a,
		0xce, 0x29, 0x4d, 0xd7, 0xcd, 0xcc, 0x2d, 0xc8, 0x2f, 0x2a, 0xd1, 0xd0,
		0x54, 0xa8, 0xe6, 0xe2, 0x2c, 0x29, 0x2a, 0x4d, 0xe5, 0xaa, 0xe5, 0x82,
		0x48, 0xa4, 0xa7, 0xe6, 0xa5, 0x16, 0x25, 0x96, 0xa4, 0x62, 0x91, 0x4a,
		0xad, 0x40, 0xd7, 0x03, 0x08, 0x00, 0x00, 0xff, 0xff, 0x18, 0x5b, 0x9d,
		0x41, 0x4c, 0x00, 0x00, 0x00,
	},
		"bash/slug.bash",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"bash/buildpack.bash": bash_buildpack_bash,
	"bash/cmd.bash": bash_cmd_bash,
	"bash/fn.bash": bash_fn_bash,
	"bash/herokuish.bash": bash_herokuish_bash,
	"bash/procfile.bash": bash_procfile_bash,
	"bash/slug.bash": bash_slug_bash,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"bash": &_bintree_t{nil, map[string]*_bintree_t{
		"buildpack.bash": &_bintree_t{bash_buildpack_bash, map[string]*_bintree_t{
		}},
		"cmd.bash": &_bintree_t{bash_cmd_bash, map[string]*_bintree_t{
		}},
		"fn.bash": &_bintree_t{bash_fn_bash, map[string]*_bintree_t{
		}},
		"herokuish.bash": &_bintree_t{bash_herokuish_bash, map[string]*_bintree_t{
		}},
		"procfile.bash": &_bintree_t{bash_procfile_bash, map[string]*_bintree_t{
		}},
		"slug.bash": &_bintree_t{bash_slug_bash, map[string]*_bintree_t{
		}},
	}},
}}