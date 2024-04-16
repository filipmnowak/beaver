package interfaces

type NetworkTestVariantResult interface{}

type NetworkTestVariant interface{}

type NetworkTest interface {
	Run() error
	Success() bool
	SplitNext() (NetworkTest, error)
	Merge(NetworkTest) error
	MergeAll([]NetworkTest) error
}
