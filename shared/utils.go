package shared

type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

// Wrap the error info in an object
func NewError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}
