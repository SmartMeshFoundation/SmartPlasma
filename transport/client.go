package transport

import (
	"fmt"
	"math/big"
	"net/rpc"

	"github.com/ethereum/go-ethereum/common"
)

// Smart Plasma RPC Methods.
const (
	AcceptTransactionMethod   = "SmartPlasma.AcceptTransaction"
	CreateProofMethod         = "SmartPlasma.CreateProof"
	AddCheckpointMethod       = "SmartPlasma.AddCheckpoint"
	CreateUIDStateProofMethod = "SmartPlasma.CreateUIDStateProof"
)

// Client is RPC client for PlasmaCash.
type Client struct {
	connect *rpc.Client
	timeout uint64
}

// NewClient creates new PlasmaCash client.
func NewClient(timeout uint64) *Client {
	return &Client{
		timeout: timeout,
	}
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

// AcceptTransaction sends raw transaction to PlasmaCash RPC server.
func (c *Client) AcceptTransaction(rawTx []byte) (resp *AcceptTransactionResp,
	err error) {
	req := &AcceptTransactionReq{rawTx}

	if err = c.connect.Call(AcceptTransactionMethod, req, &resp); err != nil {
		return nil, err
	}

	return resp, err
}

// CreateProof sends uid and block number to PlasmaCash RPC server.
// Returns merkle proof for a uid.
func (c *Client) CreateProof(uid *big.Int,
	block uint64) (resp *CreateProofResp, err error) {
	req := &CreateProofReq{UID: uid, Block: block}

	if err = c.connect.Call(CreateProofMethod, req, &resp); err != nil {
		return nil, err
	}

	return resp, err
}

// AddCheckpoint sends uid and current transaction nonce
// for inclusion in a checkpoint.
func (c *Client) AddCheckpoint(uid,
	nonce *big.Int) (resp *AddCheckpointResp, err error) {
	req := &AddCheckpointReq{
		UID:   uid,
		Nonce: nonce,
	}

	if err = c.connect.Call(AddCheckpointMethod, req, &resp); err != nil {
		return nil, err
	}

	return resp, err
}

// CreateUIDStateProof sends uid and checkpoint hash to PlasmaCash RPC server.
// Returns merkle proof for a uid.
func (c *Client) CreateUIDStateProof(uid *big.Int,
	checkpointHash common.Hash) (resp *CreateUIDStateProofResp, err error) {
	req := &CreateUIDStateProofReq{
		UID:            uid,
		CheckpointHash: checkpointHash,
	}

	if c.connect.Call(CreateUIDStateProofMethod, req, &resp); err != nil {
		return nil, err
	}

	return resp, err
}
