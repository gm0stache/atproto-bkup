package tester_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/gm0stache/atproto-ipfs-bkup/pkg/ipfs"
	"github.com/gm0stache/atproto-ipfs-bkup/test-integration/testhelper"
	"github.com/ipfs/boxo/path"
	"github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/require"
)


func TestMain(m *testing.M) {
	if _, err := testhelper.GetIPFSNodePath(); err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}

func TestUploadAndDownload(t *testing.T) {
	// arrange
	file, err := os.ReadFile("../../testdata/did:plc:5molcdxko5rtwkjivyiviss6.car")
	require.NoError(t, err)

	ipfsNodePath, err := testhelper.GetIPFSNodePath()
	require.NoError(t, err)

	ma, err := multiaddr.NewMultiaddr(ipfsNodePath)
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
