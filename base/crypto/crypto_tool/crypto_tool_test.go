package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"testing"
)

func TestFingerprint(t *testing.T) {
	block, _ := pem.Decode([]byte(pkcs1PemPrivateKey))
	if block == nil {
		panic("Invalid pemsKeyData: " + string(pkcs1PemPrivateKey))
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic("Failed to parse private key: " + err.Error())
	}

	// fingerprint uint64 = 12240908862933197005
	// rsa := crypto.NewRSACryptor()
	fmt.Println(computeFingerprint(key))
}
