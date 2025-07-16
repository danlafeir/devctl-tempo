package secrets

import "testing"

func TestFileSecretStore(t *testing.T) {
	store := NewFileSecretStore()
	key := "api_token"
	value := "supersecret"

	if err := store.WriteSecret(key, value); err != nil {
		t.Fatalf("WriteSecret failed: %v", err)
	}

	v, err := store.ReadSecret(key)
	if err != nil {
		t.Fatalf("ReadSecret failed: %v", err)
	}
	if v != value {
		t.Errorf("Expected %q, got %q", value, v)
	}

	_, err = store.ReadSecret("missing")
	if err == nil {
		t.Errorf("Expected error for missing secret, got nil")
	}
}
