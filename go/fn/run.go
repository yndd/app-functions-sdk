package fn

import (
	"fmt"
	"io/ioutil"
	"os"
)

func AsMain(input interface{}) error {
	err := func() error {
		var p ManagedResourceProcessor
		switch input := input.(type) {
		case ManagedResourceProcessorFunc:
			p = input
		default:
			return fmt.Errorf("unknown input type %T", input)
		}
		in, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return fmt.Errorf("unable to read from stdin: %v", err)
		}
		out, err := Run(p, in)
		// If there is an error, we don't return the error immediately.
		// We write out to stdout before returning any error.
		_, outErr := os.Stdout.Write(out)
		if outErr != nil {
			return outErr
		}
		return err
	}()
	if err != nil {
		Logf("failed to evaluate function: %v", err)
	}
	return err
}

// Run evaluates the function. input must be a Managed Resource in yaml format. A
// New Managed Resource will be returned
func Run(p ManagedResourceProcessor, input []byte) (out []byte, err error) {
	mr, err := ParseManagedResource(input)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Managed Resource: %v\n", mr)
	defer func() {
		// if we run into a panic, we still need to log the error to Results,
		// and return the ResourceList and error.
		v := recover()
		if v != nil {
			switch t := v.(type) {
			case errKubeObjectFields:
				err = &t
			case *errKubeObjectFields:
				err = t
			case errSubObjectFields:
				err = &t
			case *errSubObjectFields:
				err = t
			case errResultEnd:
				err = &t
			case *errResultEnd:
				err = t
			default:
				panic(v)
			}
			//mr.LogResult(err)
			//out, _ = mr.ToYAML()
		}
	}()
	/*
	success, fnErr := p.Process(mr)
	out, yamlErr := mr.ToYAML()
	if yamlErr != nil {
		return out, yamlErr
	}
	if fnErr != nil {
		return out, fnErr
	}
	if !success {
		return out, fmt.Errorf("error: function failure")
	}
	*/
	return out, nil
}