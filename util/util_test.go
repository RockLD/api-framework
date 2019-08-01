package util

import "testing"

func TestGenShortId(t *testing.T) {

	shortId, err := GenShortId()

	if shortId == "" || err != nil {
		t.Error("GenShortId failed!")
	}

	t.Log("GenShortId test pass")
}

func BenchmarkGenShortId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenShortId()
	}
}

func BenchmarkGenShortIdTimeConsuming(b *testing.B) {
	b.StopTimer()

	shortId, err := GenShortId()

	if shortId == "" || err != nil {
		b.Error(err)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		GenShortId()
	}
}
