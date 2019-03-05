package server

import (
	"fmt"
	"log"
	"net/http"

	httpListener "github.com/evenfound/even-go/node/server/http"
	rpcListener "github.com/evenfound/even-go/node/server/rpc"
)

const (
	portFormatter = ":%v"

	MessageListeningRPCServer  = "Listening RPC server [localhost:%v]"
	MessageListeningHTTPServer = "Listening HTTP server [localhost:%v]"
)

type Server struct{}

// ListenHTTP function listens http server on specified port
func (server *Server) ListenHTTP(port string) {

	fmt.Println(fmt.Sprintf(MessageListeningHTTPServer, port))

	r := httpListener.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(portFormatter, port), r))
}

// ListenRPC function listens rpc server on specified port
func (server *Server) ListenRPC(port string) {
	fmt.Println(fmt.Sprintf(MessageListeningRPCServer, port))
	rpcListener.Serve(port)
}
