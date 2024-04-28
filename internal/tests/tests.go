package tests

type TestVariantResult struct {
	Err error
	Log []byte
	KV  map[string]string
}

type TestVariant struct {
	Name        string
	Args        []string
	Result      TestVariantResult
	SuccessFunc func(TestVariant) bool
}

func (tv TestVariant) Success() bool {
	if tv.SuccessFunc != nil {
		return tv.SuccessFunc(tv)
	}
	if tv.Result.Err != nil {
		return false
	}
	return true
}

// FQN: fully qualified `Test` name; if set, it is a slice containing sequence of: `TestFamily.Name`, `TestGroup.Name`, `Test.Name` and `TestVariant.Name`.
// Name: `Test` name.
// Cmd: `Test` command, used to execute/run all `Test`'s `TestVariant`s.
// Variants: `TestVariant` slice.
// `Test` has two, composite states:
// - split: meaning it holds only a single `TestVariant`, and `FQN` is set properly.
// - primed: meaning it holds it's split and `FQN` is set to `[]string{TestGroup.Name, Test.Name, Test.Variants[0].Name}`
type Test struct {
	FQN      []string
	Name     string
	Cmd      string
	Variants []TestVariant
}

func (t Test) Run() error    { return nil }
func (t Test) Success() bool { return true }

// SplitNextVariant splits `Test` holding multiple `TestVariant`s, into separate `Test`.
func (t *Test) SplitNextVariant() (Test, error) {
	v := t.Variants[0]
	t.Variants = t.Variants[1:]
	return Test{Variants: []TestVariant{v}}, nil
}

// IsSplit checks if `Test` holds only one `TestVariant`.
func (t Test) IsSplit() bool {
	return len(t.Variants) == 1
}

// IsPrimed checks if `Test` is ready for execution.
func (t Test) IsPrimed() bool {
	if t.IsSplit() && (len(t.FQN) == 3 && len(t.FQN[0]) != 0 && len(t.FQN[1]) != 0 && len(t.FQN[2]) != 0) {
		return true
	}
	return false
}

type TestGroup struct {
	Name  string
	Tests []Test
}

type TestFamily struct {
	Name   string
	Groups []TestGroup
}

func AllTests() []TestFamily {
	ts := []TestFamily{
		{
			Name: "Network",
			Groups: []TestGroup{
				{
					Name: "DNS",
					Tests: []Test{
						{
							Name: "Resolve A record",
							Cmd:  "/usr/bin/dig",
							Variants: []TestVariant{
								{
									Name: "... of something.example.com",
									Args: []string{"A", "something.example.com"},
								},
								{
									Name: "... of xyz.example.com",
									Args: []string{"A", "xyz.example.com"},
								},
								{
									Name: "... of abc.example.com",
									Args: []string{"A", "abc.example.com"},
								},
							},
						},
					},
				},
			},
		},
	}
	return ts
}

func LookupTest(tgs []TestGroup) Test {
	return Test{}
}

func FlattenTestGroup(tgs []TestGroup) []Test {
	return []Test{}
}
