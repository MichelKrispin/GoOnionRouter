package main

import (
	"encoding/binary"
)

func buildRequest(innerBody string, addresses []string) string {

	request := innerBody
	l := len(addresses)
	for i := range addresses {
		idx := l - i - 1 // Wrap from back to front

		encryptedAddress := ""
		encryptedRequest := ""
		encryptedKey := make([]byte, 512)

		if idx != l-1 {
			// Get the public keys from the directory
			port := addresses[idx][len(addresses[idx])-4:]
			pub_pem := readStringFromFile("keys/public_" + port + ".pem")
			pub_parsed, err := ParseRsaPublicKeyFromPemStr(pub_pem)
			if err != nil {
				panic(err)
			}

			// Encrypt
			key := ""
			encryptedAddress, key = encryptAES(addresses[idx], "")
			encryptedRequest, _ = encryptAES(request, key)

			encryptedKey = []byte(encryptMessage(key, pub_parsed))
		} else { // If last one, then its the servers request
			encryptedAddress = addresses[idx]
			encryptedRequest = request
		}

		addressBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(addressBytes[0:], uint32(len(encryptedAddress)))

		contentBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(contentBytes[0:], uint32(len(encryptedRequest)))

		request = string(addressBytes) + string(contentBytes) + string(encryptedKey) +
			encryptedAddress + encryptedRequest
	}
	return request
}
