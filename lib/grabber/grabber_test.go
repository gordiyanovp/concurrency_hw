package grabber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	out   []byte
	inNum int
	hash  int
}

func TestRead(t *testing.T) {
	td := testData{
		out:   []byte{22, 12, 43, 43, 59, 43, 43, 43, 59, 90, 43, 90, 90, 23, 43, 43, 59, 43, 90, 64, 75, 43, 90, 120, 31, 72, 23, 90, 18, 122, 43, 120, 31, 90, 43, 43, 90, 31, 3, 23, 75, 122, 43, 48, 90, 25, 63, 87, 31, 75, 72, 43, 113, 124, 90, 76, 75, 3, 56, 73, 73, 56, 3, 75, 76, 119, 90, 124, 113, 38, 43, 72, 75, 34, 31, 87, 55, 63, 119, 25, 76, 90, 10, 48, 75, 43, 10, 122, 76, 75, 119, 23, 55, 3, 91, 55, 31, 5, 75, 90},
		inNum: 100,
		hash:  3571,
	}
	r := Run(100, 3571)
	assert.Equal(t, r, td.out)
}
