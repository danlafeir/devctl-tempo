//go:build darwin

package secrets

import (
	"os"
	"testing"
)

func TestKeychainSecretStore(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping keychain test in CI environment")
	}
	store := NewKeychainSecretStore()
	key := "api_token"
	value := "supersecret"

	err := store.WriteSecret(key, value)
	if err != nil {
		t.Fatalf("WriteSecret failed: %v", err)
	}

	v, err := store.ReadSecret(key)
	if err != nil {
		t.Fatalf("ReadSecret failed: %v", err)
	}
	if v != value {
		t.Errorf("Expected %q, got %q", value, v)
	}
}
