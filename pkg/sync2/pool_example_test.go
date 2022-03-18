// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package sync2_test

import (
	"bytes"
	"io"
	"os"
	"time"

	"gopkg.in/typ.v3/pkg/sync2"
)

var bufPool = sync2.Pool[*bytes.Buffer]{
	New: func() *bytes.Buffer {
		// The Pool's New function should generally only return pointer
		// types, since a pointer can be put into the return interface
		// value without an allocation:
		return new(bytes.Buffer)
	},
}

// timeNow is a fake version of time.Now for tests.
func timeNow() time.Time {
	return time.Unix(1136214245, 0)
}

func Log(w io.Writer, key, val string) {
	b := bufPool.Get()
	b.Reset()
	// Replace this with time.Now() in a real logger.
	b.WriteString(timeNow().UTC().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	w.Write(b.Bytes())
	bufPool.Put(b)
}

func ExamplePool() {
	Log(os.Stdout, "path", "/search?q=flowers")
	// Output: 2006-01-02T15:04:05Z path=/search?q=flowers
}
