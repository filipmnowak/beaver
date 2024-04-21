package azure

import (
	"codeberg.org/filipmnowak/beaver/internal/tests/groups/azure/acr"
	. "codeberg.org/filipmnowak/beaver/internal/tests/interfaces"
)

type AzureTestGroup struct {
	name  string
	tests any
}

func (azrg AzureTestGroup) Name() string       { return azrg.name }
func (azrg AzureTestGroup) Tests() any         { return acr.ACRTest{} }
func (azrg AzureTestGroup) TestResult(any) any { return "" }

func NewAzureTestGroup() AzureTestGroup {
	return AzureTestGroup{
		name:  "Azure Test Group",
		tests: acr.AllACRTests(),
	}
}
