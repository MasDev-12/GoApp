package internal

type Cache struct {
	storage map[string]any
}

func (c *Cache) set(key string, value any) error {
	err := validateKey(key)
	if err != nil {
		return err
	}

	err = validateValue(value)
	if err != nil {
		return err
	}

	c.storage[key] = value

	return nil
}

func (c *Cache) Set(key string, value any) error {
	return c.set(key, value)
}

func (c *Cache) Get(key string) (any, error) {
	if err := validateKey(key); err != nil {
		return nil, err
	}
	item, found := c.storage[key]
	if !found {
		return nil, getKeyNotFoundError(key)
	}
	return item, nil
}

func (c *Cache) Delete(key string) error {
	err := validateKey(key)
	if err != nil {
		return err
	}

	if _, found := c.storage[key]; !found {
		return getKeyNotFoundError(key)
	}

	delete(c.storage, key)
	return nil
}
