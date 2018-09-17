[![Build Status](https://travis-ci.org/SmartMeshFoundation/SmartPlasma.svg?branch=master)](https://travis-ci.org/SmartMeshFoundation/SmartPlasma)
[![GoDoc](https://godoc.org/github.com/SmartMeshFoundation/SmartPlasma?status.svg)](https://godoc.org/github.com/SmartMeshFoundation/SmartPlasma)
# SmartPlasma
Plasma, Plasma Cash &amp; Plasma XT implementation in Go

# Tests

#### For Linux and Mac
```bash
cd $GOPATH/src/github.com/SmartMeshFoundation/SmartPlasma
go test -v ./... -count=1
```

# Examples

### Simple example

The test demonstrates:
- creation of a deposit in Plasma Cash
- transfer between users
- withdrawal of tokens from Plasma Cash

Information is not displayed in the console, see comments in code.

#### For Linux and Mac
```bash
cd $GOPATH/src/github.com/SmartMeshFoundation/SmartPlasma/example/simple
go run example.go
```

### Cycle example

- Total 1000 users.
- Each user owns 20 deposits.
- Each cycle, each user transfers all his deposits to another user.
- Then the operator collects a block.
- He stores it in the database and sends the block hash to a Spectrum blockchain.

Assumptions:
1) Clients do not save transactions on their side.
2) Transactions is not checked.

#### For Linux and Mac
```bash
cd $GOPATH/src/github.com/SmartMeshFoundation/SmartPlasma/example/cycle
go run example.go
```