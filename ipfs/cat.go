package ipfs

import (
	"context"
	"github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/interface-go-ipfs-core"
	"github.com/libp2p/go-libp2p-peer"
	"io/ioutil"
	"strings"
	"time"

	"github.com/ipfs/go-ipfs/core"
	unixfs "github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-path"
)

// Fetch data from IPFS given the hash
func Cat(n *core.IpfsNode, path string, timeout time.Duration) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if !strings.HasPrefix(path, "/ipfs/") {
		path = "/ipfs/" + path
	}

	fpath, err := iface.ParsePath(path)
	if err != nil {
		return nil, err
	}

	api, err := unixfs.NewCoreAPI(n);
	if err != nil {
		return nil, err
	}

	r, err := api.Unixfs().Get(ctx, fpath)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(r.(files.File))
}

func ResolveThenCat(n *core.IpfsNode, ipnsPath path.Path, timeout time.Duration, usecache bool) ([]byte, error) {
	var ret []byte
	pid, err := peer.IDB58Decode(ipnsPath.Segments()[0])
	if err != nil {
		return nil, err
	}
	hash, err := Resolve(n, pid, timeout, usecache)
	if err != nil {
		return ret, err
	}
	p := make([]string, len(ipnsPath.Segments()))
	p[0] = hash
	for i := 0; i < len(ipnsPath.Segments())-1; i++ {
		p[i+1] = ipnsPath.Segments()[i+1]
	}
	b, err := Cat(n, path.Join(p), timeout)
	if err != nil {
		return ret, err
	}
	return b, nil
}
