package util

import (
	"fmt"
	"testing"
	"time"
)

var d time.Time

func TestIsZero(t *testing.T) {
	b := d.IsZero()
	fmt.Println(b)
}
