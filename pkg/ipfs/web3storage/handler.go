package web3storage

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ipfs/go-cid"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/multiformats/go-multihash"
	"github.com/web3-storage/go-ucanto/core/delegation"
	"github.com/web3-storage/go-ucanto/did"
	"github.com/web3-storage/go-ucanto/principal"
	"github.com/web3-storage/go-ucanto/principal/ed25519/signer"
	"github.com/web3-storage/go-w3up/capability/storeadd"
	"github.com/web3-storage/go-w3up/client"
)

type W3Storage struct {
	issuer principal.Signer
	proof  delegation.Delegation
}

// NewW3Storage initializes a new web3.storage handler.
func NewW3Storage(privateKey []byte, proof []byte) (*W3Storage, error) {
	issuer, err := signer.Decode(privateKey)
	if err != nil {
		return nil, err
	}

	proofDel, err := delegation.Extract(proof)
	if err != nil {
		return nil, err
	}

	return &W3Storage{
		issuer: issuer,
		proof:  proofDel,
	}, nil
}

// UploadFile allows uploading a CAR archive to a web3.storage space.
func (w3s *W3Storage) UploadCar(carFile *os.File, space did.DID) (*cid.Cid, error) {
	data, err := io.ReadAll(carFile)
	if err != nil {
		return nil, err
	}

	mh, err := multihash.Sum(data, multihash.SHA2_256, -1)
	if err != nil {
		return nil, err
	}

	link := cidlink.Link{
		Cid: cid.NewCidV1(0x202, mh),
	}

	rcpt, err := client.StoreAdd(w3s.issuer,
		space,
		&storeadd.Caveat{Link: link, Size: uint64(len(data))},
		client.WithProof(w3s.proof),
	)
	if err != nil {
		return nil, err
	}

	if rcpt.Out().Ok().Status == "upload" {
		req, err := http.NewRequest(http.MethodPut, *rcpt.Out().Ok().Url, bytes.NewReader(data))
		if err != nil {
			return nil, err
		}

		headerVals := map[string][]string{}
		for k, v := range rcpt.Out().Ok().Headers.Values {
			headerVals[k] = []string{v}
		}

		req.Header = headerVals
		req.ContentLength = int64(len(data))

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode < 200 || resp.StatusCode > 200 {
			return nil, fmt.Errorf("could not uploade .car file: %s", resp.Status)
		}
	}

	return &link.Cid, nil
}
