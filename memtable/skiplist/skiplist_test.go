package skiplist

import (
	"testing"

	"github.com/oscarmoya/lache/memtable"
)

func toKey(s string) memtable.Key {
	return memtable.ByteKey(s)
}

type nodeData struct {
	key   string
	value string
}

func TestSkipListSet_InsertAndOverwrite(t *testing.T) {
	type testCase struct {
		name    string
		inserts []nodeData
		check   string
		expect  string
	}

	tests := []testCase{
		{
			name: "insert single key-value pair",
			inserts: []nodeData{
				{"a", "1"},
			},
			check:  "a",
			expect: "1",
		},
		{
			name: "override existing key",
			inserts: []nodeData{
				{"k", "old"}, {"k", "new"},
			},
			check:  "k",
			expect: "new",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sl := New(16, 0.5) // Create a new skip list with a max level of 16 and default probability
			for _, kv := range tt.inserts {
				sl.Set(toKey(kv.key), []byte(kv.value))
			}
			val, ok := sl.Get(toKey(tt.check))
			if !ok {
				t.Fatalf("expected key %q to exist", tt.check)
			}
			if string(val) != tt.expect {
				t.Errorf("expected value %q, got %q", tt.expect, val)
			}
		})
	}
}

func TestSkipList_Get(t *testing.T) {
	type testCase struct {
		name    string
		inserts []nodeData
		queries []struct {
			key      string
			expected string
			found    bool
		}
	}

	tests := []testCase{
		{
			name: "get existing keys",
			inserts: []nodeData{
				{"a", "A"}, {"b", "B"}, {"c", "C"},
			},
			queries: []struct {
				key      string
				expected string
				found    bool
			}{
				{"a", "A", true},
				{"b", "B", true},
				{"c", "C", true},
			},
		},
		{
			name: "get non-existent key",
			inserts: []nodeData{
				{"x", "y"},
			},
			queries: []struct {
				key      string
				expected string
				found    bool
			}{
				{"not-there", "", false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sl := buildSkipList(tt.inserts)

			for _, q := range tt.queries {
				val, ok := sl.Get(toKey(q.key))
				if ok != q.found {
					t.Errorf("key=%q: expected found=%v, got %v", q.key, q.found, ok)
				}
				if ok && string(val) != q.expected {
					t.Errorf("key=%q: expected value=%q, got %q", q.key, q.expected, val)
				}
			}
		})
	}
}

func TestSkipList_Delete(t *testing.T) {
	type testCase struct {
		name    string
		inserts []nodeData
		deletes []string
		checks  []struct {
			key      string
			found    bool
			expected string
		}
	}

	tests := []testCase{
		{
			name: "delete existing key",
			inserts: []nodeData{
				{"a", "A"}, {"b", "B"}, {"c", "C"},
			},
			deletes: []string{"b"},
			checks: []struct {
				key      string
				found    bool
				expected string
			}{
				{"a", true, "A"},
				{"b", false, ""},
				{"c", true, "C"},
			},
		},
		{
			name: "delete non-existent key",
			inserts: []nodeData{
				{"x", "42"},
			},
			deletes: []string{"not-there"},
			checks: []struct {
				key      string
				found    bool
				expected string
			}{
				{"x", true, "42"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sl := buildSkipList(tt.inserts)
			for _, key := range tt.deletes {
				_ = sl.Delete(toKey(key))
			}
			for _, chk := range tt.checks {
				val, ok := sl.Get(toKey(chk.key))
				if ok != chk.found {
					t.Errorf("key=%q: expected found=%v, got %v", chk.key, chk.found, ok)
				}
				if ok && string(val) != chk.expected {
					t.Errorf("key=%q: expected value=%q, got %q", chk.key, chk.expected, val)
				}
			}
		})
	}
}

func buildSkipList(inserts []nodeData) *SkipList {
	sl := &SkipList{
		head:     &node{next: make([]*node, 16)},
		level:    3,
		size:     3,
		maxLevel: 16,
	}

	current := sl.head

	for _, kv := range inserts {
		n := &node{
			key:   toKey(kv.key),
			value: []byte(kv.value),
			next:  make([]*node, 1),
		}
		current.next[0] = n
		current = n
		sl.size++
	}

	return sl
}
