package secrets

import "errors"

// SecretStore defines an interface to read and write secrets.
type SecretStore interface {
	WriteSecret(key, value string) error
	ReadSecret(key string) (string, error)
}

// FileSecretStore is a simple implementation that stores secrets in memory (for now).
type FileSecretStore struct {
	store map[string]string
}

func NewFileSecretStore() *FileSecretStore {
	return &FileSecretStore{store: make(map[string]string)}
}

func (f *FileSecretStore) WriteSecret(key, value string) error {
	f.store[key] = value
	return nil
}

func (f *FileSecretStore) ReadSecret(key string) (string, error) {
	val, ok := f.store[key]
	if !ok {
		return "", errors.New("secret not found")
	}
	return val, nil
}
