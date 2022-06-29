package grabber

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Read(id, hash int) (byte, error) {
	time.Sleep(500 * time.Millisecond)
	if rand.Int()%19 == 0 {
		return 0, errors.New("smth went bad, pls rerun")
	}
	t := 0
	if id == 0 {
		t = hash
	} else {
		t = hash % id
	}
	r := fmt.Sprintf("%d", t)
	mds := sha1.Sum([]byte(r))
	return (mds[19] + mds[19]) / 2, nil
}

func Run(num int, hash int) []byte {
	return result
}
