package internal

func NewCache() *Cache {
	return &Cache{
		storage: make(map[string]any),
	}
}
