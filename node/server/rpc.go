package server

import (
	"net/http"
)

type RPCServer struct {
}

type Response string

type Args struct {
	A, B int
}

func (server *RPCServer) NewWallet(r *http.Request, args *Args, response *Response) error {
	*response = "asdasdasd"
	return nil
}
