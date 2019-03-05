package rpc

import (
	rpcWallet "github.com/evenfound/even-go/node/server/rpc/wallet"
	"github.com/powerman/rpc-codec/jsonrpc2"
	"io"
	"net"
	"net/http"
	"net/rpc"
)

var Conn net.Conn

type RPCConnection struct {
	in  io.Reader
	out io.Writer
}

func (c *RPCConnection) Read(p []byte) (n int, err error) {
	return c.in.Read(p)
}

func (c *RPCConnection) Write(d []byte) (n int, err error) {
	return c.out.Write(d)
}

func (c *RPCConnection) Close() error {
	return nil
}

// The Serve function listens json-rpc server on specified port
func Serve(port string) {
	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		panic(err)
	}

	defer listener.Close()

	var (
		wallet  = new(rpcWallet.Wallet)
		account = new(rpcWallet.Account)
	)

	server := rpc.NewServer()
	server.Register(wallet)
	server.Register(account)

	http.Serve(listener, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		serverCodec := jsonrpc2.NewServerCodec(&RPCConnection{in: r.Body, out: w}, server)

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(200)

		if err1 := server.ServeRequest(serverCodec); err1 != nil {
			http.Error(w, "Error while serving JSON request", 500)
			return
		}

	}))
}
