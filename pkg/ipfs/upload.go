package ipfs

import (
	"bytes"
	"context"
	"io"

	"github.com/ipfs/boxo/path"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/multiformats/go-multiaddr"
)

// IPFS allows interacting with an IPFS node via it's HTTP API.
type IPFS struct {
	api *rpc.HttpApi
}

// NewIPFSAPI initializes a client for interacting with a local IPFS node.
func NewCustomIPFSAPI(ma multiaddr.Multiaddr) (*IPFS, error) {
	api, err := rpc.NewApi(ma)
	if err != nil {
		return nil, err
	}

	return &IPFS{
		api: api,
	}, nil
}

// UploadToPath stores data on an IPFS node. Returns the IPFS path of the data on success.
func (ipfs *IPFS) UploadToPath(ctx context.Context, data []byte) (string, error) {
	dBuff := bytes.NewBuffer(data)
	stats, err := ipfs.api.Block().Put(ctx, dBuff)
	if err != nil {
		return "", err
	}

	return stats.Path().String(), err
}

// Get allows retrieving stored data.
func (ipfs *IPFS) GetFromPath(ctx context.Context, p path.Path) ([]byte, error) {
	r, err := ipfs.api.Block().Get(ctx, p)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(r)
}
