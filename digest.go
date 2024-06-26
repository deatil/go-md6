package md6

/*
#include <stdlib.h>
#include "md6/md6_compression_hook.c"
#include "md6/md6_compress.c"
#include "md6/md6_mode.c"

md6_state *md6_ctx_new(void) {
    md6_state *state;
    if (!(state = (md6_state *)malloc(sizeof(md6_state)))) {
        return NULL;
    }

    return state;
}

void md6_ctx_free(md6_state *state) {
    if (state) {
        memset(state, 0, sizeof(md6_state));
        free(state);
    }
}
*/
import "C"

import (
    "errors"
    "unsafe"
    "runtime"
)

// The size list of a md6 hash value in bytes
const Size224 = 28
const Size256 = 32
const Size384 = 48
const Size512 = 64

// The block size of the hash algorithm in bytes.
const BlockSize = 64

type digest struct {
    md6_ctx *C.md6_state

    d int
}

func newDigest(d int) (*digest, error) {
    md6_ctx := C.md6_ctx_new()
    if md6_ctx == nil {
        return nil, errors.New("md6: Malloc error")
    }

    ret := &digest{
        md6_ctx: md6_ctx,
        d: d,
    }

    runtime.SetFinalizer(ret, func(ret *digest) {
        C.md6_ctx_free(ret.md6_ctx)
    })

    C.md6_init(md6_ctx, C.int(d))
    return ret, nil
}

func (ctx *digest) Size() int {
    return ctx.d / 8
}

func (ctx *digest) BlockSize() int {
    return BlockSize
}

func (ctx *digest) Reset() {
    C.md6_init(ctx.md6_ctx, C.int(ctx.d))
}

func (ctx *digest) Write(data []byte) (nn int, err error) {
    if len(data) == 0 {
        return
    }

    nn = len(data)

    C.md6_update(ctx.md6_ctx, (*C.uchar)(unsafe.Pointer(&data[0])), C.uint64_t(uint64(len(data))*8))
    return
}

func (ctx *digest) Sum(in []byte) []byte {
    // Make a copy of d0 so that caller can keep writing and summing.
    d0 := *ctx
    sum := d0.checkSum()
    return append(in, sum...)
}

func (ctx *digest) checkSum() (out []byte) {
    outbuf := make([]byte, ctx.d / 8)
    C.md6_final(ctx.md6_ctx, (*C.uchar)(unsafe.Pointer(&outbuf[0])))
    return outbuf
}

func sum(d int, data []byte) (out []byte, err error) {
    nn := C.uint64_t(uint64(len(data))*8)

    outbuf := make([]byte, d / 8)
    e := C.md6_hash(C.int(d), (*C.uchar)(unsafe.Pointer(&data[0])), nn, (*C.uchar)(unsafe.Pointer(&outbuf[0])))
    if e != 0 {
        return nil, errors.New("md6: fail hash")
    }

    return outbuf, nil
}
