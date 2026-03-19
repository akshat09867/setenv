package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func selectRole(role string) {
	roles := strings.ToLower(role)
	allTools := []string{"Git", "Docker", "SSH"}
	var recommendedTools []string
	switch roles {
	case "Developer":
		recommendedTools = []string{"Git", "Docker", "SSH"}
	case "Designer":
		recommendedTools = []string{"Git"}
	case "SysAdmin":
		recommendedTools = []string{"Docker", "SSH"}
	default:
		recommendedTools = []string{"Git", "Docker", "SSH"}
	}
	defaultIndexes := []int{}
	for _, def := range recommendedTools {
		for i, tool := range allTools {
			if tool == def {
				defaultIndexes = append(defaultIndexes, i)
				break
			}
		}
	}
}

var selectCmd = &cobra.Command{
	Use:   "role",
	Short: "Recommending tools on basis of your selected role",
	Run: func(cmd *cobra.Command, args []string) {
		ro, _ := cmd.Flags().GetString("ro")
		if ro == "" {
			fmt.Println("Please specify a role with --ro or -r")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)
	selectCmd.Flags().StringP("ro", "r", "", "command to select role")
}
