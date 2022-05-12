package main

import (
	"encoding/binary"
)

func buildRequest(innerBody string, addresses []string) string {
	/*
		// TODO: Get the public keys from the directory
		pub_pem := readStringFromFile("key/public_client.pem")
		pub_parsed, err := ParseRsaPublicKeyFromPemStr(pub_pem)
		if err != nil {
			panic(err)
		}
	*/

	request := innerBody
	l := len(addresses)
	for i := range addresses {
		idx := l - i - 1 // Wrap from back to front
		addressBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(addressBytes[0:], uint32(len(addresses[idx])))

		contentBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(contentBytes[0:], uint32(len(request)))
		// request := encryptMessage(request, pub_parsed)

		request = string(addressBytes) + string(contentBytes) + addresses[idx] + request
	}
	return request
}
