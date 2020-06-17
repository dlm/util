package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"
)

func exitOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	exitOnErr(err)

	publicKey := privateKey.PublicKey

	fmt.Println("private: ", privateKey.D)
	fmt.Println("public_x:", publicKey.X)
	fmt.Println("public_y:", publicKey.Y)
}
