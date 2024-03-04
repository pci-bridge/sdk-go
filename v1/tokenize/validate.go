package tokenize

type validator interface {
	Validate() error
}
