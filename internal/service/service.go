package service

import (
	"fmt"
	"go_study/internal/repository"
)

var Test = "test"

var neTest = "ne test"

type Service struct {
	name    string
	storage repository.Storage
}

func NewService(name string, storage repository.Storage) Service {
	return Service{name, storage}
}

func (s *Service) SaveData(data string) {
	s.storage.Set("data", data)
	newData := s.storage.Get("data")
	fmt.Println(newData)
}

func PrintNeTest() {
	fmt.Println(neTest)
}
