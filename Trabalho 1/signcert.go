package main

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
)

// Implementar o payload com a chave pública
// 1 - Gerar a chave privada e pegar a chave pública
// 2 - Assinar o certificado
// 3 - Verificar a assinatura

// Self Signed Certificate
// Certificate Authority?
// x.509?

type Payload struct {
	Name  string
	Email string
}

func main() {
	payload := Payload{Name: os.Args[1], Email: os.Args[2]}

	fmt.Println(payload)

	returnedPublicKey, returnedHash, returnedSign := sign(payload)

	signatureVerification := verifySignature(returnedPublicKey, returnedHash, returnedSign)

	fmt.Println(signatureVerification)
}

func sign(p Payload) (rsa.PublicKey, []byte, []byte) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.PublicKey

	payloadToBytes := new(bytes.Buffer)
	json.NewEncoder(payloadToBytes).Encode(p)

	hashedPayload := sha256.New()
	hashedPayload.Write((payloadToBytes.Bytes()))

	hashSum := hashedPayload.Sum(nil)

	fmt.Printf("SHA256 Hash: ", hashSum)

	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, hashSum, nil)
	if err != nil {
		panic(err)
	}

	return publicKey, hashSum, signature
}

func verifySignature(publicKey rsa.PublicKey, hashSum []byte, signature []byte) string {
	err := rsa.VerifyPSS(&publicKey, crypto.SHA256, hashSum, signature, nil)

	if err != nil {
		fmt.Println("Erro ao verificar a assinatura ", err)
	}

	fmt.Println("Assinatura verificada com sucesso")

	return ""
}
