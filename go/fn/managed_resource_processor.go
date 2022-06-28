package fn

type ManagedResourceProcessor interface {
	Process(mr *KubeObject) (*KubeObject, error)
}

// ManagedResourceProcessorFunc converts a compatible function to a ManagedResourceProcessor.
type ManagedResourceProcessorFunc func(mr *KubeObject) (*KubeObject, error)

func (p ManagedResourceProcessorFunc) Process(mr *KubeObject) (*KubeObject, error) {
	return p(mr)
}
