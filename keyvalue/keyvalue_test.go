package keyvalue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ErrorBucketsCount(t *testing.T) {
	_, err := New[string](Config{
		BucketsCount: 0,
	})

	assert.NotNil(t, err)
}

func Test_SetGet(t *testing.T) {
	kv := newKeyvalue(2)
	kv.Set("a", "1")
	v, ok := kv.Get("a")

	assert.Equal(t, "1", v)
	assert.True(t, ok)
}

func Test_Update(t *testing.T) {
	kv := newKeyvalue(2)
	kv.Set("a", "1")
	kv.Set("a", "2")
	v, ok := kv.Get("a")

	assert.Equal(t, "2", v)
	assert.True(t, ok)
	assert.Equal(t, 1, kv.Size())
}

func Test_Forget(t *testing.T) {
	kv := newKeyvalue(2)
	kv.Set("a", "1")
	kv.Forget("a")
	_, ok := kv.Get("a")

	assert.False(t, ok)
}

func Test_ForgetMissing(t *testing.T) {
	kv := newKeyvalue(2)
	kv.Set("a", "1")
	kv.Forget("a")
	kv.Forget("a")
	_, ok := kv.Get("a")

	assert.False(t, ok)
}

func Test_Size(t *testing.T) {
	kv := newKeyvalue(2)
	assert.Equal(t, 0, kv.Size())

	kv.Set("a", "1")
	assert.Equal(t, 1, kv.Size())

	kv.Set("b", "2")
	assert.Equal(t, 2, kv.Size())

	kv.Forget("b")
	assert.Equal(t, 1, kv.Size())

	kv.Forget("a")
	assert.Equal(t, 0, kv.Size())
}

func Test_OneBucket(t *testing.T) {
	kv := newKeyvalue(1)

	kv.Set("a", "1")
	v, ok := kv.Get("a")
	assert.Equal(t, v, "1")
	assert.True(t, ok)
	assert.Equal(t, 1, kv.Size())

	kv.Forget("a")
	assert.Equal(t, 0, kv.Size())
}

func newKeyvalue(size int) *keyvalue[string] {
	kv, _ := New[string](Config{
		BucketsCount: size,
	})

	return kv
}
