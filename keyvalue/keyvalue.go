// KeyValue хранилище с сегментацией на бакеты

package keyvalue

import (
	"errors"
	"hash/fnv"
	"sync"
)

type Keyvalue[T any] interface {
	Get(key string) (string, bool)
	Set(key string, val T)
	Forget(key string)
}

type keyvalue[T any] struct {
	buckets      []bucket[T]
	bucketsCount int
	itemsCount   int
	zero         T
}

type bucket[T any] struct {
	mu    *sync.RWMutex
	items map[string]T
}

type Config struct {
	BucketsCount int
}

func New[T any](cfg Config) (*keyvalue[T], error) {
	if cfg.BucketsCount < 1 {
		return &keyvalue[T]{}, errors.New("количество бакетов keyvalue не может быть менее одного")
	}

	// Увеличивается размер на 1 для того, чтобы при рассчете номера бакета
	// в bucketNumByKey() не делать уменьшения полученного значения на 1.
	// Фактически, нулевой бакет будет всегда пустым

	buckets := make([]bucket[T], cfg.BucketsCount+1)
	for num := range buckets {
		buckets[num] = bucket[T]{
			mu:    &sync.RWMutex{},
			items: map[string]T{},
		}
	}

	return &keyvalue[T]{
		buckets:      buckets,
		bucketsCount: cfg.BucketsCount,
	}, nil
}

func (k *keyvalue[T]) Get(key string) (T, bool) {
	b := k.bucketByKey(key)
	if b == nil {
		return k.zero, false
	}

	b.mu.RLock()
	defer b.mu.RUnlock()

	val, ok := b.items[key]
	return val, ok
}

func (k *keyvalue[T]) Set(key string, val T) {
	b := k.bucketByKey(key)
	b.mu.Lock()
	defer b.mu.Unlock()

	_, ok := b.items[key]
	if !ok {
		k.itemsCount++
	}

	b.items[key] = val
}

func (k *keyvalue[T]) Forget(key string) {
	b := k.bucketByKey(key)
	if b == nil {
		return
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	_, ok := b.items[key]
	if ok {
		delete(b.items, key)
		k.itemsCount--
	}
}

func (k *keyvalue[_]) Size() int {
	return k.itemsCount
}

func (k *keyvalue[T]) bucketByKey(key string) *bucket[T] {
	num := k.bucketNumByKey(key)

	return &k.buckets[num]
}

func (k *keyvalue[_]) bucketNumByKey(key string) int {
	hash := fnv.New32a()
	hash.Write([]byte(key))

	return k.bucketsCount % int(hash.Sum32())
}
