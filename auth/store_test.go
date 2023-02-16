package auth

import (
	"testing"
	"github.com/stretchr/testify/require"
)

// Test creating KeyStore object
func TestCreateStore(t *testing.T) {
	store, err := CreateKeyStore()
	require.Equal(t, store != nil, true)
	require.Equal(t, err == nil, true)
}

// Test inserting a key
func TestStoreInsertKey(t *testing.T) {
	store, _ := CreateKeyStore()
	key, err := store.Insert("grub", []byte("password"))
	require.Equal(t, key != nil, true)
	require.Equal(t, err == nil, true)
}

// Test inserting two keys
func TestStoreInsertTwoKeys(t *testing.T) {
	store, _ := CreateKeyStore()
	store.Insert("grub", []byte("password"))
	key, err := store.Insert("grub2", []byte("password2"))
	require.Equal(t, key != nil, true)
	require.Equal(t, err == nil, true)
}

// Testing inserting a key twice fails
func TestStoreInsertTwoSameKey(t *testing.T) {
	store, _ := CreateKeyStore()
	store.Insert("grub", []byte("password"))
	_, err := store.Insert("grub", []byte("password2"))
	require.Equal(t, err != nil, true)
}

// Test querying existing key
func TestStoreQueryKey(t *testing.T) {
	store, _ := CreateKeyStore()
	store.Insert("grub", []byte("password"))
	key, ok := store.Query("grub")
	require.Equal(t, key, []byte("password"))
	require.Equal(t, ok, true)
}

// Test querying non-existing key
func TestStoreQueryBadKey(t *testing.T) {
	store, _ := CreateKeyStore()
	key, ok := store.Query("grub")
	require.Equal(t, key, []byte(nil))
	require.Equal(t, ok, false)
}
