package task_7

import (
	"fmt"
	"sync"
)

type Cache struct {
	cars  map[string]int
	mutex sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		cars:  make(map[string]int),
		mutex: sync.RWMutex{},
	}
}

func (c *Cache) Get(name string) string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	_, ok := c.cars[name]
	if !ok {
		return "Empty"
	}
	res := fmt.Sprintf("GET %s: %v", name, c.cars[name])
	return res
}

func (c *Cache) Set(name string, value int) string {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cars[name] = value

	res := fmt.Sprintf("SET %s: %v", name, value)
	return res
}

func Task7() {
	cache := NewCache()    // инициализация пустого кеша
	wg := sync.WaitGroup{} // инициализация пустой wg (для отслеживания выполнения горутин)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(cache.Set("car1", 1))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(cache.Set("car2", 2))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(cache.Set("car1", 3))
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			fmt.Println(cache.Get(name))
		}("car1")
	}
	wg.Wait()
}
