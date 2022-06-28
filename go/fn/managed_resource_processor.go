package fn

type ManagedResourceProcessor interface {
	Process(mr *KubeObject) (bool, error)
}

// ManagedResourceProcessorFunc converts a compatible function to a ManagedResourceProcessor.
type ManagedResourceProcessorFunc func(mr *KubeObject) (bool, error)

func (p ManagedResourceProcessorFunc) Process(mr *KubeObject) (bool, error) {
	return p(mr)
}
