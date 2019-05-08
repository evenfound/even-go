package core

import (
	"github.com/evenfound/even-go/net/repointer"
)

// StartPointerRepublisher - setup republisher for IPNS
func (n *EvenNode) StartPointerRepublisher() {
	n.PointerRepublisher = net.NewPointerRepublisher(n.IpfsNode, n.Datastore, n.PushNodes, n.IsModerator)
	go n.PointerRepublisher.Run()
}
