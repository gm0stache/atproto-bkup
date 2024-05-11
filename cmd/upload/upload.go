package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gm0stache/atproto-ipfs-bkup/pkg/ipfs/web3storage"
	"github.com/web3-storage/go-ucanto/did"
)

func main() {
	carFilePath := flag.String("car", "", ".car file path")
	privateKeyFilePath := flag.String("privatekey", "", "path of the file containing the private key")
	proofFilePath := flag.String("proof", "", "path of the file containing the UCAN proof")
	spaceDIDStr := flag.String("space", "", "DID of the web3.storage space where the file will be uploaded")

	flag.Parse()

	carFile, err := os.Open(*carFilePath)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := os.ReadFile(*privateKeyFilePath)
	if err != nil {
		log.Fatal(err)
	}

	proof, err := os.ReadFile(*proofFilePath)
	if err != nil {
		log.Fatal(err)
	}

	space, err := did.Parse(*spaceDIDStr)
	if err != nil {
		log.Fatal(err)
	}

	w3s, err := web3storage.NewW3Storage(privateKey, proof)
	if err != nil {
		log.Fatal(err)
	}

	cid, err := w3s.UploadCar(carFile, space)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("upload successful.\ncontent ID: %q", cid.String())
}
