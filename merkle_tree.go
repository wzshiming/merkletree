package merkletree

import (
	"hash"
)

type MerkleTreeLevel uint32

const (
	leafLevel = MerkleTreeLevel(1) // 最低节点
)

type MerkleTree struct {
	tree      map[MerkleTreeLevel][]HashData
	maxLevel  MerkleTreeLevel
	maxDegree int
	hash      hash.Hash
}

func NewMerkleTree(maxDegree int, h hash.Hash) *MerkleTree {
	if maxDegree < 2 {
		maxDegree = 2
	}
	return &MerkleTree{
		tree:      map[MerkleTreeLevel][]HashData{},
		maxLevel:  leafLevel,
		maxDegree: maxDegree,
		hash:      h,
	}
}

// 添加hash 如果生成高层的节点则删除低层的
func (m *MerkleTree) Append(nextHash HashData) error {
	m.tree[leafLevel] = append(m.tree[leafLevel], nextHash)
	currentLevel := leafLevel
	for {
		currentLevelHashes := m.tree[currentLevel]
		if len(currentLevelHashes) <= m.maxDegree {
			return nil
		}
		nextLevelHash := SumHash(currentLevelHashes, m.hash)

		delete(m.tree, currentLevel)
		nextLevel := currentLevel + 1
		m.tree[nextLevel] = append(m.tree[nextLevel], nextLevelHash)
		if nextLevel > m.maxLevel {
			m.maxLevel = nextLevel
		}
		currentLevel = nextLevel
	}
}

// 已经完成不再发生变化
func (m *MerkleTree) Done() error {
	var h HashData
	for i := leafLevel; i < m.maxLevel; {
		currentLevelHashes := m.tree[i]
		switch len(currentLevelHashes) {
		case 0:
			i++
			continue
		case 1:
			h = currentLevelHashes[0]
		default:
			h = SumHash(currentLevelHashes, m.hash)
		}
		delete(m.tree, i)
		i++
		m.tree[i] = append(m.tree[i], h)
	}

	finalHashes := m.tree[m.maxLevel]
	if len(finalHashes) > m.maxDegree {
		delete(m.tree, m.maxLevel)
		m.maxLevel++
		m.tree[m.maxLevel] = []HashData{SumHash(finalHashes, m.hash)}
	}
	return nil
}

func (m *MerkleTree) GetMaxLevel() MerkleTreeLevel {
	return m.maxLevel
}

func (m *MerkleTree) GetMaxLevelHashes() []HashData {
	return m.tree[m.maxLevel]
}

func (m *MerkleTree) IsEmpty() bool {
	return m.maxLevel == 1 && len(m.tree[m.maxLevel]) == 0
}
