package sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloor(t *testing.T) {
	assert.Equal(t, 1, Floor(3/2))
}

func TestShellSort(t *testing.T) {
	arr := []int{1,7,6,2,8,2,8,4}
	ShellSort(arr)
	fmt.Println(arr)
}
