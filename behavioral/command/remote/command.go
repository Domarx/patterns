package remote

type Command interface {
	Execute() error
}
