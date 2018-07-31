package memory

import "testing"

var (
	testVal = []byte("")
)

func TestDBConcurrent(t *testing.T) {
	i := 1000000

	e := make(chan error)

	db := NewDB()
	go func() {
		for {
			cur, err := db.Current()
			if err != nil {
				<-e
				return
			}
			db.Get(cur + 1)
		}
	}()

	for {
		if i == 0 {
			break
		}

		err := db.Set(testVal)
		if err != nil {
			t.Fatal(err)
		}
		i--
	}

	_, err := db.Current()
	if err != nil {
		t.Fatal(err)
	}

	select {
	case err := <-e:
		t.Fatal(err)
	default:
	}
}
