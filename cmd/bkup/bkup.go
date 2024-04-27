package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gm0stache/atproto-ipfs-bkup/pkg/atproto"
	"github.com/joho/godotenv"
)

const (
	envBskyHandle string = "BSKY_HANDLE"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	handle := os.Getenv(envBskyHandle)
	if handle == "" {
		log.Fatalf("env var %q must be set", envBskyHandle)
	}

	ctx := context.Background()
	atID, err := atproto.GetATID(ctx, handle)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("handle: %q, PDS URL: %q\n", handle, atID.PDSEndpoint())

	path := atID.DID.String() + ".car"
	if err := atproto.DownloadRepo(ctx, path, atID); err != nil {
		log.Fatal(err)
	}

	// todo: uploade to IPFS

	fmt.Println("done.")
}
