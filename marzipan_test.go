// package marzipan provides a marzipan writer
package marzipan

import (
	"bytes"
	"testing"
)

func TestWrite(t *testing.T) {
	var b bytes.Buffer
	w := NewWriter(&b)

	w.Write([]byte("hello"))
	w.Write([]byte("world"))
	want := "hellomarzipanworld"
	if got := string(b.Bytes()); got != want {
		t.Fatalf("bad output. got = %s, want = %s", got, want)
	}

	w.Write([]byte("hej"))
	w.Write([]byte("världen"))
	want += "hejmarzipanvärlden"
	if got := string(b.Bytes()); got != want {
		t.Fatalf("bad output. got = %s, want = %s", got, want)
	}
}
