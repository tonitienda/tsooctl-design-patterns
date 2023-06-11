package contracts

type CommandResult struct {
	Message string
}

type Command interface {
	Execute() CommandResult
}
