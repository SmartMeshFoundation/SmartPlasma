package transport

import (
	"net"
	"net/http"
	"net/rpc"
	"strconv"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/backend"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/checkpoints"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block/transactions"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
	"github.com/SmartMeshFoundation/SmartPlasma/database"
)

const (
	tcpProtocol = "tcp"
)

// Server is RPC server to Plasma Cash service.
type Server struct {
	port   uint16
	server *http.Server
}

// NewServer creates new RPC server to Plasma Cash service.
func NewServer(timeout int, port uint16,
	session *rootchain.RootChainSession, backend backend.Backend,
	blockBase, chptBase database.Database) *Server {
	rpcServer := rpc.NewServer()

	rpcServer.RegisterName("SmartPlasma", &SmartPlasma{
		timeout:      timeout,
		currentChpt:  checkpoints.NewBlock(),
		currentBlock: transactions.NewTxBlock(),
		blockBase:    blockBase,
		chptBase:     chptBase,
		session:      session,
		backend:      backend,
	})

	httpServer := &http.Server{
		Handler: rpcServer,
	}

	return &Server{
		port:   port,
		server: httpServer,
	}
}

// ListenAndServe starts RPC server to Plasma Cash service.
func (srv *Server) ListenAndServe() error {
	l, err := net.Listen(tcpProtocol, ":"+strconv.Itoa(int(srv.port)))
	if err != nil {
		return err
	}

	return srv.server.Serve(l)
}

// Close stops RPC server to Plasma Cash service.
func (srv *Server) Close() error {
	return srv.server.Close()
}
