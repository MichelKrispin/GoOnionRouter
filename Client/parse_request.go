package main

import (
	"bufio"
	"encoding/binary"
	"log"
	"net"
)

func parseRequest(c net.Conn) (address string, content string) {
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

	// Read the address
	for i := 0; i < int(addressSize); i++ {
		b, err := buf.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				log.Fatalln(err)
			}
			break
		}
		address += string(b)
	}

	// Read the content
	for i := 0; i < int(contentSize); i++ {
		b, err := buf.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				log.Fatalln(err)
			}
			break
		}
		content += string(b)
	}
	return address, content
}
