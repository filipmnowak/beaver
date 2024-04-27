package azure

import (
	"codeberg.org/filipmnowak/beaver/internal/tests/groups/azure/acr"
	. "codeberg.org/filipmnowak/beaver/internal/tests/interfaces"
)

type TestGroup struct {
	Name  string
	Tests []Test
}

func AzureTestGroup() TestGroup {
	// https://go.dev/wiki/InterfaceSlice
	acr_tests := acr.AllACRTests()

	ts := make([]Test, len(acr_tests))
	for i, d := range acr_tests {
		ts[i] = &d
	}
	return TestGroup{
		Name:  "Network Test Group",
		Tests: ts,
	}
}
