package filesig

import (
	"bytes"
	"testing"
)

func TestPdf(t *testing.T) {
	buff := bytes.NewReader([]byte("%PDF-1.3"))

	valid := IsPdf(buff)
	if !valid {
		t.Error("error: buffer not valid PDF file")
	}
}

func TestMpg0(t *testing.T) {
	buff := bytes.NewReader(MPG_0)

	valid := IsMpg(buff)
	if !valid {
		t.Error("error: buffer not valid MPG file")
	}
}

func TestMpg1(t *testing.T) {
	buff := bytes.NewReader(MPG_1)

	valid := IsMpg(buff)
	if !valid {
		t.Error("error: buffer not valid MPG file")
	}
}
