package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/danlafeir/devctl-tempo/pkg/secrets"
	"github.com/danlafeir/devctl-tempo/pkg/config"
)

type Config struct {
	TempoAPIToken string `json:"tempo_api_token"`
}

var apiToken string
var issueId string

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure your Tempo API token and default issue ID",
	Run: func(cmd *cobra.Command, args []string) {
		// API Token logic
		token := strings.TrimSpace(apiToken)
		if token == "" {
			fmt.Print("Enter your Tempo API token: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			token = strings.TrimSpace(scanner.Text())
		}
		if token == "" {
			fmt.Println("Token cannot be empty.")
			os.Exit(1)
		}
		store := secrets.NewKeychainSecretStore()
		if err := store.WriteTempoAPIToken(token); err != nil {
			fmt.Println("Failed to write token to keychain:", err)
			os.Exit(1)
		}
		fmt.Println("Tempo API token saved securely to keychain.")

		// Issue ID logic
		id := strings.TrimSpace(issueId)
		if id == "" {
			fmt.Print("Enter your default Issue ID (alphanumeric): ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			id = strings.TrimSpace(scanner.Text())
		}
		if id == "" {
			fmt.Println("Issue ID cannot be empty.")
			os.Exit(1)
		}
		// Validate alphanumeric
		for _, c := range id {
			if !(('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9')) {
				fmt.Println("Issue ID must be alphanumeric.")
				os.Exit(1)
			}
		}
		if err := config.InitConfig(); err != nil {
			fmt.Println("Failed to initialize config:", err)
			os.Exit(1)
		}
		viper.Set("issue_id", id)
		if err := config.SaveConfig(); err != nil {
			fmt.Println("Failed to save config:", err)
			os.Exit(1)
		}
		fmt.Println("Default issue ID saved to config.")
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
	configureCmd.Flags().StringVar(&apiToken, "api-token", "", "Tempo API token to store in the keychain")
	configureCmd.Flags().StringVar(&issueId, "issue-id", "", "Default issue ID (alphanumeric) to store in config")
}

