package ipfs_test

import (
	"context"
	"os"
	"testing"

	"github.com/gm0stache/atproto-ipfs-bkup/pkg/ipfs"
	"github.com/ipfs/boxo/path"
	"github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/require"
)

func TestUpload(t *testing.T) {
	// arrange
	file, err := os.ReadFile("../../testdata/did:plc:5molcdxko5rtwkjivyiviss6.car")
	require.NoError(t, err)

	ma, err := multiaddr.NewMultiaddr("/ip4/127.0.0.1/tcp/5001")
	require.NoError(t, err)

	api, err := ipfs.NewCustomIPFSAPI(ma)
	require.NoError(t, err)

	ctx := context.Background()

	// act
	contentPath, err := api.UploadToPath(ctx, file)

	// assert
	require.NoError(t, err)

	p, err := path.NewPath(contentPath)
	require.NoError(t, err)

	byts, err := api.GetFromPath(ctx, p)
	require.NoError(t, err)
	require.EqualValues(t, file, byts)
}
