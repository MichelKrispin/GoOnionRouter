package main

import (
	"encoding/binary"
)

func buildRequest(innerBody string, addresses []string) string {
	request := innerBody
	for _, address := range addresses {
		addressBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(addressBytes[0:], uint32(len(address)))

		contentBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(contentBytes[0:], uint32(len(request)))

		request = string(addressBytes) + string(contentBytes) + address + request
	}
	return request
}
