package hash

import (
	"fmt"
	"testing"
)

func TestHashCode(t *testing.T) {
	for i := 0; i < 1000; i++ {
		fmt.Println(HashCode(fmt.Sprintf("hello:%d", i))&15)
	}
}
