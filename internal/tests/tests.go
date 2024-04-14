package tests

import (
	"codeberg.org/filipmnowak/beaver/internal/tests/groups"
)

func AllTests() any {
	return groups.NewTestGroups()
}
