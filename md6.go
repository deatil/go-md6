package md6

import (
    "hash"
)

// New returns a new hash.Hash computing the md6 checksum.
func New(d int) (hash.Hash, error) {
    return newDigest(d)
}

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

// ===========

// Sum returns the md6 checksum.
func Sum(d int, data []byte) (out []byte, err error) {
    return sum(d, data)
}

// Sum224 returns the md6 checksum.
func Sum224(data []byte) (out [Size224]byte) {
    sum, _ := Sum(224, data)
    copy(out[:], sum)

    return
}

// Sum256 returns the md6 checksum.
func Sum256(data []byte) (out [Size256]byte) {
    sum, _ := Sum(256, data)
    copy(out[:], sum)

    return
}

// Sum384 returns the md6 checksum.
func Sum384(data []byte) (out [Size384]byte) {
    sum, _ := Sum(384, data)
    copy(out[:], sum)

    return
}

// Sum512 returns the md6 checksum.
func Sum512(data []byte) (out [Size512]byte) {
    sum, _ := Sum(512, data)
    copy(out[:], sum)

    return
}
