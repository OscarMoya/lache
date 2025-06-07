# LSM Tree Development Roadmap

## Phase 1: In-Memory Infrastructure

**Goal**: Have a working in-memory cache using a skip list.

- [ ] Implement skip list with `Set` and `Get`
- [ ] Wrap skip list in a `MemTable` interface (`Set`, `Get`, `Range`)
- [ ] Unit tests for `MemTable`

**Deliverable**: An ordered in-memory key/value store.

---

## Phase 2: Bloom Filter

**Goal**: Filter out disk reads for missing keys.

- [ ] Implement a basic Bloom filter
- [ ] Add tests to measure false positive rate and memory usage

**Deliverable**: Space-efficient probabilistic key presence checker.

---

## Phase 3: SSTables

**Goal**: Persist `MemTable` contents to immutable sorted files on disk.

- [ ] Implement `sstable.Writer`: flush `MemTable` to disk
- [ ] Implement `sstable.Reader`: search for a key using binary search and Bloom filter
- [ ] Integrate Bloom filters per SSTable

**Deliverable**: Durable, searchable on-disk tables.

---

## Phase 4: Basic Compaction

**Goal**: Merge SSTables and discard obsolete keys.

- [ ] Implement `sstable.Compactor`: merge multiple SSTables
- [ ] Support versioning and tombstone cleanup

**Deliverable**: Cleaned, optimized storage via compaction.

---

## Phase 5: LSM Tree Core

**Goal**: Orchestrate `MemTable`, SSTables, and compaction.

- [ ] Implement `lsm.Tree` with basic `Set`, `Get`, `Flush`
- [ ] Search in `MemTable` first, then in SSTables in order
- [ ] Trigger flush when `MemTable` exceeds size threshold
- [ ] Trigger compaction when too many SSTables are present

**Deliverable**: A functioning basic LSM Tree.

---

## Phase 6: Testing, Benchmarking, CLI

**Goal**: Validate performance and usability.

- [ ] Add benchmarks using `testing.B`
- [ ] Implement simple CLI in `cmd/` for interactive usage (`put`, `get`, `scan`)

**Deliverable**: Usable and measurable system.