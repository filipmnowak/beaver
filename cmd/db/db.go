package db

import (
	"fmt"
	"os"

	"codeberg.org/filipmnowak/beaver/internal/db/sqlite"
	"codeberg.org/filipmnowak/beaver/internal/tests"
	"github.com/spf13/cobra"
)

func NewDBCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "db",
		Short: "DB-related commands",
	}
	cmd.AddCommand(NewDBInitCmd())
	return cmd
}

func NewDBInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initalize database",
		Run: func(cmd *cobra.Command, _ []string) {
			DBInit(cmd)
		},
	}

	cmd.Flags().String("db-path", "data/beaver.sqlite3", "optional; filesystem path to SQLite DB")
	return cmd
}

func DBInit(cmd *cobra.Command) {
	dbPath, _ := cmd.Flags().GetString("db-path")
	db := sqlite.NewDB(nil, dbPath, "")
	if !db.Init() {
		for _, e := range db.InitErrors {
			fmt.Println(e)
		}
		os.Exit(1)
	}
}

func TestDummy(cmd *cobra.Command) {
	dbPath, _ := cmd.Flags().GetString("db-path")
	db := sqlite.NewDB(nil, dbPath, "")
	db.Init()
	fmt.Printf("%s\n", db.InitErrors)
	input := []map[string]string{
		{"family": "network", "_group": "DNS", "test": "resolve A record", "variant": "of something1.org", "key": "/key1", "value": "/abc/def/value/001"},
		{"family": "network", "_group": "DNS", "test": "resolve A record", "variant": "of something1.org", "key": "/key2", "value": "/abc/def/value/00x"},
		{"family": "network", "_group": "DNS", "test": "resolve A record", "variant": "of something1.org", "key": "/key3", "value": "/abc/def/value/001"},
		{"family": "network", "_group": "DNS", "test": "resolve A record", "variant": "of something2.org", "key": "/key4", "value": "/abc/def/value/001"},
		{"family": "network", "_group": "DNS", "test": "resolve A record", "variant": "of something3.org", "key": "/key5", "value": "/abc/def/value/001"},
		{"family": "network", "_group": "DNS", "test": "resolve AAAA record", "variant": "of something4.org", "key": "/key6", "value": "/abc/def/value/001"},
		{"family": "network", "_group": "DNS", "test": "resolve AAAA record", "variant": "of something5.org", "key": "/key7", "value": "/abc/def/value/001"},
	}
	out, err := db.TransactUpserts(input, "test_results", "family, _group, test, variant, key")
	fmt.Printf("out: \n%s\nerr:\n%s\n", out, err)
	flattenedTests := tests.FlattenTests(tests.AllTests())
	for _, t := range flattenedTests {
		fmt.Printf("%s\n", t.Variants)
	}
}
