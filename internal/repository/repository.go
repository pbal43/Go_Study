package repository

type Storage struct {
	store map[string]string
}

func NewStorage() *Storage {
	return &Storage{store: make(map[string]string)}
}

// Сеттер
func (s *Storage) Set(key, value string) {
	s.store[key] = value
}

// Геттер
func (s *Storage) Get(key string) string {
	return s.store[key]
}

// Сокрыт от использования в других пакетах, так как написан с маленькой буквы
func (s *Storage) clearMap() {
	for key := range s.store {
		delete(s.store, key)
	}
}
