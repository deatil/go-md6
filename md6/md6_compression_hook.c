/* File:    md6.h
** Author:  nabijaczleweli
** Email:   nabijaczleweli@gmail.com
** Date:    4/22/2020
**
** (The following license is known as "The MIT License")
**
** Copyright (c) 2020 nabijaczleweli
**
** Permission is hereby granted, free of charge, to any person obtaining a copy
** of this software and associated documentation files (the "Software"), to deal
** in the Software without restriction, including without limitation the rights
** to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
** copies of the Software, and to permit persons to whom the Software is
** furnished to do so, subject to the following conditions:
**
** The above copyright notice and this permission notice shall be included in
** all copies or substantial portions of the Software.
**
** THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
** IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
** FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
** AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
** LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
** OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
** THE SOFTWARE.
**
** (end of license)
**
** This file provides the definition of the global compression_hook variable.
** The original code had it in md6.h, which made Clang 11 unhappy when linking:
** see https://github.com/nabijaczleweli/md6-rs/issues/1.
*/


#include "md6.h"


void (* compression_hook)(md6_word *C,
			  const md6_word *Q,
			  md6_word *K,
			  int ell,
			  int i,
			  int r,
			  int L,
			  int z,
			  int p,
			  int keylen,
			  int d,
			  md6_word *N
			  );
