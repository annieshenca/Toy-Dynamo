package main

type kvs struct {
	db map[string]string
}

// Contains returns true if the dbAccess object contains an object with key equal to the input
func (k *kvs) Contains(key string) bool {
	_, ok := k.db[key]
	return ok
}

// Get returns the value associated with a particular key. If the key does not exist it returns ""
func (k *kvs) Get(key string) string {
	if k.Contains(key) {
		return k.db[key]
	}
	return ""
}

// Delete removes a key-value pair from the object. If the key does not exist it returns false.
func (k *kvs) Delete(key string) bool {
	if k.Contains(key) {
		delete(k.db, key)
		return true
	} else {
		return false
	}

}

// Put adds a key-value pair to the DB. If the key already exists, then it overwrites the existing value. If the key does not exist then it is added.
func (k *kvs) Put(key string, val string) bool {
	//k.db[key] = val
	if k.Contains(key) {
		k.db[key] = val
		return true
	} else if !(k.Contains(key)) {
		k.db[key] = val
		return true
	} else { //something something error checking will get to later
		return false
	}
}

// ServiceUp returns true if the interface is able to communicate with the DB
func (k *kvs) ServiceUp() bool {
	return true
}
