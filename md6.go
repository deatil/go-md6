package md6

import (
    "hash"
)

// New224 returns a new hash.Hash computing the md6 checksum.
func New224() hash.Hash {
    h, _ := New(224)
    return h
}

// New256 returns a new hash.Hash computing the md6 checksum.
func New256() hash.Hash {
    h, _ := New(256)
    return h
}

// New384 returns a new hash.Hash computing the md6 checksum.
func New384() hash.Hash {
    h, _ := New(384)
    return h
}

// New512 returns a new hash.Hash computing the md6 checksum.
func New512() hash.Hash {
    h, _ := New(512)
    return h
}
