package test

import (
	. "codeberg.org/filipmnowak/beaver/internal/tests"
	. "codeberg.org/filipmnowak/beaver/internal/tests/runner"
	"github.com/spf13/cobra"
)

func NewTestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "infra testing-related commands",
	}
	cmd.AddCommand(NewTestRunCmd())
	return cmd
}

func NewTestRunCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run infra tests.",
		Run: func(cmd *cobra.Command, _ []string) {
			TestRun(cmd)
		},
	}

	//cmd.Flags().String("db-path", "data/beaver.sqlite3", "optional; filesystem path to SQLite DB")
	return cmd
}

func TestRun(cmd *cobra.Command) {
	tests := FlattenTests(AllTests())
	chs := RunTests(tests)
	ch := Merge(chs)
	PersistResults(ch)
}
