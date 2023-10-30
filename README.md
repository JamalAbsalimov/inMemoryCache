# Usage example

```go
func main() {

	cacheStore := NewCacheStore()

	cacheStore.Set(123, "Москва")
	cacheStore.Set(456, "Санкт-Петербург")

	fmt.Println(cacheStore.Get(123))
	fmt.Println(cacheStore.Get(456))
	fmt.Println(cacheStore.Delete(456))
	fmt.Println(cacheStore.Delete(123))
	fmt.Println(cacheStore.Get(456))
	fmt.Println(cacheStore.Get(123))
}

```