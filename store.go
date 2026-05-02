package main

import (
	"fmt"
	"sync"
)

type Store struct {
	data map[string]string  //Simple hashmap (short -> long)
	mu sync.RWMutex
}

func NewStore() *Store{ // Creating the store or instance of store itself 
	return &Store{
		data : make(map[string]string),
	}
}

func (s *Store) Save(short, long string) { // Saving the gen link
	s.mu.Lock()      //Write lock 
	defer s.mu.Unlock() // Unlock after successful write
	s.data[short] = long
	fmt.Println("Saving:", short, "->", long)
}

func (s *Store) GetLink(short string) (string, bool) {
	s.mu.RLock()             //Read lock ( can be used by multiple routines)
	defer s.mu.RUnlock()
	link , ok := s.data[short]   //extract the og link (ok to verify )
	return link,ok
}