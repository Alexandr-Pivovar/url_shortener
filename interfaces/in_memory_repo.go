package interfaces

import (
	"errors"
	"sync"
)

type Store struct {
	urls map[string]string
	mu   sync.Mutex
}

func New() *Store {
	return &Store{
		urls: make(map[string]string),
	}
}

func (s *Store) Save(key, url string) error {
	s.mu.Lock()

	_, ok := s.urls[url]
	if ok {
		s.mu.Unlock()
		return errors.New("key exists")
	}
	s.urls[key] = url

	s.mu.Unlock()
	return nil
}

func (s *Store) Get(key string) (string, error) {
	s.mu.Lock()

	shortUrl, ok := s.urls[key]
	if !ok {
		s.mu.Unlock()
		return "", errors.New("key does not exist")
	}

	s.mu.Unlock()
	return shortUrl, nil
}
