# devctl-tempo

A Go CLI for interacting with the Tempo API, using Cobra and Viper. Stores secrets securely on MacOS using the system keychain.

## Prerequisites
- Go 1.18+
- (For MacOS secrets) Keychain access

## Setup
Clone the repository and install dependencies:
```sh
git clone <repo-url>
cd devctl-tempo
go mod tidy
```

## Testing
Run all tests:
```sh
make test
```

## Building
Build for your current OS and architecture:
```sh
make build
```
The binary will be output to `bin/devctl-tempo`.

Build for all supported OS/arch:
```sh
make build-all
```
Binaries for Linux, MacOS, and Windows (amd64/arm64) will be in `bin/`.

## Running
After building, run the CLI:
```sh
./bin/devctl-tempo <command>
```

## Configuration
To configure your Tempo API token and default issue ID:
```sh
./bin/devctl-tempo configure --api-token <YOUR_TOKEN> --issue-id <ISSUE_ID>
```
You can also omit flags to be prompted interactively.

- The API token is stored securely in the MacOS keychain.
- The default issue ID is stored in `$HOME/.devctl/plugins/tempo/config.yaml`.

## Available Commands
- `configure` — Set up your Tempo API token and default issue ID.
- (Hidden) `get-week` — Fetch your current week's timecard from the Tempo API.
- (Hidden) `completion` — Generate shell completion scripts.

## Notes
- Only MacOS is currently supported for secure secrets storage.
- For other OS support, contributions are welcome!

---

For more information, see the code and comments or open an issue.
