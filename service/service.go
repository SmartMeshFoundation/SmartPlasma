package service

import (
	"github.com/SmartMeshFoundation/SmartPlasma/blockchan/block"
	"github.com/SmartMeshFoundation/SmartPlasma/database"
)

type Service struct {
	currentBlock block.Block
	blockBase    database.Database
	chptBase     database.Database
}
