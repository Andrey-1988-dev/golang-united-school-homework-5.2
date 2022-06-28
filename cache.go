package cache

import (
	"time"
)

type Items struct {
	deadline         time.Time
	value            string
	infinityDeadline bool
}

type Cache struct {
	item map[string]Items
}

func NewCache() Cache {
	return Cache{item: make(map[string]Items)}
}

func (c Cache) Get(key string) (string, bool) {
	var value string
	// Проверяем наличиие ключа
	items, exists := c.item[key]
	if exists {
		// Проверяем deadline
		infinityDeadline := items.infinityDeadline
		if infinityDeadline {
			// Если дедлайн не ограничен
			value = items.value
			return value, exists
		} else {
			// Если дедлайн ограничен
			deadline := items.deadline
			timeNow := time.Now()
			elapsed := deadline.Sub(timeNow).Microseconds()
			if elapsed > 0 {
				// deadline ещё не достугнут
				value = items.value
				return value, exists
			} else {
				// deadline превышен
				return value, false
			}
		}
	} else {
		// Ключ не найден
		return value, exists
	}
}

func (c Cache) Put(key, value string) {
	deadline := time.Now()
	c.item[key] = Items{
		deadline:         deadline,
		value:            value,
		infinityDeadline: true,
	}
}

func (c Cache) Keys() []string {
	var values []string
	for key := range c.item {
		// Проверяем на deadline
		_, exists := c.Get(key)
		if exists {
			values = append(values, key)
		}
	}
	return values
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.item[key] = Items{
		deadline:         deadline,
		value:            value,
		infinityDeadline: false,
	}
}
