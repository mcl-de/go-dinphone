package dinphone

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		din := Parse("0800 2222 333")
		assert.Equal(t, "0800 2222-333", din, "")
	})
	t.Run("Test 2", func(t *testing.T) {
		din := Parse("(07021) 893666 3925")
		assert.Equal(t, "+49 7021 893666-3925", din, "")
	})
	t.Run("Test 3", func(t *testing.T) {
		din := Parse("+49 241 1696 100")
		assert.Equal(t, "+49 241 1696-100", din, "")
	})
	t.Run("Test 4", func(t *testing.T) {
		din := Parse("07021 574 246")
		assert.Equal(t, "+49 7021 574-246", din, "")
	})
	t.Run("Test 5", func(t *testing.T) {
		din := Parse("089 7446 1112")
		assert.Equal(t, "+49 89 7446-1112", din, "")
	})
	t.Run("Test 6", func(t *testing.T) {
		din := Parse("040 4106973")
		assert.Equal(t, "+49 40 4106973", din, "")
	})
	t.Run("Test 7", func(t *testing.T) {
		din := Parse("757692950")
		assert.Equal(t, "", din, "")
	})
	t.Run("Test 8", func(t *testing.T) {
		din := Parse("027725052234")
		assert.Equal(t, "", din, "")
	})
	t.Run("Test 9", func(t *testing.T) {
		din := Parse("05977 935 305")
		assert.Equal(t, "+49 5977 935-305", din, "")
	})
	t.Run("Test 10", func(t *testing.T) {
		din := Parse("0043 241 1696 100")
		assert.Equal(t, "+43 241 1696-100", din, "")
	})
}

func BenchmarkParse(b *testing.B) {
	b.Run("Test 1" , func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Parse("0800 2222 333")
		}
	})
	b.Run("Test 2" , func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Parse("(07021) 893666 3925")
		}
	})
	b.Run("Test 3" , func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Parse("0043 241 1696 100")
		}
	})
	b.Run("Test 4" , func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Parse("+49 241 1696 100")
		}
	})
}