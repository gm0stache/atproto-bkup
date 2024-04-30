package main_test

import (
	"log"
	"os"
	"testing"

	"github.com/gm0stache/atproto-ipfs-bkup/pkg/testhelper"
)

func TestMain(m *testing.M) {
	if _, err := testhelper.GetIPFSNodePath(); err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}
