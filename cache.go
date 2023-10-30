package main

type Cache interface {
	Set(key any, value any) bool
	Get(key any) any
	Delete(key any) bool
}

type Store struct {
	item map[any]any
}

func NewCacheStore() Store {
	return Store{item: make(map[any]any)}
}

func (r Store) Set(key any, value any) bool {
	r.item[key] = value

	return true
}

func (r Store) Get(key any) any {
	return r.item[key]
}

func (r Store) Delete(key any) bool {
	delete(r.item, key)

	return true
}
