package network

import (
	"codeberg.org/filipmnowak/beaver/internal/tests/groups/network/dns"
	"codeberg.org/filipmnowak/beaver/internal/tests/groups/network/http"
)

type NetworkTests interface {
	Run() error
	Success() bool
}

type NetworkTestGroup struct {
	Name  string
	Tests []NetworkTests
}

func NewNetworkTestGroup() NetworkTestGroup {
	// https://go.dev/wiki/InterfaceSlice
	dns_tests := dns.AllDNSTests()
	http_tests := http.AllHTTPTests()

	ts := make([]NetworkTests, len(dns_tests)+len(http_tests))
	for i, d := range dns_tests {
		ts[i] = d
	}
	for i, d := range http_tests {
		ts[i+len(dns_tests)] = d
	}
	return NetworkTestGroup{
		Name:  "Network Test Group",
		Tests: ts,
	}
}
