package main

import (
	"bytes"
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

type Payload struct {
	Name  string
	Email string
}

func main() {
	payload := Payload{Name: os.Args[1], Email: os.Args[2]}

	fmt.Println(payload)
}

func sign(p Payload) {
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
}

func verifySignature() {

}
