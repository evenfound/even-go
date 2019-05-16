package commands

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/repo/config"

	"gx/ipfs/QmTjNRVt2fvaRFu93keEC7z5M1GS1iH6qZ9227htQioTUY/go-ipfs-cmds"
	"gx/ipfs/QmceUdzxkimdYsgtX733uNgzf1DLHyBKN6ehGSp85ayppM/go-ipfs-cmdkit"
	"gx/ipfs/QmceUdzxkimdYsgtX733uNgzf1DLHyBKN6ehGSp85ayppM/go-ipfs-cmdkit/files"
)

type Context struct {
	Online     bool
	ConfigRoot string
	ReqLog     *ReqLog

	config     *config.Config
	LoadConfig func(path string) (*config.Config, error)

	node          *core.IpfsNode
	ConstructNode func() (*core.IpfsNode, error)
}

// GetConfig returns the config of the current Command execution
// context. It may load it with the provided function.
func (c *Context) GetConfig() (*config.Config, error) {
	var err error
	if c.config == nil {
		if c.LoadConfig == nil {
			return nil, errors.New("nil LoadConfig function")
		}
		c.config, err = c.LoadConfig(c.ConfigRoot)
	}
	return c.config, err
}

// GetNode returns the node of the current Command execution
// context. It may construct it with the provided function.
func (c *Context) GetNode() (*core.IpfsNode, error) {
	var err error
	if c.node == nil {
		if c.ConstructNode == nil {
			return nil, errors.New("nil ConstructNode function")
		}
		c.node, err = c.ConstructNode()
	}
	return c.node, err
}

// Context returns the node's context.
func (c *Context) Context() context.Context {
	n, err := c.GetNode()
	if err != nil {
		log.Debug("error getting node: ", err)
		return context.Background()
	}

	return n.Context()
}

// LogRequest adds the passed request to the request log and
// returns a function that should be called when the request
// lifetime is over.
func (c *Context) LogRequest(req *cmds.Request) func() {
	rle := &ReqLogEntry{
		StartTime: time.Now(),
		Active:    true,
		Command:   strings.Join(req.Path, "/"),
		Options:   req.Options,
		Args:      req.Arguments,
		ID:        c.ReqLog.nextID,
		log:       c.ReqLog,
	}
	c.ReqLog.AddEntry(rle)

	return func() {
		c.ReqLog.Finish(rle)
	}
}

// Close cleans up the application state.
func (c *Context) Close() {
	// let's not forget teardown. If a node was initialized, we must close it.
	// Note that this means the underlying req.Context().Node variable is exposed.
	// this is gross, and should be changed when we extract out the exec Context.
	if c.node != nil {
		log.Info("Shutting down node...")
		c.node.Close()
	}
}

// Request represents a call to a command from a consumer
type Request interface {
	Path() []string
	Option(name string) *cmdkit.OptionValue
	Options() cmdkit.OptMap
	Arguments() []string
	StringArguments() []string
	Files() files.File
	Context() context.Context
	InvocContext() *Context
	Command() *Command
}
