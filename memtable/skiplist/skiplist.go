package skiplist

import (
	"math/rand"

	"github.com/oscarmoya/lache/memtable"
)

func (s *SkipList) Set(key memtable.Key, value []byte) {
	// Insert a new key-value pair into the skip list.
	// If the key already exists, update its value.

	update := make([]*node, s.maxLevel)
	current := s.head

	// Traverse the skip list to find the position for the new key.
	for i := s.level - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].key.Less(key) {
			current = current.next[i]
		}
		update[i] = current
	}

	// If the key already exists, update its value.
	if current.next[0] != nil && current.next[0].key.Equal(key) {
		current.next[0].value = value
		return
	}

	// Determine the level for the new node.
	lvl := randomLevel(s.maxLevel, s.probability)
	if lvl > s.level {
		// If the new level is greater than the current level, update the skip list's level.
		for i := s.level; i < lvl; i++ {
			update[i] = s.head // Update the update array for new levels.
		}
		s.level = lvl
	}

	newNode := &node{
		key:   key,
		value: value,
		next:  make([]*node, lvl),
	}
	for i := 0; i < lvl; i++ {
		if update[i] == nil {
			update[i] = s.head
		}
		newNode.next[i] = update[i].next[i]
		update[i].next[i] = newNode
	}

	s.size++
}

func randomLevel(maxLevel int, probability float64) int {
	level := 1
	for rand.Float64() < probability && level < maxLevel {
		level++
	}
	return level
}

func (s *SkipList) Get(key memtable.Key) ([]byte, bool) {
	// Retrieve the value associated with the given key.
	// Return the value and a boolean indicating if the key exists.
	current := s.head
	for i := s.level - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].key.Less(key) {
			current = current.next[i]
		}
	}
	if current.next[0] != nil && current.next[0].key.Equal(key) {
		return current.next[0].value, true
	}
	return nil, false
}

func (s *SkipList) ScanF(f func(key memtable.Key, value []byte) bool) {
	// Scan through the skip list and apply the function f to each key-value pair.
	current := s.head.next[0]
	if current == nil {
		return // Skip list is empty, nothing to scan.
	}
	for current != nil {
		// Apply the function f to the current key-value pair.
		if !f(current.key, current.value) {
			break
		}
		current = current.next[0]
	}

}

func (s *SkipList) Delete(key memtable.Key) bool {
	current := s.head
	update := make([]*node, s.maxLevel)
	for i := s.level - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].key.Less(key) {
			current = current.next[i]
		}
		update[i] = current
	}
	if current.next[0] != nil && current.next[0].key.Equal(key) {
		// Key found, remove it from the skip list.
		for i := 0; i < s.level; i++ {
			if update[i].next[i] != current.next[0] {
				break
			}
			update[i].next[i] = current.next[0].next[i]
		}
		s.size--
		// If the highest level is empty, reduce the level of the skip list.
		for s.level > 0 && s.head.next[s.level-1] == nil {
			s.level--
		}
		return true
	}
	return false

}

func (s *SkipList) Size() int {
	// Return the size of the skip list.
	return s.size
}

func (s *SkipList) MaxLevel() int {
	// Return the maximum level of the skip list.
	return s.maxLevel
}

func (s *SkipList) Level() int {
	// Return the current level of the skip list.
	return s.level
}

func (s *SkipList) Head() *node {
	// Return the head node of the skip list.
	return s.head
}
