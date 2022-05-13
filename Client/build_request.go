package main

import (
	"encoding/binary"
)

func buildRequest(innerBody string, addresses []string, publicKeys []string) (string, []string) {
	keys := make([]string, len(addresses))
	// keys[0] = "hUFHodcVSIhdhlbUwMAsWvufvELjBIec"
	// keys[1] = "BHUMoHQtScWrtjFNNMxbUnNxlaHdQtyQ"
	// keys[2] = "cWknRXFBnkHWWGmzREKTweGtVGEBhHwV"

	request := innerBody
	l := len(addresses)
	for i := range addresses {
		idx := l - i - 1 // Wrap from back to front

		// Get the public keys from the directory
		pub_pem := readStringFromFile(publicKeys[idx])
		pub_parsed, err := ParseRsaPublicKeyFromPemStr(pub_pem)
		if err != nil {
			panic(err)
		}

		// Encrypt
		// key := keys[idx]
		encryptedAddress, key := encryptAES(addresses[idx], "")
		// encryptedAddress, _ := encryptAES(addresses[idx], key)
		encryptedRequest, _ := encryptAES(request, key)
		keys[idx] = key

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
