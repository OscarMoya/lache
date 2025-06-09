package memtable

type Memtable interface {
	// Set inserts or updates a key-value pair in the memtable.
	Set(key Key, value []byte)

	// Get retrieves the value associated with a key from the memtable.
	// It returns the value and a boolean indicating whether the key was found.
	Get(key Key) ([]byte, bool)

	// ScanF iterates over the memtable and applies the provided function
	ScanF(f func(key Key, value []byte) bool)
}
