package ipfs

import (
	"bytes"
	"context"
	"io"

	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/client/rpc"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/multiformats/go-multiaddr"
)

// IPFSAPI allows interacting with an IPFS node via it's HTTP API.
type IPFSAPI struct {
	api *rpc.HttpApi
}

// NewIPFSAPI initializes a client for interacting with a local IPFS node.
func NewLocalIPFSAPI() (*IPFSAPI, error) {
	ma, err := multiaddr.NewMultiaddr("/ip4/127.0.0.1/tcp/5001")
	if err != nil {
		return nil, err
	}

	api, err := rpc.NewApi(ma)
	if err != nil {
		return nil, err
	}

	return &IPFSAPI{
		api: api,
	}, nil
}

// Upload stores data on an IPFS node.
func (ipfsAPI *IPFSAPI) Upload(ctx context.Context, data []byte) (string, error) {
	builder := basicnode.Prototype.Any.NewBuilder()
	if err := builder.AssignBytes(data); err != nil {
		return "", err
	}

	// node := builder.Build()
	// ln := ipldlegacy.LegacyNode{
	// 	Block: nil,
	// 	Node:  node,
	// }

	dBuff := bytes.NewBuffer(data)
	stats, err := ipfsAPI.api.Block().Put(ctx, dBuff)
	if err != nil {
		return "", err
	}
	return stats.Path().String(), err

	//	if err := api.Dag().Add(ctx, &ln); err != nil {
	//		return "", err
	//	}
	//
	// return "", nil // todo: return (c)id which can be used to retrieve the stored content.
}

// Get allows retrieving stored data.
func (ipfsAPI *IPFSAPI) Get(ctx context.Context, id string) ([]byte, error) {
	p, err := path.NewPath(id)
	if err != nil {
		return nil, err
	}

	r, err := ipfsAPI.api.Block().Get(ctx, p)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(r)
}
