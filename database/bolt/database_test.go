package bolt

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/pborman/uuid"
	"math/big"
)

const (
	dbName = "database"
)

var (
	testVal = []byte("tx")
)

func TestDB(t *testing.T) {
	dir, err := ioutil.TempDir("", uuid.NewUUID().String())
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(dir)

	database, err := NewDB(filepath.Join(dir, dbName), BlocksBucket, nil)
	if err != nil {
		t.Fatal(err)
	}

	var number int64 = 5

	err = database.Set([]byte(big.NewInt(number).String()), testVal)
	if err != nil {
		t.Fatal(err)
	}

	val, err := database.Get([]byte(new(big.Int).SetInt64(number).String()))
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(testVal, val) {
		t.Fatalf("expect %s, got %s", testVal, val)
	}
}
