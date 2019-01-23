package logs

import (
	"fmt"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	now := time.Now()
	a, d, h := formatTimeHeader(now)
	fmt.Println(string(a), d, h)
}
