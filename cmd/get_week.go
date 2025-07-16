package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var getWeekCmd = &cobra.Command{
	Hidden: true,
	Use:   "get-week",
	Short: "Fetch your current week's timecard from Tempo API",
	Run: func(cmd *cobra.Command, args []string) {
		apiToken := os.Getenv("TEMPO_API_TOKEN")
		if apiToken == "" {
			fmt.Println("TEMPO_API_TOKEN environment variable is not set.")
			os.Exit(1)
		}

		// Calculate current week start and end (Monday-Sunday)
		now := time.Now()
		weekday := int(now.Weekday())
		if weekday == 0 { // Sunday
			weekday = 7
		}
		monday := now.AddDate(0, 0, -weekday+1)
		sunday := monday.AddDate(0, 0, 6)
		startDate := monday.Format("2006-01-02")
		endDate := sunday.Format("2006-01-02")

		url := fmt.Sprintf("https://api.tempo.io/core/3/worklogs/user/me?from=%s&to=%s", startDate, endDate)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Printf("Failed to create request: %v\n", err)
			os.Exit(1)
		}
		req.Header.Set("Authorization", "Bearer "+apiToken)
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("Request failed: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			fmt.Printf("API returned status: %s\n", resp.Status)
			os.Exit(1)
		}

		var result map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			fmt.Printf("Failed to decode response: %v\n", err)
			os.Exit(1)
		}

		pretty, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(pretty))
	},
}

func init() {
	rootCmd.AddCommand(getWeekCmd)
}
