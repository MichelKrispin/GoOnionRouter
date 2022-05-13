package main

import "encoding/binary"

func buildResponse(response string, key string) string {
	// All of these address and keys in front of the
	// content will be ignored anyway.
	// But the size is still important for parsing.
	dummyAddress := "none"
	addressBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(addressBytes[0:], uint32(len(dummyAddress)))

	// Encrypt the content with the AES key of the client
	response, _ = encryptAES(response, key)
	contentBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(contentBytes[0:], uint32(len(response)))

	dummyKeyBytes := make([]byte, 512)

	response = string(addressBytes) + string(contentBytes) + string(dummyKeyBytes) +
		dummyAddress + response

	return response
}
