package dns

import (
	. "codeberg.org/filipmnowak/beaver/internal/tests/groups/network/interfaces"
)

type DNSTestVariantResult struct {
	Success bool
	Err     error
	Log     []byte
	KV      map[string]string
}

type DNSTestVariant struct {
	Name      string
	Arguments map[string]any
	Result    DNSTestVariantResult
	Expected  string
}

type DNSTest struct {
	Name     string
	Cmd      string
	Variants []DNSTestVariant
}

func (dnst DNSTest) Run() error                      { return nil }
func (dnst DNSTest) Success() bool                   { return true }
func (dnst DNSTest) SplitNext() (NetworkTest, error) { return DNSTest{}, nil }
func (dnst DNSTest) Merge(NetworkTest) error         { return nil }
func (dnst DNSTest) MergeAll([]NetworkTest) error    { return nil }

func AllDNSTests() []DNSTest {
	return []DNSTest{
		{
			Name: "Resolve A record",
			Cmd:  "/usr/bin/dig",
			Variants: []DNSTestVariant{
				{
					Name: "... of something.example.com",
					Arguments: map[string]any{
						"args": []string{"A", "something.example.com"},
					},
				},
				{
					Name: "... of xyz.example.com",
					Arguments: map[string]any{
						"args": []string{"A", "xyz.example.com"},
					},
				},
				{
					Name: "... of abc.example.com",
					Arguments: map[string]any{
						"args": []string{"A", "abc.example.com"},
					},
				},
			},
		},
		{
			Name: "Resolve AAAA record",
			Cmd:  "/usr/bin/dig",
			Variants: []DNSTestVariant{
				{
					Name: "... of something.example.com",
					Arguments: map[string]any{
						"args": []string{"AAAA", "something.example.com"},
					},
				},
				{
					Name: "... of xyz.example.com",
					Arguments: map[string]any{
						"args": []string{"AAAA", "xyz.example.com"},
					},
				},
				{
					Name: "... of abc.example.com",
					Arguments: map[string]any{
						"args": []string{"AAAA", "abc.example.com"},
					},
				},
			},
		},
		{
			Name: "Resolve MX record",
			Cmd:  "/usr/bin/dig",
			Variants: []DNSTestVariant{
				{
					Name: "... of something.example.com",
					Arguments: map[string]any{
						"args": []string{"MX", "something.example.com"},
					},
				},
				{
					Name: "... of xyz.example.com",
					Arguments: map[string]any{
						"args": []string{"MX", "xyz.example.com"},
					},
				},
				{
					Name: "... of abc.example.com",
					Arguments: map[string]any{
						"args": []string{"MX", "abc.example.com"},
					},
				},
			},
		},
	}
}
