package interfaces

type TestVariantResult interface {
	Success() bool
}

type TestVariant interface{}

type Test interface {
	Run() error
	SplitNextVariant() (Test, error)
}
