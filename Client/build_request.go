package main

import (
	"encoding/binary"
)

func buildRequest(innerBody string, r route) (string, []string) {
	keys := make([]string, len(r.Nodes)-1)

	request := innerBody
	l := len(r.Nodes)

	// Wrap from back to front
	for idx := l - 1; idx > 0; idx-- {
		// Get the public keys from the directory
		pub_pem := r.Nodes[idx-1].PublicKey
		pub_parsed, err := ParseRsaPublicKeyFromPemStr(pub_pem)
		if err != nil {
			panic(err)
		}

		// Encrypt
		encryptedAddress, key := encryptAES(r.Nodes[idx].Address, "")
		encryptedRequest, _ := encryptAES(request, key)
		keys[idx-1] = key

		encryptedKey := []byte(encryptMessage(key, pub_parsed))

		addressBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(addressBytes[0:], uint32(len(encryptedAddress)))

		contentBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(contentBytes[0:], uint32(len(encryptedRequest)))
		request = string(addressBytes) + string(contentBytes) + string(encryptedKey) +
			encryptedAddress + encryptedRequest
	}
	return request, keys
}
