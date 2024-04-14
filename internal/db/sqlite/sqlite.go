package sqlite

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

func dbFileSuffix() string {
	return time.Now().Format(time.RFC3339Nano)
}

type DBInitErr struct {
	Level string
	Error error
}

type DB struct {
	DBPath      string
	InitSQLFunc func() string
	SQLITE3Cmd  string
	InitErrors  []DBInitErr
}

func (dbi DB) Success() bool {
	for _, e := range dbi.InitErrors {
		if e.Level == "error" {
			return false
		}
	}
	return true
}

func (dbi *DB) Init() bool {
	if err := os.MkdirAll(filepath.Dir(dbi.DBPath), 0700); err != nil && !os.IsNotExist(err) {
		dbi.InitErrors = append(dbi.InitErrors, DBInitErr{"error", err})
	}
	if err := os.Rename(dbi.DBPath, dbi.DBPath+"_"+dbFileSuffix()); err != nil && !os.IsNotExist(err) {
		dbi.InitErrors = append(dbi.InitErrors, DBInitErr{"warning", err})
	}
	cmd := exec.Command(dbi.SQLITE3Cmd, dbi.DBPath, dbi.InitSQLFunc())
	if err := cmd.Run(); err != nil {
		dbi.InitErrors = append(dbi.InitErrors, DBInitErr{"error", err})
	}
	return dbi.Success()
}

func TemplateUpserts(values []map[string]string, table, conflictOn string) (string, error) {
	ConflictOn := conflictOn
	if conflictOn == "" {
		// in most of the cases tables will have two columns: key and value.
		ConflictOn = "key"
	}

	cols := make([][]string, len(values))
	vals := make([][]string, len(values))
	for i, vs := range values {
		for k, v := range vs {
			cols[i] = append(cols[i], k)
			vals[i] = append(vals[i], "'"+v+"'")
		}
	}
	_cols := make([]string, len(values))
	_vals := make([]string, len(values))
	for i := range cols {
		_cols[i] = strings.Join(cols[i], ",")
		_vals[i] = strings.Join(vals[i], ",")
	}

	t, err := template.New("t").Parse(
		`{{range $i1, $v := .Cols}}
		INSERT INTO {{$.Table}}({{index $._Cols $i1}}) VALUES({{index $._Vals $i1}})
		ON CONFLICT({{$.ConflictOn}}) DO UPDATE SET ({{index $._Cols $i1}}) = ({{index $._Vals $i1}});
		{{end}}
		`,
	)
	if err != nil {
		return "", err
	}
	buffer := &bytes.Buffer{}
	input := map[string]any{
		"Table":      table,
		"ConflictOn": ConflictOn,
		"Cols":       cols,
		"_Cols":      _cols,
		"_Vals":      _vals,
	}
	err = t.Execute(buffer, input)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func (dbi DB) TransactUpserts(values []map[string]string, table, conflictOn string) (string, error) {
	upserts, err := TemplateUpserts(values, table, conflictOn)
	if err != nil {
		return "", err
	}
	transaction := "BEGIN TRANSACTION;\n" + upserts + "COMMIT;\n"
	var output string
	if output, err = dbi.RunStatement(transaction, true, false, false); err != nil {
		return output, err
	}
	return output, nil
}

func (dbi DB) RunStatement(statement string, rw, unsafe, noJSONOutput bool) (string, error) {
	flags := []string{
		dbi.DBPath,
		statement,
	}
	if !rw {
		flags = append(flags, "--readonly")
	}
	if !unsafe {
		flags = append(flags, "--safe")
	}
	if !noJSONOutput {
		flags = append(flags, "--json")
	}
	cmd := exec.Command(dbi.SQLITE3Cmd, flags...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output[:]), err
	}
	return string(output[:]), nil
}

func NewDB(InitSQLFunc func() string, DBPath, SQLITE3Cmd string) DB {
	sqlite3Cmd := "/usr/bin/sqlite3"
	initSQLFunc := func() string {
		create_table := `
		CREATE TABLE test_results(
			_group STRING,
			test STRING,
			variant STRING,
			key STRING,
			value STRING,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (_group, test, variant, key));`
		return create_table
	}
	if InitSQLFunc != nil {
		initSQLFunc = InitSQLFunc
	}
	if SQLITE3Cmd != "" {
		sqlite3Cmd = SQLITE3Cmd
	}
	return DB{
		DBPath:      DBPath,
		SQLITE3Cmd:  sqlite3Cmd,
		InitSQLFunc: initSQLFunc,
	}
}
