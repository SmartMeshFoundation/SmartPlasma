package transport

import (
	"bytes"
	"math/big"
	"net/http/httptest"
	"net/rpc"
	"testing"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/account"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/transaction"
)

type testHTTPServer struct {
	server *httptest.Server
}

func testServer() *testHTTPServer {
	rpcServer := rpc.NewServer()

	rpcServer.RegisterName("SmartPlasma", &SmartPlasma{
		timeout: 100,
	})

	httpServer := httptest.NewServer(rpcServer)

	return &testHTTPServer{
		server: httpServer,
	}
}

func testClient(t *testing.T, srv *testHTTPServer) *Client {
	cli := NewClient(100)
	err := cli.ConnectString(srv.server.URL[7:])
	if err != nil {
		t.Fatal(err)
	}
	return cli
}

func TestNewClient(t *testing.T) {
	s := testServer()

	defer s.server.Close()

	cli := testClient(t, s)

	defer cli.Close()

	newOwner := account.Account(account.GenKey())

	tx, err := transaction.NewTransaction(big.NewInt(0), big.NewInt(2),
		big.NewInt(3), big.NewInt(4), newOwner.From)
	if err != nil {
		t.Fatal(err)
	}

	buf := bytes.NewBuffer([]byte{})

	err = tx.EncodeRLP(buf)
	if err != nil {
		t.Fatal(err)
	}

	req := &SentTxReq{buf.Bytes()}

	var resp *SentTxResp

	if err := cli.connect.Call(SentTxMethod, req, &resp); err != nil {
		t.Fatal(err)
	}

	if resp.Error != "" {
		t.Fatal("error")
	}
}
