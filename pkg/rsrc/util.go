package rsrc

import "fmt"

func identifier(i any) (*int, *string, error) {
	if id, ok := i.(int); ok {
		return &id, nil, nil
	} else if name, ok := i.(string); ok {
		return nil, &name, nil
	}
	return nil, nil, fmt.Errorf("wrong identifier %v(%T) for resource identifier", i, i)
}
