package cmd

import (
	"fmt"

	"setenv/tui/bubble"

	tea "charm.land/bubbletea/v2"
	"github.com/spf13/cobra"
)

func NewToolsCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "tools",
		Short: "Interactively select tools",
		RunE: func(cmd *cobra.Command, args []string) error {
			p := tea.NewProgram(
				bubble.InitialModel(),
			)

			finalModel, err := p.Run()
			if err != nil {
				return fmt.Errorf("TUI failed: %w", err)
			}

			m, ok := finalModel.(bubble.Mo)
			if !ok {
				return fmt.Errorf("unexpected model type after TUI exit")
			}

			var selected []string
			for idx := range m.Selected {
				selected = append(selected, m.Choices[idx])
			}

			if len(selected) == 0 {
				fmt.Println("\nNo tools selected.")
				return nil
			}

			fmt.Println("\nSelected tools:")
			for _, item := range selected {
				fmt.Printf("  • %s\n", item)
			}

			for _, v := range selected {
				checkTool(v)
			}

			return nil
		},
	}
}

func init() {
	rootCmd.AddCommand(NewToolsCommand())
}
