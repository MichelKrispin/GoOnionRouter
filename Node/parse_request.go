package main

import (
	"bufio"
	"crypto"
	"crypto/rsa"
	"encoding/binary"
	"log"
	"net"
)

// PROTOCOL
// size of address | size of content | Encrypted 256 AES key | address | content
// 	   4 bytes     |    4 bytes	     \   512 byes            |   ...   \  ...
func parseRequest(c net.Conn, port string) (address string, content string, key string) {
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
	encryptedAddressBytes := make([]byte, int(addressSize))
	for i := 0; i < int(addressSize); i++ {
		b, err := buf.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				log.Fatalln(err)
			}
			break
		}
		encryptedAddressBytes[i] = b
	}
	encryptedAddress := string(encryptedAddressBytes)

	// Read the content
	encryptedContentBytes := make([]byte, int(contentSize))
	for i := 0; i < int(contentSize); i++ {
		b, err := buf.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				log.Fatalln(err)
			}
			break
		}
		encryptedContentBytes[i] = b
	}
	encryptedContent := string(encryptedContentBytes)

	// Decrypt (if port is not empty)
	if port != "" {
		priv_pem := readStringFromFile("keys/private_" + port + ".pem")
		priv_parsed, _ := ParseRsaPrivateKeyFromPemStr(priv_pem)

		decryptedKeyBytes, err := priv_parsed.Decrypt(nil, encryptedKey, &rsa.OAEPOptions{Hash: crypto.SHA256})
		if err != nil {
			panic(err)
		}
		key = string(decryptedKeyBytes)
		log.Print("Decrypted key \"", key, "\"\n")

		address = decryptAES(encryptedAddress, key)
		content = decryptAES(encryptedContent, key)
	} else {
		address = encryptedAddress
		content = encryptedContent
	}

	return address, content, key
}
