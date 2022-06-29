package fn

import (
	"fmt"

	"github.com/yndd/app-functions-sdk/go/fn/internal"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

const (
	ResourceContextKind       = "ResourceContext"
	ResourceContextAPIVersion = "app.yndd.io/v1"
)

type ResourceContext struct {
	Input   *ResourceContextInputs `yaml:"input" json:"input"`                         // the input CR(s)
	Outputs KubeObjects            `yaml:"outputs,omitempty" json:"outputs,omitempty"` // the rendered output CR
	Results Results                `yaml:"results,omitempty" json:"results,omitempty"` // result context
}

type ResourceContextInputs struct {
	Origin *KubeObject `yaml:"origin" json:"origin"`                     // the origin CR in the intent/app
	Target *KubeObject `yaml:"target,omitempty" json:"target,omitempty"` // could be node or target
	Items  KubeObjects `yaml:"items,omitempty" json:"items,omitempty"`   // additional input items like OC
}

// ParseResourceContext parses a ResourceContext from the input byte array. This function can be used to parse either KRM fn input
// or KRM fn output
func ParseResourceContext(in []byte) (*ResourceContext, error) {
	rc := &ResourceContext{}
	rcObj, err := ParseKubeObject(in)
	if err != nil {
		return nil, fmt.Errorf("failed to parse input bytes: %w", err)
	}
	if rcObj.GetKind() != ResourceContextKind {
		return nil, fmt.Errorf("input was of unexpected kind %q; expected %s",ResourceContextKind, rcObj.GetKind())
	}
	
	return rc, nil
}

// toYNode converts the ResourceList to the yaml.Node representation.
func (rl *KubeObject) toYNode() (*yaml.Node, error) {
	reMap := internal.NewMap(nil)
	//reObj := &KubeObject{SubObject{reMap}}
	//reObj.SetAPIVersion(kio.ResourceListAPIVersion)
	//reObj.SetKind(kio.ResourceListKind)

	/*
		if rl.Items != nil && len(rl.Items) > 0 {
			itemsSlice := internal.NewSliceVariant()
			for i := range rl.Items {
				itemsSlice.Add(rl.Items[i].node())
			}
			if err := reMap.SetNestedSlice(itemsSlice, "items"); err != nil {
				return nil, err
			}
		}
		if !rl.FunctionConfig.IsEmpty() {
			if err := reMap.SetNestedMap(rl.FunctionConfig.node(), "functionConfig"); err != nil {
				return nil, err
			}
		}

		if rl.Results != nil && len(rl.Results) > 0 {
			resultsSlice := internal.NewSliceVariant()
			for _, result := range rl.Results {
				mv, err := internal.TypedObjectToMapVariant(result)
				if err != nil {
					return nil, err
				}
				resultsSlice.Add(mv)
			}
			if err := reMap.SetNestedSlice(resultsSlice, "results"); err != nil {
				return nil, err
			}
		}
	*/

	return reMap.Node(), nil
}

// ToYAML converts the ResourceList to yaml.
func (mr *KubeObject) ToYAML() ([]byte, error) {
	// Sort the resources first.
	mr.Sort()
	ynode, err := mr.toYNode()
	if err != nil {
		return nil, err
	}
	doc := internal.NewDoc([]*yaml.Node{ynode}...)
	return doc.ToYAML()
}

// Sort sorts the ResourceList.items by apiVersion, kind, namespace and name.
func (mr *KubeObject) Sort() {
	//sort.Sort(mr.Items)
}
