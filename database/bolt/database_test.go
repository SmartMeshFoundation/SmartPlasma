package bolt

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/pborman/uuid"
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

	err = database.Set(testVal)
	if err != nil {
		t.Fatal(err)
	}

	number, err := database.Current()
	if err != nil {
		t.Fatal(err)
	}

	val, err := database.Get(number)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(testVal, val) {
		t.Fatalf("expect %s, got %s", testVal, val)
	}
}
