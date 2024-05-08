package tester_test

import (
	"context"
	"log"
	"os"
	"testing"

	envconfig "github.com/caarlos0/env/v11"
	"github.com/gm0stache/atproto-ipfs-bkup/pkg/ipfs"
	"github.com/gm0stache/atproto-ipfs-bkup/test-integration/testhelper"
	"github.com/ipfs/boxo/path"
	"github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/require"
)

type EnvConfig struct {
	TestIPFSHost string `env:"IPFS_HOST"`
	TestIPFSPort string `env:"IPFS_PORT"`
}

type testConfig struct {
	ipfsNodePath string
}

type testSuit struct {
	t   *testing.T
	cfg *testConfig
}

func TestAll(t *testing.T) {
	cfg := &EnvConfig{}
	if err := envconfig.Parse(cfg); err != nil {
		log.Fatalf("could not read test config from env: %+v\n", err)
	}

	ipfsNodePath, err := testhelper.GetIPFSNodePath()
	if err != nil {
		log.Fatal(err)
	}

	tSuit := testSuit{
		t: t,
		cfg: &testConfig{
			ipfsNodePath: ipfsNodePath,
		},
	}

	testcases := map[string]func(*testSuit){
		"up- and download Blob from IPFS": uploadAndDownload,
	}

	for name, test := range testcases {
		t.Run(name, func(_ *testing.T) {
			test(&tSuit)
		})
	}
}

func uploadAndDownload(ts *testSuit) {
	// arrange
	file, err := os.ReadFile("../../testdata/did:plc:5molcdxko5rtwkjivyiviss6.car")
	require.NoError(ts.t, err)

	ma, err := multiaddr.NewMultiaddr(ts.cfg.ipfsNodePath)
	require.NoError(ts.t, err)

	api, err := ipfs.NewCustomIPFSAPI(ma)
	require.NoError(ts.t, err)

	ctx := context.Background()

	// act
	contentPath, err := api.UploadToPath(ctx, file)

	// assert
	require.NoError(ts.t, err)

	p, err := path.NewPath(contentPath)
	require.NoError(ts.t, err)

	byts, err := api.GetFromPath(ctx, p)
	require.NoError(ts.t, err)
	require.EqualValues(ts.t, file, byts)
}
