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

// func createSelfSignedCertificate() {
	
// 	cert := &x509.Certificate{
// 		SerialNumber: big.NewInt(1658),
// 		Subject: pkix.Name{
// 			Organization:  []string{""},
// 			Country:       []string{""},
// 			Province:      []string{""},
// 			Locality:      []string{""},
// 			StreetAddress: []string{""},
// 			PostalCode:    []string{""},
// 		},
// 		IPAddresses:  []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
// 		NotBefore:    time.Now(),
// 		NotAfter:     time.Now().AddDate(10, 0, 0),
// 		SubjectKeyId: []byte{1, 2, 3, 4, 6},
// 		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
// 		KeyUsage:     x509.KeyUsageDigitalSignature,
// 	}

// 	return cert
// }

// func createSelfSignedCertkey() {
// 	certPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
// 	if err != nil {
// 		return err
// 	}
// }

// func signCertificate() {
// 	certBytes, err := x509.CreateCertificate(rand.Reader, cert, ca, &certPrivKey.PublicKey, caPrivKey)
// 	if err != nil {
// 		return err
// 	}
// }