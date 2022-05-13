package main

import (
	"crypto/aes"
	"crypto/cipher"
	cryptorand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func GenerateRsaKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
	privkey, _ := rsa.GenerateKey(cryptorand.Reader, 4096)
	return privkey, &privkey.PublicKey
}

func ExportRsaPrivateKeyAsPemStr(privkey *rsa.PrivateKey) string {
	privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
	privkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privkey_bytes,
		},
	)
	return string(privkey_pem)
}

func ParseRsaPrivateKeyFromPemStr(privPEM string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

func ExportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) (string, error) {
	pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return "", err
	}
	pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey_bytes,
		},
	)

	return string(pubkey_pem), nil
}

func ParseRsaPublicKeyFromPemStr(pubPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break // fall through
	}
	return nil, errors.New("Key type is not RSA")
}

func writeStringToFile(filename string, data string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(data)
}

func readStringFromFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func generateAndSaveRsaKeyPair(suffix string) {
	// Create the keys
	priv, pub := GenerateRsaKeyPair()

	// Export the keys to pem string
	priv_pem := ExportRsaPrivateKeyAsPemStr(priv)
	pub_pem, _ := ExportRsaPublicKeyAsPemStr(pub)

	// Write them to files
	priv_name := "private"
	pub_name := "public"
	if suffix != "" {
		priv_name += "_" + suffix
		pub_name += "_" + suffix
	}

	writeStringToFile(priv_name+".pem", priv_pem)
	writeStringToFile(pub_name+".pem", pub_pem)
}

func encryptMessage(msg string, publicKey *rsa.PublicKey) string {
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		cryptorand.Reader,
		publicKey,
		[]byte(msg),
		nil)
	if err != nil {
		panic(err)
	}
	return string(encryptedBytes)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func random256Key() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 32)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func encryptAES(message string, key string) (string, string) {
	if key == "" {
		key = random256Key()
	}
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(cryptorand.Reader, nonce); err != nil {
		fmt.Println(err)
	}
	return string(gcm.Seal(nonce, nonce, []byte(message), nil)), key
}

func decryptAES(encryptedMessage string, key string) string {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}

	nonceSize := gcm.NonceSize()
	if len(encryptedMessage) < nonceSize {
		panic(err)
	}

	nonce, ciphertext := encryptedMessage[:nonceSize], encryptedMessage[nonceSize:]
	message, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		panic(err)
	}
	return string(message)
}
