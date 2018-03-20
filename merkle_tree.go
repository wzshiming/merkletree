package merkletree

import (
	"hash"
)

const (
	leafLevel = 1 // 最低节点
)

type MerkleTree struct {
	tree      map[int][]HashData
	maxLevel  int
	maxDegree int
	hash      hash.Hash
}

func NewMerkleTree(maxDegree int, h hash.Hash) *MerkleTree {
	if maxDegree < 2 {
		maxDegree = 2
	}
	return &MerkleTree{
		tree:      map[int][]HashData{},
		maxLevel:  leafLevel,
		maxDegree: maxDegree,
		hash:      h,
	}
}

func (m *MerkleTree) BlockSize() int {
	return m.hash.BlockSize()
}

func (m *MerkleTree) Reset() {
	*m = *NewMerkleTree(m.maxDegree, m.hash)
}

func (m *MerkleTree) Write(p []byte) (n int, err error) {
	m.Append(NewHash(p, m.hash))
	return len(p), nil
}

// 添加hash 如果生成高层的节点则删除低层的
func (m *MerkleTree) Append(nextHash HashData) {
	m.tree[leafLevel] = append(m.tree[leafLevel], nextHash)
	m.update(leafLevel)
}

func (m *MerkleTree) update(currentLevel int) {
	hashs := m.tree[currentLevel]
	if len(hashs) <= m.maxDegree {
		return
	}
	next := SumHash(hashs, m.hash)

	delete(m.tree, currentLevel)
	nextLevel := currentLevel + 1
	m.tree[nextLevel] = append(m.tree[nextLevel], next)
	if nextLevel > m.maxLevel {
		m.maxLevel = nextLevel
	}
	m.update(nextLevel)
}

func (m *MerkleTree) Sum() []byte {
	return m.SumHash()[:]
}

func (m *MerkleTree) SumHash() HashData {
	return m.sum(m.maxLevel)
}

func (m *MerkleTree) sum(cl int) HashData {

	if cl == 0 {
		return nil
	}
	sum := m.tree[cl]

	if ms := m.sum(cl - 1); ms != nil {
		sum = append(sum, ms)
	}

	return SumHash(sum, m.hash)
}

func (m *MerkleTree) GetMaxLevel() int {
	return m.maxLevel
}

func (m *MerkleTree) IsEmpty() bool {
	return m.maxLevel == 1 && len(m.tree[m.maxLevel]) == 0
}
