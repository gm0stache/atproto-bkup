package ipfs_test

import (
	"context"
	"os"
	"testing"

	"github.com/gm0stache/atproto-ipfs-bkup/pkg/ipfs"
	"github.com/stretchr/testify/require"
)

func TestUpload(t *testing.T) {
	// arrange
	file, err := os.ReadFile("../../testdata/did:plc:5molcdxko5rtwkjivyiviss6.car")
	require.NoError(t, err)

	api, err := ipfs.NewLocalIPFSAPI()
	require.NoError(t, err)

	ctx := context.Background()

	// act
	cid, err := api.Upload(ctx, file)

	// assert
	require.NoError(t, err)

	byts, err := api.Get(ctx, cid)
	require.NoError(t, err)
	require.EqualValues(t, file, byts)
}
