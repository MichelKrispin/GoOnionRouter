package main

import (
	"bufio"
	"encoding/binary"
	"log"
	"net"
)

func parseRequest(c net.Conn, keys []string) (address string, content string) {
	buf := bufio.NewReader(c)
	// Read the address size
	addressSizeBytes := []byte{0, 0, 0, 0}
	for i := 0; i < 4; i++ {
		var err error
		addressSizeBytes[i], err = buf.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				log.Fatalln(err)
			}
			break
		}
	}
	addressSize := binary.BigEndian.Uint32(addressSizeBytes)

	// Read the content size
	contentSizeBytes := []byte{0, 0, 0, 0}
	for i := 0; i < 4; i++ {
		var err error
		contentSizeBytes[i], err = buf.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				log.Fatalln(err)
			}
			break
		}
	}
	contentSize := binary.BigEndian.Uint32(contentSizeBytes)

	// Read the encrypted key
	encryptedKey := make([]byte, 512)
	for i := 0; i < 512; i++ {
		var err error
		encryptedKey[i], err = buf.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				log.Fatalln(err)
			}
			break
		}
	}

	// Read the address
	addressBytes := make([]byte, int(addressSize))
	for i := 0; i < int(addressSize); i++ {
		var err error
		addressBytes[i], err = buf.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				log.Fatalln(err)
			}
			break
		}
	}
	address = string(addressBytes)

	// Read the content
	contentBytes := make([]byte, int(contentSize))
	for i := 0; i < int(contentSize); i++ {
		var err error
		contentBytes[i], err = buf.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				log.Fatalln(err)
			}
			break
		}
	}
	content = string(contentBytes)

	// Decrypt the content multiple times using the keys
	for _, key := range keys {
		content = decryptAES(content, key)
	}

	return address, content
}
