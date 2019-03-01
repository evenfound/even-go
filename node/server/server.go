package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"log"
	"net/http"

	httpListener "github.com/evenfound/even-go/node/server/http"
)

const (
	portFormatter = ":%v"

	MessageListeningRPCServer  = "Listening RPC server [localhost:%v]"
	MessageListeningHTTPServer = "Listening HTTP server [localhost:%v]"
)

type Server struct {
	rpc RPCServer
}

func (server *Server) ListenHTTP(port string) {

	fmt.Println(fmt.Sprintf(MessageListeningHTTPServer, port))

	var router = httpListener.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(portFormatter, port), router))
}

// ListenRPC function listens rpc server on specified port
func (server *Server) ListenRPC(port string) {
	fmt.Println(fmt.Sprintf(MessageListeningRPCServer, port))
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")
	s.RegisterService(server.rpc, "")
	r := mux.NewRouter()
	r.Handle("/", s)
	http.ListenAndServe(fmt.Sprintf(portFormatter, port), r)
}
