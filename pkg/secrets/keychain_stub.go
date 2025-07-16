//go:build !darwin

package secrets

import "errors"

type KeychainSecretStore struct{}

func NewKeychainSecretStore() *KeychainSecretStore {
	return &KeychainSecretStore{}
}

func (k *KeychainSecretStore) WriteTempoAPIToken(token string) error {
	return errors.New("keychain not supported on this OS")
}

func (k *KeychainSecretStore) ReadTempoAPIToken() (string, error) {
	return "", errors.New("keychain not supported on this OS")
}

func (k *KeychainSecretStore) WriteSecret(key, value string) error {
	return errors.New("keychain not supported on this OS")
}

func (k *KeychainSecretStore) ReadSecret(key string) (string, error) {
	return "", errors.New("keychain not supported on this OS")
}
