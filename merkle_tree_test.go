package merkletree

import (
	"crypto/sha256"
	"fmt"
	"testing"

	ffmt "gopkg.in/ffmt.v1"
)

func TestB(t *testing.T) {

	hash := sha256.New()

	mt := NewMerkleTree(3, hash)

	for i := 0; i != 10; i++ {
		mt.Write([]byte(fmt.Sprintln("hello", i)))
	}

	ffmt.Mark(mt.SumHash())

}
