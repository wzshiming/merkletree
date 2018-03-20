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
	switch len(hashs) {
	case 0:
		return nil
	case 1:
		return hashs[0]
	}
	h.Reset()
	for _, v := range hashs {
		h.Write(v[:])
	}
	return HashData(h.Sum(nil))
}

func (h HashData) String() string {
	return "0x" + hex.EncodeToString(h[:])
}

func (h HashData) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(h[:]))
}

func (h *HashData) UnmarshalJSON(data []byte) error {
	s := ""
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	d, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	*h = HashData(d)
	return nil
}
