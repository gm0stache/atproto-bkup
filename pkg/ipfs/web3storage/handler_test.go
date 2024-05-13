package web3storage_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/gm0stache/atproto-ipfs-bkup/pkg/ipfs/web3storage"
	"github.com/web3-storage/go-ucanto/did"
)

func TestUpload(t *testing.T) {
	carFilePath := "/Users/g/Development/projects/atproto-ipfs-bkup/main/testdata/did:plc:5molcdxko5rtwkjivyiviss6.car"
	privateKeyFilePath := "/Users/g/Development/projects/atproto-ipfs-bkup/main/private.key"
	proofFilePath := "/Users/g/Development/projects/atproto-ipfs-bkup/main/proof.ucan"
	spaceDIDStr := "did:key:z6Mkr1NdzBnquYjyaMTLjsUPPThhvupFwziE73uzGiUyWZrA"

	carFile, err := os.Open(carFilePath)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := os.ReadFile(privateKeyFilePath)
	if err != nil {
		log.Fatal(err)
	}

	proof, err := os.ReadFile(proofFilePath)
	if err != nil {
		log.Fatal(err)
	}

	space, err := did.Parse(spaceDIDStr)
	if err != nil {
		log.Fatal(err)
	}

	w3s, err := web3storage.NewW3Storage(privateKey, proof)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v, %+v, %+v", carFile, space, w3s)
	// cid, err := w3s.UploadCar(carFile, space)
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	// fmt.Printf("upload successful.\ncontent ID: %q", cid.String())
}
