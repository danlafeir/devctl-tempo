package secrets

import (
	"errors"
	"fmt"
	"runtime"
)

// NewSecretStore returns the appropriate SecretStore for the current OS.
func NewSecretStore() (SecretStore, error) {
	switch runtime.GOOS {
	case "darwin":
		return NewKeychainSecretStore(), nil
	default:
		err := errors.New("no supported secrets backend for this OS")
		fmt.Println("[secrets] ERROR:", err)
		return nil, err
	}
}
