package transport

import (
	"fmt"
	"net/rpc"
)

// Smart Plasma RPC Methods.
const (
	SentTxMethod = "SmartPlasma.SentTx"
)

type Client struct {
	connect *rpc.Client
	timeout uint64
}

func NewClient(timeout uint64) *Client {
	return &Client{
		timeout: timeout,
	}
}

func (c *Client) Connect(address string, port uint16) error {
	client, err := rpc.DialHTTP(tcpProtocol,
		fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		return err
	}

	c.connect = client
	return nil
}

func (c *Client) ConnectString(str string) error {
	client, err := rpc.DialHTTP(tcpProtocol, str)
	if err != nil {
		return err
	}

	c.connect = client
	return nil
}

func (c *Client) Close() error {
	return c.connect.Close()
}

func (c *Client) SentTx(rawTx []byte) (resp *SentTxResp, err error) {
	req := &SentTxReq{rawTx}

	if err := c.connect.Call(SentTxMethod, req, &resp); err != nil {
		return nil, err
	}

	return resp, err
}
