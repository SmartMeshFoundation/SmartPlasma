package transport

import (
	"bytes"
	"fmt"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/backend"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/transaction"
)

// SmartPlasma implements PlasmaCash methods to RPC server.
type SmartPlasma struct {
	timeout int
	backend backend.Backend
}

// SentTxReq is request for send Plasma transaction to PRC server.
type SentTxReq struct {
	Tx []byte
}

// SentTxResp is response for send Plasma transaction to PRC server.
type SentTxResp struct {
	Error string
}

// SentTx accepts a raw transaction and returns a response.
func (api *SmartPlasma) SentTx(req *SentTxReq, resp *SentTxResp) error {
	buf := bytes.NewBuffer(req.Tx)

	tx := &transaction.Transaction{}

	if err := transaction.DecodeRLP(buf, tx); err != nil {
		resp.Error = err.Error()
		return nil
	}

	fmt.Println(tx.Hash().String()) // TODO: dummy
	return nil
}
