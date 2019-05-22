package ipfs

import (
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-ipfs/core"
)

/* Recursively un-pin a directory given its hash.
   This will allow it to be garbage collected. */
func UnPinDir(n *core.IpfsNode, rootHash string) error {

	cid, err := cid.Parse([]string{"/ipfs/" + rootHash});
	if err != nil {
		return err
	}

	return n.Pinning.Unpin(n.Context(), cid, true)
}


