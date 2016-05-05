package highbrow

type RetryError struct {
	msg    string
	errors []error
}

func (e *RetryError) Errors() string {
	x := ""

	for _, er := range e.errors {
		x += er.Error()
		x += " | "
	}

	return x
}
func (e *RetryError) Error() string { return e.msg }
