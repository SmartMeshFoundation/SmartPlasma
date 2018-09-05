package service

import (
	"bytes"
	"context"
	"testing"

	"github.com/SmartMeshFoundation/Spectrum/common"
)

func TestAcceptUidState(t *testing.T) {
	i := newInstance(t)
	if err := i.service.currentChpt.AddCheckpoint(one, two); err != nil {
		t.Fatal(err)
	}
}

func TestInitCheckpoint(t *testing.T) {
	i := newInstance(t)

	err := i.service.currentChpt.AddCheckpoint(one, two)
	if err != nil {
		t.Fatal(err)
	}

	_, err = i.service.BuildCheckpoint()
	if err != nil {
		t.Fatal(err)
	}

	hash := i.service.CurrentCheckpoint().Hash()

	i.service.InitCheckpoint()

	hash2 := i.service.CurrentCheckpoint().Hash()

	if bytes.Equal(hash.Bytes(), hash2.Bytes()) || (hash2 != common.Hash{}) {
		t.Fatal("the current checkpoint was not initialized")
	}
}

func TestCreateUidStateProof(t *testing.T) {
	i := newInstance(t)

	err := i.service.currentChpt.AddCheckpoint(one, two)
	if err != nil {
		t.Fatal(err)
	}

	_, err = i.service.BuildCheckpoint()
	if err != nil {
		t.Fatal(err)
	}

	chpt := i.service.CurrentCheckpoint()

	tx, err := i.service.SendChptHash(context.Background(), chpt.Hash())
	if err != nil {
		t.Fatal(err)
	}

	err = i.service.mineTx(context.Background(), tx)
	if err != nil {
		t.Fatal(err)
	}

	err = i.service.SaveCheckpointToDB(chpt)
	if err != nil {
		t.Fatal(err)
	}

	proof, err := i.service.CreateUIDStateProof(one, chpt.Hash())
	if err != nil {
		t.Fatal(err)
	}

	if len(proof) == 0 {
		t.Fatal("empty proof")
	}

	valid, err := i.service.IsValidCheckpoint(
		context.Background(), one, two, chpt.Hash(), proof)
	if err != nil {
		t.Fatal(err)
	}

	if !valid {
		t.Fatal("checkpoint invalid")
	}

	valid, err = i.service.IsValidCheckpoint(
		context.Background(), one, three, chpt.Hash(), proof)
	if err != nil {
		t.Fatal(err)
	}

	if valid {
		t.Fatal("checkpoint valid")
	}

	proof2, err := i.service.CreateUIDStateProof(two, chpt.Hash())
	if err != nil {
		t.Fatal(err)
	}

	if len(proof2) == 0 {
		t.Fatal("empty proof")
	}

	valid, err = i.service.IsValidCheckpoint(
		context.Background(), two, two, chpt.Hash(), proof2)
	if err != nil {
		t.Fatal(err)
	}

	if valid {
		t.Fatal("checkpoint valid")
	}
}
