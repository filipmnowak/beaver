package dns

import (
	. "codeberg.org/filipmnowak/beaver/internal/tests/interfaces"
)

type DNSTestVariantResult struct {
	Err         error
	Log         []byte
	KV          map[string]string
	SuccessFunc func(TestVariantResult) bool
}

func (dtvr DNSTestVariantResult) Success() bool {
	if dtvr.SuccessFunc != nil {
		return dtvr.SuccessFunc(dtvr)
	}
	if dtvr.Err != nil {
		return false
	}
	return true
}

type DNSTestVariant struct {
	Name      string
	Arguments map[string]any
	Result    DNSTestVariantResult
}

type DNSTest struct {
	Name     string
	Cmd      string
	Variants []DNSTestVariant
}

func (dnst DNSTest) Run() error    { return nil }
func (dnst DNSTest) Success() bool { return true }
func (dnst *DNSTest) SplitNextVariant() (Test, error) {
	v := dnst.Variants[0]
	dnst.Variants = dnst.Variants[1:]
	return &DNSTest{Variants: []DNSTestVariant{v}}, nil
}

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
