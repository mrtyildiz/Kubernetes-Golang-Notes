package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func generateRSAKeyPair(bits int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func saveRSAKeyToFile(filename string, key *rsa.PrivateKey) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(file, privateKeyPEM)
	if err != nil {
		return err
	}
	return nil
}

func saveRSAPublicKeyToFile(filename string, key *rsa.PublicKey) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	publicKeyPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(key),
	}

	err = pem.Encode(file, publicKeyPEM)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	bits := 2048 // You can adjust the key size as needed (e.g., 2048, 4096)

	privateKey, err := generateRSAKeyPair(bits)
	if err != nil {
		panic(err)
	}

	err = saveRSAKeyToFile("private_key.pem", privateKey)
	if err != nil {
		panic(err)
	}

	// Public key can be extracted from the private key
	publicKey := &privateKey.PublicKey
	err = saveRSAPublicKeyToFile("public_key.pem", publicKey)
	if err != nil {
		panic(err)
	}

	fmt.Println("RSA key pair generated and saved successfully.")
}
