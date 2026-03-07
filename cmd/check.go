package cmd

import (
	"context"
	"fmt"
	"setenv/checks"
	"strings"

	"github.com/spf13/cobra"
)

var (
	flag bool
)

func checkTool(tool string) {
	ctx := context.Background()
	tool = strings.ToLower(tool)
	flag = true
	switch tool {
	case "git":
		out, err := checks.Checkgit(ctx)
		if err != nil {
			fmt.Printf("✘ git: %v\n", err)
			fix, err2 := checks.Fixgit(ctx, false)
			if err2 != nil {
				fmt.Printf("error during fixing: %v\n", err2)
				flag = false
			}
			if flag {
				fmt.Printf("fixed %v", fix)
			}
		} else {
			fmt.Printf("✔ git: %s\n", out)
		}
	case "docker":
		out, err := checks.Checkdocker(ctx)
		if err != nil {
			fix, err2 := checks.Fixdocker(ctx, false)
			if err2 != nil {
				fmt.Printf("error during fixing: %v\n", err2)
				flag = false
			}
			if flag {
				fmt.Printf("fixed %v", fix)
			}
		} else {
			fmt.Printf("✔ docker: %s\n", out)
		}
	case "ssh":
		out, err := checks.Checkssh(ctx)
		if err != nil {
			fix, err2 := checks.Fixssh(ctx, false)
			if err2 != nil {
				fmt.Printf("error during fixing: %v\n", err2)
				flag = false
			}
			if flag {
				fmt.Printf("fixed %v", fix)
			}
		} else {
			fmt.Printf("✔ ssh: %s\n", out)
		}
	default:
		fmt.Printf("Unknown tool: %s\n", tool)
	}
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Checks whether tool is installed or not",
	Run: func(cmd *cobra.Command, args []string) {
		tool, _ := cmd.Flags().GetString("tool")
		if tool == "" {
			fmt.Println("Please specify a tool with --tool or -c")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().StringP("tool", "c", "", "command to check installation")
}
