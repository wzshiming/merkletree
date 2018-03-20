package merkletree

import (
	"encoding/hex"
	"encoding/json"
	"hash"
)

type HashData []byte

func NewHash(data []byte, h hash.Hash) HashData {
	h.Reset()
	h.Write(data)
	return HashData(h.Sum(nil))
}

func SumHash(hashs []HashData, h hash.Hash) HashData {
	h.Reset()
	for _, v := range hashs {
		h.Write(v[:])
	}
	return HashData(h.Sum(nil))
}

func (h HashData) String() string {
	return hex.EncodeToString(h[:])
}

func (h HashData) MarshalJSON() ([]byte, error) {
	return json.Marshal(h.String())
}
