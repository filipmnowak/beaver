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
	cmd.AddCommand(NewTestRunOnceCmd())
	cmd.AddCommand(NewTestRunForeverCmd())
	cmd.AddCommand(NewTestRunForeverAndServeCmd())
	return cmd
}

func NewTestRunOnceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run-once",
		Short: "Run infra tests once and exit.",
		Run: func(cmd *cobra.Command, _ []string) {
			TestRunOnce(cmd)
		},
	}

	cmd.Flags().String("db-path", "data/beaver.sqlite3", "optional; filesystem path to SQLite DB")
	return cmd
}

func TestRunOnce(cmd *cobra.Command) {
	dbPath, _ := cmd.Flags().GetString("db-path")
	tests := FlattenTests(AllTests())
	chs := RunTests(tests)
	ch := Merge(chs)
	PersistResults(ch, dbPath)
}

func NewTestRunForeverCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run-forever",
		Short: "Run infra tests forever.",
		Run: func(cmd *cobra.Command, _ []string) {
			TestRunForever(cmd)
		},
	}

	cmd.Flags().String("db-path", "data/beaver.sqlite3", "optional; filesystem path to SQLite DB")
	return cmd
}

func TestRunForever(cmd *cobra.Command) {
	dbPath, _ := cmd.Flags().GetString("db-path")
	tests := FlattenTests(AllTests())
	chs := RunTests(tests)
	ch := Merge(chs)
	PersistResults(ch, dbPath)
}

func NewTestRunForeverAndServeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run-forever-and-serve",
		Short: "Run infra tests forever and serve results dashboard.",
		Run: func(cmd *cobra.Command, _ []string) {
			TestRunForeverAndServe(cmd)
		},
	}

	cmd.Flags().String("db-path", "data/beaver.sqlite3", "optional; filesystem path to SQLite DB")
	return cmd
}

func TestRunForeverAndServe(cmd *cobra.Command) {
	dbPath, _ := cmd.Flags().GetString("db-path")
	tests := FlattenTests(AllTests())
	chs := RunTests(tests)
	ch := Merge(chs)
	PersistResults(ch, dbPath)
}
