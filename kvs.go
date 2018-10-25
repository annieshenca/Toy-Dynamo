// KVS.go
//
// CMPS 128 Fall 2018
//
// Lawrence Lawson            lelawson
// Pete Wilcox                pcwilcox
//
// This source file defines the KVS object used for a local data store. The struct
// implements a map to store key-value pairs, and implements the dbAccess interface
// defined in dbAccess.go in order to serve as a db object for the REST API.
//

package main

import "sync"

// KVS represents a key-value store and implements the dbAccess interface
type KVS struct {
	db    map[string]string
	mutex *sync.RWMutex
}

// NewKVS initializes a KVS object and returns a pointer to it function
func NewKVS() *KVS {
	var k KVS
	k.db = make(map[string]string)
	var m sync.RWMutex
	k.mutex = &m
	return &k
}

// Contains returns true if the dbAccess object contains an object with key equal to the input
func (k *KVS) Contains(key string) bool {
	// Grab a read lock
	k.mutex.RLock()
	defer k.mutex.RUnlock()

	// Once the read lock has been obtained, call the non-locking contains() method
	ok := k.contains(key)
	return ok
}

// contains is the unexported version of Contains() and does not hold a read lock
func (k *KVS) contains(key string) bool {
	_, ok := k.db[key]
	return ok
}

// Get returns the value associated with a particular key. If the key does not exist it returns ""
func (k *KVS) Get(key string) string {
	// Grab a read lock
	k.mutex.RLock()
	defer k.mutex.RUnlock()

	// Call the non-locking contains() method
	if k.contains(key) {
		return k.db[key]
	}
	return ""
}

// Delete removes a key-value pair from the object. If the key does not exist it returns false.
func (k *KVS) Delete(key string) bool {
	// Grab a write lock
	k.mutex.Lock()
	defer k.mutex.Unlock()

	// Call the nonlocking contains method
	if k.contains(key) {
		delete(k.db, key)
		return true
	}
	return false

}

// Put adds a key-value pair to the DB. If the key already exists, then it overwrites the existing value. If the key does not exist then it is added.
func (k *KVS) Put(key string, val string) bool {
	maxVal := 1048576 // 1 megabyte
	maxKey := 200     // 200 characters
	keyLen := len(key)
	valLen := len(val)

<<<<<<< HEAD
	if keyLen < maxKey && valLen < maxVal {
		if k.Contains(key) {
			k.db[key] = val
			return true
		} else if !(k.Contains(key)) {
=======
	if keyLen <= maxKey && valLen <= maxVal {
		// Grab a write lock
		k.mutex.Lock()
		defer k.mutex.Unlock()
		if k.contains(key) {
>>>>>>> hw2
			k.db[key] = val
			return true
		}
		k.db[key] = val
		return true
	}
	return false
}

// ServiceUp returns true if the interface is able to communicate with the DB
func (k *KVS) ServiceUp() bool {
	return true
}
