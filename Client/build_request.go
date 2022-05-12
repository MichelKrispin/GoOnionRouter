package main

import (
	"encoding/binary"
)

func buildRequest(innerBody string, addresses []string) string {
	request := innerBody
	l := len(addresses)
	for i := range addresses {
		idx := l - i - 1 // Wrap from back to front
		addressBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(addressBytes[0:], uint32(len(addresses[idx])))

		contentBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(contentBytes[0:], uint32(len(request)))

		request = string(addressBytes) + string(contentBytes) + addresses[idx] + request
	}
	return request
}
