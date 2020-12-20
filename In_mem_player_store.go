package main

import "sync"

// we probably need to have a map here that keeps track of players names and scores
type InMemoryPlayerStore struct {
	store map[string]int
	lock  sync.RWMutex
}

// this is going to update the dict that we make above
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.store[name]++
}

// this is going to return the players value in the map
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.store[name]
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
		sync.RWMutex{},
	}
}
