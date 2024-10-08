package tests

import (
	"os/exec"
	"slices"
)

type TestVariantResult struct {
	Err error
	Log []byte
}

type TestVariant struct {
	Name        string
	Args        []string
	Result      TestVariantResult
	SuccessFunc func(TestVariant) bool
	ErrorFunc func(TestVariant) string
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

func (tv TestVariant) Error() string {
	if tv.SuccessFunc != nil {
		return tv.ErrorFunc(tv)
	}
	if tv.Result.Err != nil {
		return tv.Result.Err.Error()
	}
	return ""
}

// FQN: fully qualified `Test` name; if set, it is a slice containing sequence of: `TestFamily.Name`, `TestGroup.Name`, `Test.Name` and `TestVariant.Name`.
// Name: `Test` name.
// Cmd: `Test` command, used to execute/run all `Test` `TestVariant`s.
// Variants: `TestVariant` slice.
// `Test` has two, implicit states:
// - split: meaning it holds only a single `TestVariant`.
// - primed: `Test` is split and its `FQN` is set to `[]string{TestFamily.Name, TestGroup.Name, Test.Name, Test.Variants[0].Name}`
type Test struct {
	FQN      []string
	Name     string
	Cmd      string
	Variants []TestVariant
}

func (t Test) Run() error {
	cmd := exec.Command(t.Cmd, t.Variants[0].Args...)
	output, err := cmd.CombinedOutput()
	t.Variants[0].Result.Err = err
	t.Variants[0].Result.Log = output

	return nil
}
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

// PrimedMaybe checks if `Test` is ready for execution.
func (t Test) PrimedMaybe() bool {
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
									Args: []string{"+short", "A", "something.example.com"},
								},
								{
									Name: "... of xyz.example.com",
									Args: []string{"+short", "A", "xyz.example.com"},
								},
								{
									Name: "... of abc.example.com",
									Args: []string{"+short", "A", "abc.example.com"},
								},
							},
						},
					},
				},
				{
					Name: "Time",
					Tests: []Test{
						{
							Name: "Print date",
							Cmd:  "/usr/bin/date",
							Variants: []TestVariant{
								{
									Name: "... in seconds",
									Args: []string{"+%s"},
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

func LookupTest(fqn [4]string, tf []TestFamily) *Test {
	i1 := slices.IndexFunc(tf, func(tf TestFamily) bool {
		return tf.Name == fqn[0]
	})
	if i1 == -1 {
		return &Test{}
	}
	i2 := slices.IndexFunc(tf[i1].Groups, func(tg TestGroup) bool {
		return tg.Name == fqn[1]
	})
	if i2 == -1 {
		return &Test{}
	}
	i3 := slices.IndexFunc(tf[i1].Groups[i2].Tests, func(t Test) bool {
		return t.Name == fqn[2]
	})
	if i3 == -1 {
		return &Test{}
	}
	return &tf[i1].Groups[i2].Tests[i3]
}

func FlattenTests(tfs []TestFamily) []*Test {
	tests := []*Test{}
	for i1, tf := range tfs {
		for i2, tg := range tf.Groups {
			for i3, t := range tg.Tests {
				for i4, v := range t.Variants {
					_t := Test{
						FQN:      []string{tf.Name, tg.Name, t.Name, v.Name},
						Cmd:      tfs[i1].Groups[i2].Tests[i3].Cmd,
						Name:     tfs[i1].Groups[i2].Tests[i3].Name,
						Variants: []TestVariant{tfs[i1].Groups[i2].Tests[i3].Variants[i4]},
					}
					tests = append(tests, &_t)
				}
			}
		}
	}
	return tests
}
