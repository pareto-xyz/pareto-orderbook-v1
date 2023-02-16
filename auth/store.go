package auth

import "errors"

// KeyStore - Inherits the `api.KeyStore` interface
type KeyStore struct {
	db map[string][]byte
}

// CreateKeyStore - Create KeyStore instance
func CreateKeyStore() (*KeyStore, error) {
	var store KeyStore;
	store.db = make(map[string][]byte)
	return &store, nil
}

// Insert - Inserts a new (userID, key) combo
func (store *KeyStore) Insert(userID string, key []byte) ([]byte, error) {
	if _, ok := store.db[userID]; ok {
		return nil, errors.New("Insert: a key exists already")
	}
	store.db[userID] = key;
	return key, nil
}

// Query - Look up key in store
func (store *KeyStore) Query(userID string) ([]byte, bool) {
	key, ok := store.db[userID]
	if !ok {
		return nil, ok
	}
	return key, ok
}
