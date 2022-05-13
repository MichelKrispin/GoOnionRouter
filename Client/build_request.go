package main

import (
	"encoding/binary"
)

func buildRequest(innerBody string, addresses []string, keys []string) string {

	request := innerBody
	l := len(addresses)
	for i := range addresses {
		idx := l - i - 1 // Wrap from back to front

		// Get the public keys from the directory
		pub_pem := readStringFromFile(keys[idx])
		pub_parsed, err := ParseRsaPublicKeyFromPemStr(pub_pem)
		if err != nil {
			panic(err)
		}

		// Encrypt
		key := "HZJYmrSFCyOzwopuVuPwhQVuCnnErtuc"
		encryptedKey := []byte(encryptMessage(key, pub_parsed))
		/*
			key := ""
			encryptedAddress, key := encryptAES(addresses[idx], "")
			encryptedRequest := ""
			encryptedRequest, _ = encryptAES(request, key)
		*/

		addressBytes := make([]byte, 4)
		// binary.BigEndian.PutUint32(addressBytes[0:], uint32(len(encryptedAddress)))
		binary.BigEndian.PutUint32(addressBytes[0:], uint32(len(addresses[idx])))

		contentBytes := make([]byte, 4)
		// binary.BigEndian.PutUint32(contentBytes[0:], uint32(len(encryptedRequest)))
		binary.BigEndian.PutUint32(contentBytes[0:], uint32(len(request)))
		request = string(addressBytes) + string(contentBytes) + string(encryptedKey) +
			addresses[idx] + request
		//encryptedAddress + encryptedRequest
	}
	return request
}
