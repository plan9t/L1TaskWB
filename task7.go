// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu   sync.RWMutex
	data map[string]string
}

func main() {
	safeMap := NewSafeMap() // Создание экземпляра мапы

	var wg sync.WaitGroup // счётчик для работы с горутинами

	// Горутина 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		safeMap.Set("key1", "value1")
	}()

	// Горутина 2
	wg.Add(1)
	go func() {
		defer wg.Done()
		safeMap.Set("key2", "value2")
	}()

	// Горутина 3 для чтения
	wg.Add(1)
	go func() {
		defer wg.Done()
		val, exists := safeMap.Get("key1")
		fmt.Printf("key1: %s, exists: %v\n", val, exists)
	}()

	wg.Wait()

	// Чтение данных из горутин
	val1, exists1 := safeMap.Get("key1")
	val2, exists2 := safeMap.Get("key2")

	fmt.Printf("key1: %s, exists1: %v\n", val1, exists1)
	fmt.Printf("key2: %s, exists2: %v\n", val2, exists2)
}

// Создание экземпляра
func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]string),
	}
}

// Установка значений
func (sm *SafeMap) Set(key, value string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

// Чтение значений по ключу
func (sm *SafeMap) Get(key string) (string, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	val, ok := sm.data[key]
	return val, ok
}
