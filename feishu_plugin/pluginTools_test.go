package feishu_plugin

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrInArr(t *testing.T) {
	// mock StrInArr
	var testArr = []string{
		"a",
		"b",
		"c",
	}
	t.Logf("~> mock StrInArr")
	// do StrInArr
	t.Logf("~> do StrInArr")
	// verify StrInArr
	assert.True(t, StrInArr("c", testArr))
	assert.False(t, StrInArr("d", testArr))
}

func BenchmarkStrInArr(b *testing.B) {
	var testArr = []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
	}
	for i := 0; i < b.N; i++ {
		assert.True(b, StrInArr("f", testArr))
	}
}
