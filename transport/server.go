package transport

import (
	"net"
	"net/http"
	"net/rpc"
	"strconv"
)

const (
	tcpProtocol = "tcp"
)

type Server struct {
	port   uint16
	server *http.Server
}

func NewServer(timeout int, port uint16) *Server {
	rpcServer := rpc.NewServer()

	rpcServer.RegisterName("SmartPlasma", &SmartPlasma{
		timeout: timeout,
	})

	httpServer := &http.Server{
		Handler: rpcServer,
	}

	return &Server{
		port:   port,
		server: httpServer,
	}
}

func (srv *Server) ListenAndServe() error {
	l, err := net.Listen(tcpProtocol, ":"+strconv.Itoa(int(srv.port)))
	if err != nil {
		return err
	}

	return srv.server.Serve(l)
}

func (srv *Server) Close() error {
	return srv.server.Close()
}
