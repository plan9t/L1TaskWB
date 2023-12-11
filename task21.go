//

package main

import "fmt"

// "Адаптер" (Adapter) используется для обеспечения совместной работы двух классов,
// которые без него не могли бы взаимодействовать из-за несовместимого интерфейса.
// Адаптер предоставляет интерфейс, ожидаемый клиентом, для существующего класса.
// MemoryCardReader - старый Memory Card Reader (устаревший интерфейс)

type MemoryCardReader struct {
}

// ReadMemoryCard - метод для чтения карты памяти старого стандарта
func (mcr *MemoryCardReader) ReadMemoryCard() string {
	return "Reading old memory card..."
}

type NewMemoryCardAdapter struct {
	OldMemoryCardReader *MemoryCardReader
}

// ReadNewMemoryCard - метод для чтения карты памяти нового стандарта через адаптер
func (adapter *NewMemoryCardAdapter) ReadNewMemoryCard() string {
	// Используем старый Memory Card Reader через адаптер
	result := adapter.OldMemoryCardReader.ReadMemoryCard()
	// Дополнительные действия для адаптации к новому стандарту, если необходимо
	result += " Adapted to new memory card standard."
	return result
}

func main() {
	// Создаем экземпляр старого Memory Card Reader
	oldReader := &MemoryCardReader{}

	// Создаем адаптер
	adapter := &NewMemoryCardAdapter{OldMemoryCardReader: oldReader}

	// Используем адаптер вместе с новым устройством
	result := adapter.ReadNewMemoryCard()

	// Выводим результат
	fmt.Println(result)
}
