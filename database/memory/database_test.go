package memory

import (
	"math/big"
	"strconv"
	"sync"
	"testing"
)

var (
	testVal = []byte("")
)

func TestDBConcurrent(t *testing.T) {
	i := 1000000

	e := make(chan error)

	var count int64

	mtx := sync.Mutex{}

	db := NewDB()
	go func() {
		for {
			mtx.Lock()
			cur := count
			mtx.Unlock()

			_, err := db.Get(new(big.Int).SetInt64(cur + 1).Bytes())
			if err != nil {
				e <- err
			}
		}
	}()

	for {
		if i == 0 {
			break
		}

		mtx.Lock()
		cur := count
		mtx.Unlock()

		err := db.Set(strconv.AppendInt(nil, cur, 10), testVal)
		if err != nil {
			t.Fatal(err)
		}
		i--
		mtx.Lock()
		count++
		mtx.Unlock()
	}

	select {
	case err := <-e:
		t.Fatal(err)
	default:
	}
}
