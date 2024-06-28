package test

import (
	"net"
	"time"

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

	defaultTestTimeout, _ := time.ParseDuration("15s")

	cmd.PersistentFlags().String("db-path", "data/beaver.sqlite3", "filesystem path to SQLite DB")
	cmd.PersistentFlags().DurationP("test-timeout", "t", defaultTestTimeout, "timeout set globally for all of the tests")
	cmd.PersistentFlags().UintP("test-buffer-size", "u", 32, "test results channel buffer size")
	cmd.PersistentFlags().UintP("test-batch-size", "a", 16, "maximum tests executed at the same time")
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

	cmd.Flags().Uint32P("port", "p", 8080, "dashboard listening TCP port")
	cmd.Flags().IPP("ip", "i", net.IP{0, 0, 0, 0}, "dashboard listening IPv4 address")

	return cmd
}

func TestRunForeverAndServe(cmd *cobra.Command) {
	dbPath, _ := cmd.Flags().GetString("db-path")
	tests := FlattenTests(AllTests())
	chs := RunTests(tests)
	ch := Merge(chs)
	PersistResults(ch, dbPath)
}
