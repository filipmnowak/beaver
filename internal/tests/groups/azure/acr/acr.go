package acr

import (
	. "codeberg.org/filipmnowak/beaver/internal/tests/groups/network/interfaces"
)

type ACRTestVariantResult struct {
	Success bool
	Err     error
	Log     []string
	KV      map[string]any
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

func (dnst ACRTest) Run() error                      { return nil }
func (dnst ACRTest) Success() bool                   { return true }
func (dnst ACRTest) SplitNext() (NetworkTest, error) { return ACRTest{}, nil }
func (dnst ACRTest) Merge(NetworkTest) error         { return nil }
func (dnst ACRTest) MergeAll([]NetworkTest) error    { return nil }

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
