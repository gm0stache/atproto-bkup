package atproto

import (
	"context"
	"os"

	"github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/atproto/identity"
	"github.com/bluesky-social/indigo/atproto/syntax"
	"github.com/bluesky-social/indigo/xrpc"
)

// GetATID resolves a given handle to an ATproto identity.
func GetATID(ctx context.Context, handle string) (*identity.Identity, error) {
	atID, err := syntax.ParseAtIdentifier(handle)
	if err != nil {
		return nil, err
	}
	idDir := identity.DefaultDirectory()
	ident, err := idDir.Lookup(ctx, *atID)
	if err != nil {
		return nil, err
	}
	return ident, nil
}

// DownloadRepo downloads the PDS content for a given user.
func DownloadRepo(ctx context.Context, path string, id *identity.Identity) error {
	client := xrpc.Client{
		Host: id.PDSEndpoint(),
	}
	repoByts, err := atproto.SyncGetRepo(ctx, &client, id.DID.String(), "")
	if err != nil {
		return err
	}
	return os.WriteFile(path, repoByts, 0666)
}
