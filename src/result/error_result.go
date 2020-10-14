package result

type ErrorResult struct {
	err error
}

func (r *ErrorResult) Unwrap() interface{} {
	if r.err != nil {
		panic(r.err.Error())
	}
	return nil
}

func Result(err error) *ErrorResult {
	return &ErrorResult{err: err}
}
