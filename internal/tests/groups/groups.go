package groups

import (
	"codeberg.org/filipmnowak/beaver/internal/tests/groups/azure"
	"codeberg.org/filipmnowak/beaver/internal/tests/groups/network"
)

type TestGroup interface{}

type TestGroups struct {
	name  string
	tests TestGroup
}

func (tg TestGroups) Name() string       { return tg.name }
func (tg TestGroups) Tests() any         { return nil }

func NewTestGroups() TestGroups {
	return TestGroups{
		name: "All Test Groups",
		tests: []TestGroup{
			network.NetworkTestGroup(),
			azure.AzureTestGroup(),
		},
	}
}
