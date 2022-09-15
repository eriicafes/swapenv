package utils

import "github.com/spf13/cobra"

type args[T any] struct {
	parser func(cmd *cobra.Command, args []string) (T, error)
	Fields T
}

// Cobra Positional Args function
func (a *args[T]) Validate(cmd *cobra.Command, rawArgs []string) error {
	fields, err := a.parser(cmd, rawArgs)

	if err != nil {
		return err
	}

	// mutate internal fields after successful validation
	a.Fields = fields
	return nil
}

// Create new Args
func ParseArgs[T any](parser func(cmd *cobra.Command, rawArgs []string) (T, error)) args[T] {
	return args[T]{
		parser: parser,
	}
}
