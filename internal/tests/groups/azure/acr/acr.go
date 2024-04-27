package acr

import (
	. "codeberg.org/filipmnowak/beaver/internal/tests/interfaces"
)

type ACRTestVariantResult struct {
	Err         error
	Log         []byte
	KV          map[string]string
	SuccessFunc func(TestVariantResult) bool
}

func (acrtv ACRTestVariantResult) Success() bool {
	if acrtv.SuccessFunc != nil {
		return acrtv.SuccessFunc(acrtv)
	}
	if acrtv.Err != nil {
		return false
	}
	return true
}

type ACRTestVariant struct {
	Name      string
	Arguments map[string]any
	Result    ACRTestVariantResult
	Expected  string
}

type ACRTest struct {
	Name        string
	Cmd         string
	SuccessFunc func() bool
	Variants    []ACRTestVariant
}

func (dnst ACRTest) Run() error    { return nil }
func (dnst ACRTest) Success() bool { return true }
func (dnst *ACRTest) SplitNextVariant() (Test, error) {
	v := dnst.Variants[0]
	dnst.Variants = dnst.Variants[1:]
	return &ACRTest{Variants: []ACRTestVariant{v}}, nil
}

func AllACRTests() []ACRTest {
	return []ACRTest{
		{
			Name: "Something something ACR 1",
			Cmd:  "/usr/bin/true",
			Variants: []ACRTestVariant{
				{
					Name:      "... variant 1",
					Arguments: map[string]any{},
				},
				{
					Name:      "... variant 2",
					Arguments: map[string]any{},
				},
				{
					Name:      "... variant 3",
					Arguments: map[string]any{},
				},
			},
		},
		{
			Name: "Something something ACR 2",
			Cmd:  "/usr/bin/true",
			Variants: []ACRTestVariant{
				{
					Name:      "... variant 1",
					Arguments: map[string]any{},
				},
				{
					Name:      "... variant 2",
					Arguments: map[string]any{},
				},
				{
					Name:      "... variant 3",
					Arguments: map[string]any{},
				},
			},
		},
		{
			Name: "Something something ACR 3",
			Cmd:  "/usr/bin/true",
			Variants: []ACRTestVariant{
				{
					Name:      "... variant 1",
					Arguments: map[string]any{},
				},
				{
					Name:      "... variant 2",
					Arguments: map[string]any{},
				},
				{
					Name:      "... variant 3",
					Arguments: map[string]any{},
				},
			},
		},
	}
}
