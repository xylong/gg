package result

import (
	"fmt"
)

type ErrorResult struct {
	data interface{}
	err  error
}

func (r *ErrorResult) Unwrap() interface{} {
	if r.err != nil {
		panic(r.err.Error())
	}
	return r.data
}

func Result(any ...interface{}) *ErrorResult {
	switch len(any) {
	case 1:
		if any[0] == nil {
			return &ErrorResult{nil, nil}
		} else if e, ok := any[0].(error); ok {
			return &ErrorResult{nil, e}
		}
	case 2:
		if any[1] == nil {
			return &ErrorResult{any[0], nil}
		} else if e, ok := any[1].(error); ok {
			return &ErrorResult{any[0], e}
		}
	}
	return &ErrorResult{nil, fmt.Errorf("error result")}
}
