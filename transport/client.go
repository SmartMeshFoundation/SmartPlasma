package transport

import (
	"context"
	"fmt"
	"net/rpc"
	"time"

	"github.com/SmartMeshFoundation/Spectrum/accounts/abi/bind"
	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/pkg/errors"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/account"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/backend"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/build"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/mediator"
	"github.com/SmartMeshFoundation/SmartPlasma/contract/rootchain"
)

const (
	tcpProtocol = "tcp"
)

// Errors.
var (
	ErrTransactor = errors.New("transactor is missing")
)

// Client is RPC client for PlasmaCash.
type Client struct {
	connect          *rpc.Client
	backend          backend.Backend
	sessionMediator  *mediator.MediatorSession
	sessionRootChain *rootchain.RootChainSession
	opts             *account.PlasmaTransactOpts
	timeout          uint64 // in seconds
	med              *build.Contract
	root             *build.Contract
}

// NewClient creates new PlasmaCash client.
// The Client must initialize RemoteEthereumClient or DirectEthereumClient.
func NewClient(timeout uint64, opts *account.PlasmaTransactOpts) *Client {
	return &Client{
		timeout: timeout,
		opts:    opts,
	}
}

// RemoteEthereumClient initializes work with remote ethereum client.
// Ethereum transactions are generated locally, signed locally,
// packaged and sent to the server. If this function is not called,
// then all transactions are sent directly to the Ethereum.
func (c *Client) RemoteEthereumClient(root, med *build.Contract) {
	c.med = med
	c.root = root
}

// DirectEthereumClient initializes work with direct ethereum client.
func (c *Client) DirectEthereumClient(opts bind.TransactOpts,
	mediatorAddress, rootChainAddress common.Address, backend backend.Backend) {
	mSession, err := mediator.NewMediatorSession(
		opts, mediatorAddress, backend)
	if err != nil {
		panic(err)
	}

	rSession, err := rootchain.NewRootChainSession(
		opts, rootChainAddress, backend)
	if err != nil {
		panic(err)
	}

	c.sessionRootChain = rSession
	c.sessionMediator = mSession
	c.backend = backend
}

// Connect tries to connect to a PlasmaCash RPC server.
func (c *Client) Connect(address string, port uint16) error {
	client, err := rpc.DialHTTP(tcpProtocol,
		fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		return err
	}

	c.connect = client
	return nil
}

// ConnectString tries to connect to a PlasmaCash RPC server.
func (c *Client) ConnectString(str string) error {
	client, err := rpc.DialHTTP(tcpProtocol, str)
	if err != nil {
		return err
	}

	c.connect = client
	return nil
}

// Close closes connection to PlasmaCash RPC server.
func (c *Client) Close() error {
	return c.connect.Close()
}

func (c *Client) newContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(
		context.Background(), time.Duration(c.timeout)*time.Second)
}

// Opts returns Plasma Cash transact options.
func (c *Client) Opts() *account.PlasmaTransactOpts {
	return c.opts
}
