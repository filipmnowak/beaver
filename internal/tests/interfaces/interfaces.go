package interfaces

type TestVariantResult interface{}

type TestVariant interface{}

type Test interface {
	Run() error
	Success() bool
	SplitNext() (Test, error)
	Merge(Test) error
	MergeAll([]Test) error
}
