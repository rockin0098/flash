/*
 *  Copyright (c) 2017, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mtproto

import (
	"bufio"
	"net"
)

type BufferedReader struct {
	r        *bufio.Reader
	net.Conn // So that most methods are embedded
}

func NewBufferedReader(c net.Conn) *BufferedReader {
	return &BufferedReader{bufio.NewReader(c), c}
}

// func NewBufferedReaderSize(c net.Conn, n int) *BufferedReader {
// 	return &BufferedReader{bufio.NewReaderSize(c, n), c}
// }

func (b *BufferedReader) Peek(n int) ([]byte, error) {
	return b.r.Peek(n)
}

func (b *BufferedReader) Read(p []byte) (int, error) {
	return b.r.Read(p)
}

func (b *BufferedReader) Discard(n int) (int, error) {
	return b.r.Discard(n)
}

func (b *BufferedReader) BufioReader() *bufio.Reader {
	return b.r
}
