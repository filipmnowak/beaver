package network

import (
	"codeberg.org/filipmnowak/beaver/internal/tests/groups/network/dns"
	"codeberg.org/filipmnowak/beaver/internal/tests/groups/network/http"
	. "codeberg.org/filipmnowak/beaver/internal/tests/interfaces"
)

type TestGroup struct {
	Name  string
	Tests []Test
}

func NetworkTestGroup() TestGroup {
	// https://go.dev/wiki/InterfaceSlice
	dns_tests := dns.AllDNSTests()
	http_tests := http.AllHTTPTests()

	ts := make([]Test, len(dns_tests)+len(http_tests))
	for i, d := range dns_tests {
		ts[i] = &d
	}
	for i, d := range http_tests {
		ts[i+len(dns_tests)] = &d
	}
	return TestGroup{
		Name:  "Network Test Group",
		Tests: ts,
	}
}
