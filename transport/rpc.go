package transport

import (
	"bytes"
	"fmt"

	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/backend"
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/transaction"
)

type SmartPlasma struct {
	timeout int
	backend backend.Backend
}

type SentTxReq struct {
	Tx []byte
}

type SentTxResp struct {
	Error string
}

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
