package cmd

import (
	"codeberg.org/filipmnowak/beaver/cmd/db"
	"codeberg.org/filipmnowak/beaver/cmd/test"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "beaver",
	Short: "Test K8s and cloud infra",
}

func init() {
	rootCmd.AddCommand(db.NewDBCmd())
	rootCmd.AddCommand(test.NewTestCmd())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
