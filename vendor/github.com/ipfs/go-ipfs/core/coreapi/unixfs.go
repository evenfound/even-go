package coreapi

import (
	"context"
	"io"

	coreiface "github.com/ipfs/go-ipfs/core/coreapi/interface"
	coreunix "github.com/ipfs/go-ipfs/core/coreunix"
	uio "github.com/ipfs/go-ipfs/unixfs/io"

	cid "gx/ipfs/QmcZfnkapfECQGcLZaf9B79NRg7cRa9EnZh4LSbkCzwNvY/go-cid"
	ipld "gx/ipfs/Qme5bWv7wtjUNGsK2BNGVUFPKiuxWrsqrtvYwCLRw8YFES/go-ipld-format"
)

type UnixfsAPI CoreAPI

// Add builds a merkledag node from a reader, adds it to the blockstore,
// and returns the key representing that node.
func (api *UnixfsAPI) Add(ctx context.Context, r io.Reader) (coreiface.Path, error) {
	k, err := coreunix.AddWithContext(ctx, api.node, r)
	if err != nil {
		return nil, err
	}
	c, err := cid.Decode(k)
	if err != nil {
		return nil, err
	}
	return ParseCid(c), nil
}

// Cat returns the data contained by an IPFS or IPNS object(s) at path `p`.
func (api *UnixfsAPI) Cat(ctx context.Context, p coreiface.Path) (coreiface.Reader, error) {
	dget := api.node.DAG // TODO: use a session here once routing perf issues are resolved

	dagnode, err := resolveNode(ctx, dget, api.node.Namesys, p)
	if err != nil {
		return nil, err
	}

	r, err := uio.NewDagReader(ctx, dagnode, dget)
	if err == uio.ErrIsDir {
		return nil, coreiface.ErrIsDir
	} else if err != nil {
		return nil, err
	}
	return r, nil
}

// Ls returns the contents of an IPFS or IPNS object(s) at path p, with the format:
// `<link base58 hash> <link size in bytes> <link name>`
func (api *UnixfsAPI) Ls(ctx context.Context, p coreiface.Path) ([]*ipld.Link, error) {
	dagnode, err := api.core().ResolveNode(ctx, p)
	if err != nil {
		return nil, err
	}

	var ndlinks []*ipld.Link
	dir, err := uio.NewDirectoryFromNode(api.node.DAG, dagnode)
	switch err {
	case nil:
		l, err := dir.Links(ctx)
		if err != nil {
			return nil, err
		}
		ndlinks = l
	case uio.ErrNotADir:
		ndlinks = dagnode.Links()
	default:
		return nil, err
	}

	links := make([]*ipld.Link, len(ndlinks))
	for i, l := range ndlinks {
		links[i] = &ipld.Link{Name: l.Name, Size: l.Size, Cid: l.Cid}
	}
	return links, nil
}

func (api *UnixfsAPI) core() coreiface.CoreAPI {
	return (*CoreAPI)(api)
}
