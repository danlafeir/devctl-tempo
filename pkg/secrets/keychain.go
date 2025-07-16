//go:build darwin

package secrets

import (
	"errors"
	"github.com/keybase/go-keychain"
)

// KeychainSecretStore implements SecretStore using the MacOS Keychain.
const tempoService = "cli.devctl.tempo"
const tempoAccount = "tempo_api_token"

type KeychainSecretStore struct{}

func NewKeychainSecretStore() *KeychainSecretStore {
	return &KeychainSecretStore{}
}

// WriteTempoAPIToken stores the Tempo API token securely in the Keychain.
func (k *KeychainSecretStore) WriteTempoAPIToken(token string) error {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(tempoService)
	item.SetAccount(tempoAccount)
	item.SetLabel(tempoAccount)
	item.SetData([]byte(token))
	item.SetAccessible(keychain.AccessibleWhenUnlocked)
	item.SetSynchronizable(keychain.SynchronizableNo)
	_ = keychain.DeleteGenericPasswordItem(tempoService, tempoAccount)
	return keychain.AddItem(item)
}

// ReadTempoAPIToken retrieves the Tempo API token from the Keychain.
func (k *KeychainSecretStore) ReadTempoAPIToken() (string, error) {
	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassGenericPassword)
	query.SetService(tempoService)
	query.SetAccount(tempoAccount)
	query.SetReturnData(true)
	query.SetMatchLimit(keychain.MatchLimitOne)
	results, err := keychain.QueryItem(query)
	if err != nil {
		return "", err
	}
	if len(results) == 0 || len(results[0].Data) == 0 {
		return "", errors.New("secret not found")
	}
	return string(results[0].Data), nil
}

// Generic interface methods remain for compatibility
func (k *KeychainSecretStore) WriteSecret(key, value string) error {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(tempoService)
	item.SetAccount(key)
	item.SetLabel(key)
	item.SetData([]byte(value))
	item.SetAccessible(keychain.AccessibleWhenUnlocked)
	item.SetSynchronizable(keychain.SynchronizableNo)
	// Delete any existing item first
	_ = keychain.DeleteGenericPasswordItem(tempoService, key)
	return keychain.AddItem(item)
}

func (k *KeychainSecretStore) ReadSecret(key string) (string, error) {
	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassGenericPassword)
	query.SetService(tempoService)
	query.SetAccount(key)
	query.SetReturnData(true)
	query.SetMatchLimit(keychain.MatchLimitOne)
	results, err := keychain.QueryItem(query)
	if err != nil {
		return "", err
	}
	if len(results) == 0 || len(results[0].Data) == 0 {
		return "", errors.New("secret not found")
	}
	return string(results[0].Data), nil
}
