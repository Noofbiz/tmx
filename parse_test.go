package tmx

import (
	"errors"
	"testing"
)

type failReader int

func (failReader) Read(b []byte) (int, error) {
	return 0, errors.New("test error")
}

func TestParseFail(t *testing.T) {
	_, err := Parse(failReader(0))
	if err == nil {
		t.Errorf("Parsed a reader when it threw an error")
	}
}
