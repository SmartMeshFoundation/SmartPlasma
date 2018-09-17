package transport

import (
	"net"
	"net/http"
	"net/rpc"
	"strconv"

	"github.com/SmartMeshFoundation/SmartPlasma/service"
)

// Server is RPC server to Plasma Cash service.
type Server struct {
	port    uint16
	server  *http.Server
	service *service.Service
}

// NewServer creates new RPC server to Plasma Cash service.
func NewServer(timeout int, port uint16, service *service.Service) *Server {
	rpcServer := rpc.NewServer()

	rpcServer.RegisterName(
		"SmartPlasma", NewSmartPlasma(timeout, service))

	httpServer := &http.Server{
		Handler: rpcServer,
	}

	return &Server{
		port:    port,
		server:  httpServer,
		service: service,
	}
}

// ListenAndServe starts RPC server to Plasma Cash service.
func (srv *Server) ListenAndServe() error {
	l, err := net.Listen("tcp", ":"+strconv.Itoa(int(srv.port)))
	if err != nil {
		return err
	}

	return srv.server.Serve(l)
}

// Close stops RPC server to Plasma Cash service.
func (srv *Server) Close() error {
	srv.service.Close()
	return srv.server.Close()
}
